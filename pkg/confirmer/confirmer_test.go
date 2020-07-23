// Copyright (c) 2020 Proton Technologies AG
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

package confirmer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfirmerYes(t *testing.T) {
	c := New()

	req := c.NewRequest(1 * time.Second)

	go func() {
		assert.NoError(t, c.SetResponse(req.ID(), true))
	}()

	res, err := req.Result()
	assert.NoError(t, err)
	assert.True(t, res)
}

func TestConfirmerNo(t *testing.T) {
	c := New()

	req := c.NewRequest(1 * time.Second)

	go func() {
		assert.NoError(t, c.SetResponse(req.ID(), false))
	}()

	res, err := req.Result()
	assert.NoError(t, err)
	assert.False(t, res)
}

func TestConfirmerTimeout(t *testing.T) {
	c := New()

	req := c.NewRequest(1 * time.Second)

	go func() {
		time.Sleep(2 * time.Second)
		assert.NoError(t, c.SetResponse(req.ID(), true))
	}()

	_, err := req.Result()
	assert.Error(t, err)
}
