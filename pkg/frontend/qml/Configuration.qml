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

import QtQuick 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.impl 2.12

import Proton 4.0

Rectangle {
    id: root

    property ColorScheme colorScheme
    property string title
    property string hostname
    property string port
    property string username
    property string password
    property string security

    implicitWidth: 304
    implicitHeight: content.height + 2*root._margin

    color: root.colorScheme.background_norm
    radius: 9

    property int _margin: 24

    ColumnLayout {
        id: content
        width: root.width - 2*root._margin
        anchors{
            top: root.top
            left: root.left
            leftMargin   : root._margin
            rightMargin  : root._margin
            topMargin    : root._margin
            bottomMargin : root._margin
        }

        spacing: 12

        Label {
            colorScheme: root.colorScheme
            text: root.title
            type: Label.Body_semibold
        }

        Item{}

        ConfigurationItem{ colorScheme: root.colorScheme; label: qsTr("Hostname") ; value: root.hostname }
        ConfigurationItem{ colorScheme: root.colorScheme; label: qsTr("Port")     ; value: root.port     }
        ConfigurationItem{ colorScheme: root.colorScheme; label: qsTr("Username") ; value: root.username }
        ConfigurationItem{ colorScheme: root.colorScheme; label: qsTr("Password") ; value: root.password }
        ConfigurationItem{ colorScheme: root.colorScheme; label: qsTr("Security") ; value: root.security }
    }
}

