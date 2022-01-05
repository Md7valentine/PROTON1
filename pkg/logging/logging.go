// Copyright (c) 2021 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package logging

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/ljanyst/peroxide/pkg/constants"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

const (
	// MaxLogSize defines the maximum log size we should permit: 5 MB
	//
	// The Zendesk limit for an attachement is 50MB and this is what will
	// be allowed via the API. However, if that fails for some reason, the
	// fallback is sending the report via email, which has a limit of 10mb
	// total or 7MB per file. Since we can produce up to 6 logs, and we
	// compress all the files (avarage compression - 80%), we need to have
	// a limit of 30MB total before compression, hence 5MB per log file.
	MaxLogSize = 5 * 1024 * 1024

	// MaxLogs defines how many log files should be kept.
	MaxLogs = 3
)

func Init(logsPath string) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.StampMilli,
	})
	logrus.AddHook(&writer.Hook{
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		},
	})

	rotator, err := NewRotator(MaxLogSize, func() (io.WriteCloser, error) {
		return os.Create(filepath.Join(logsPath, getLogName(constants.Version, constants.Revision)))
	})
	if err != nil {
		return err
	}

	logrus.SetOutput(rotator)
	return nil
}

// SetLevel will change the level of logging and in case of Debug or Trace
// level it will also prevent from writing to file. Setting level to Info or
// higher will not set writing to file again if it was previously cancelled by
// Debug or Trace.
func SetLevel(level string) {
	if lvl, err := logrus.ParseLevel(level); err == nil {
		logrus.SetLevel(lvl)
	}

	if logrus.GetLevel() == logrus.DebugLevel || logrus.GetLevel() == logrus.TraceLevel {
		// The hook to print panic, fatal and error to stderr is always
		// added. We want to avoid log duplicates by replacing all hooks
		_ = logrus.StandardLogger().ReplaceHooks(logrus.LevelHooks{})
		logrus.SetOutput(os.Stderr)
	}
}

func getLogName(version, revision string) string {
	return fmt.Sprintf("v%v_%v_%v.log", version, revision, time.Now().Unix())
}

func MatchLogName(name string) bool {
	return regexp.MustCompile(`^v.*\.log$`).MatchString(name)
}
