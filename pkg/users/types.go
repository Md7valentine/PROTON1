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

package users

import (
	"github.com/ljanyst/peroxide/pkg/store"
	"github.com/ljanyst/peroxide/pkg/users/credentials"
)

type Locator interface {
	Clear() error
}

type CredentialsStorer interface {
	List() (userIDs []string, err error)
	Add(userID, userName, uid, ref string, mailboxPassword []byte, emails []string) (*credentials.Credentials, error)
	Get(userID string) (*credentials.Credentials, error)
	SwitchAddressMode(userID string) (*credentials.Credentials, error)
	UpdateEmails(userID string, emails []string) (*credentials.Credentials, error)
	UpdatePassword(userID string, password []byte) (*credentials.Credentials, error)
	UpdateToken(userID, uid, ref string) (*credentials.Credentials, error)
	Logout(userID string) (*credentials.Credentials, error)
	Delete(userID string) error
}

type StoreMaker interface {
	New(user store.BridgeUser) (*store.Store, error)
	Remove(userID string) error
}
