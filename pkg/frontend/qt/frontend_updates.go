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

//go:build build_qt
// +build build_qt

package qt

import (
	"sync"

	"github.com/ljanyst/peroxide/pkg/config/settings"
	"github.com/ljanyst/peroxide/pkg/updater"
)

var checkingUpdates = sync.Mutex{}

func (f *FrontendQt) checkUpdates() error {
	version, err := f.updater.Check()
	if err != nil {
		return err
	}

	f.SetVersion(version)
	return nil
}

func (f *FrontendQt) checkUpdatesAndNotify(isRequestFromUser bool) {
	checkingUpdates.Lock()
	defer checkingUpdates.Unlock()
	defer f.qml.CheckUpdatesFinished()

	if err := f.checkUpdates(); err != nil {
		f.log.WithError(err).Error("An error occurred while checking updates")
		if isRequestFromUser {
			f.qml.UpdateManualError()
		}
		return
	}

	if !f.updater.IsUpdateApplicable(f.newVersionInfo) {
		f.log.Debug("No need to update")
		if isRequestFromUser {
			f.qml.UpdateIsLatestVersion()
		}
		return
	}

	if !f.updater.CanInstall(f.newVersionInfo) {
		f.log.Debug("A manual update is required")
		f.qml.UpdateManualReady(f.newVersionInfo.Version.String())
		return
	}
}

func (f *FrontendQt) updateForce() {
	checkingUpdates.Lock()
	defer checkingUpdates.Unlock()

	version := ""
	if err := f.checkUpdates(); err == nil {
		version = f.newVersionInfo.Version.String()
	}

	f.qml.UpdateForce(version)
}

func (f *FrontendQt) setIsAutomaticUpdateOn() {
	f.qml.SetIsAutomaticUpdateOn(f.settings.GetBool(settings.AutoUpdateKey))
}

func (f *FrontendQt) toggleAutomaticUpdate(makeItEnabled bool) {
	f.qml.SetIsAutomaticUpdateOn(makeItEnabled)
	isEnabled := f.settings.GetBool(settings.AutoUpdateKey)
	if makeItEnabled == isEnabled {
		return
	}

	f.settings.SetBool(settings.AutoUpdateKey, makeItEnabled)

	f.checkUpdatesAndNotify(false)
}

func (f *FrontendQt) setIsBetaEnabled() {
	channel := f.bridge.GetUpdateChannel()
	f.qml.SetIsBetaEnabled(channel == updater.EarlyChannel)
}

func (f *FrontendQt) toggleBeta(makeItEnabled bool) {
	channel := updater.StableChannel
	if makeItEnabled {
		channel = updater.EarlyChannel
	}

	f.bridge.SetUpdateChannel(channel)

	f.setIsBetaEnabled()

	// Immediately check the updates to set the correct landing page link.
	f.checkUpdates()
}
