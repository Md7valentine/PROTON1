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

// Package base implements a common application base currently shared by bridge and IE.
// The base includes the following:
//  - access to standard filesystem locations like config, cache, logging dirs
//  - an extensible crash handler
//  - versioned cache directory
//  - persistent settings
//  - event listener
//  - credentials store
//  - pmapi Manager
// In addition, the base initialises logging and reacts to command line arguments
// which control the log verbosity and enable cpu/memory profiling.
package base

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/allan-simon/go-singleinstance"
	"github.com/ljanyst/peroxide/pkg/api"
	"github.com/ljanyst/peroxide/pkg/config/cache"
	"github.com/ljanyst/peroxide/pkg/config/settings"
	"github.com/ljanyst/peroxide/pkg/config/tls"
	"github.com/ljanyst/peroxide/pkg/config/useragent"
	"github.com/ljanyst/peroxide/pkg/constants"
	"github.com/ljanyst/peroxide/pkg/cookies"
	"github.com/ljanyst/peroxide/pkg/events"
	"github.com/ljanyst/peroxide/pkg/keychain"
	"github.com/ljanyst/peroxide/pkg/listener"
	"github.com/ljanyst/peroxide/pkg/locations"
	"github.com/ljanyst/peroxide/pkg/logging"
	"github.com/ljanyst/peroxide/pkg/pmapi"
	"github.com/ljanyst/peroxide/pkg/users/credentials"
	"github.com/sirupsen/logrus"
)

type Base struct {
	Locations *locations.Locations
	Settings  *settings.Settings
	Lock      *os.File
	Cache     *cache.Cache
	Listener  listener.Listener
	Creds     *credentials.Store
	CM        pmapi.Manager
	CookieJar *cookies.Jar
	UserAgent *useragent.UserAgent
	TLS       *tls.TLS

	Name    string // the app's name
	usage   string // the app's usage description
	command string // the command used to launch the app (either the exe path or the launcher path)
	restart bool   // whether the app is currently set to restart

	teardown []func() error // actions to perform when app is exiting
}

func New( // nolint[funlen]
	appName,
	appUsage,
	configName,
	updateURLName,
	keychainName,
	cacheVersion string,
) (*Base, error) {
	userAgent := useragent.New()

	rand.Seed(time.Now().UnixNano())
	os.Args = StripProcessSerialNumber(os.Args)

	locationsProvider, err := locations.NewDefaultProvider(filepath.Join(constants.VendorName, configName))
	if err != nil {
		return nil, err
	}

	locations := locations.New(locationsProvider, configName)

	logsPath, err := locations.ProvideLogsPath()
	if err != nil {
		return nil, err
	}
	if err := logging.Init(logsPath); err != nil {
		return nil, err
	}

	if err := migrateFiles(configName); err != nil {
		logrus.WithError(err).Warn("Old config files could not be migrated")
	}

	if err := locations.Clean(); err != nil {
		return nil, err
	}

	settingsPath, err := locations.ProvideSettingsPath()
	if err != nil {
		return nil, err
	}
	settingsObj := settings.New(settingsPath)

	lock, err := singleinstance.CreateLockFile(locations.GetLockFile())
	if err != nil {
		logrus.Warnf("%v is already running", appName)
		return nil, api.CheckOtherInstanceAndFocus(settingsObj.GetInt(settings.APIPortKey))
	}

	cachePath, err := locations.ProvideCachePath()
	if err != nil {
		return nil, err
	}
	cache, err := cache.New(cachePath, cacheVersion)
	if err != nil {
		return nil, err
	}
	if err := cache.RemoveOldVersions(); err != nil {
		return nil, err
	}

	listener := listener.New()
	events.SetupEvents(listener)

	// If we can't load the keychain for whatever reason,
	// we signal to frontend and supply a dummy keychain that always returns errors.
	kc, err := keychain.NewKeychain(settingsObj, keychainName)
	if err != nil {
		return nil, err
	}

	cfg := pmapi.NewConfig(configName, constants.Version)
	cfg.GetUserAgent = userAgent.String
	cfg.UpgradeApplicationHandler = func() { listener.Emit(events.UpgradeApplicationEvent, "") }
	cfg.TLSIssueHandler = func() { listener.Emit(events.TLSCertIssue, "") }

	cm := pmapi.New(cfg)

	cm.AddConnectionObserver(pmapi.NewConnectionObserver(
		func() { listener.Emit(events.InternetOffEvent, "") },
		func() { listener.Emit(events.InternetOnEvent, "") },
	))

	jar, err := cookies.NewCookieJar(settingsObj)
	if err != nil {
		return nil, err
	}

	cm.SetCookieJar(jar)

	return &Base{
		Locations: locations,
		Settings:  settingsObj,
		Lock:      lock,
		Cache:     cache,
		Listener:  listener,
		Creds:     credentials.NewStore(kc),
		CM:        cm,
		CookieJar: jar,
		UserAgent: userAgent,
		TLS:       tls.New(settingsPath),

		Name: appName,
	}, nil
}
