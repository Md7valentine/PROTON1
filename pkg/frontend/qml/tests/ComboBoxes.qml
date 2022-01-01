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
import QtQuick.Window 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

import "../Proton"

RowLayout {
    id: root
    property ColorScheme colorScheme

    ColumnLayout {
        Layout.fillWidth: true
        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            colorScheme: root.colorScheme
        }

        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            editable: true
            colorScheme: root.colorScheme
        }
    }

    ColumnLayout {
        Layout.fillWidth: true
        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            colorScheme: root.colorScheme
            enabled: false
        }

        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            editable: true
            colorScheme: root.colorScheme
            enabled: false
        }
    }

    ColumnLayout {
        Layout.fillWidth: true
        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            colorScheme: root.colorScheme
            LayoutMirroring.enabled: true
        }

        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            editable: true
            colorScheme: root.colorScheme
            LayoutMirroring.enabled: true
        }
    }

    ColumnLayout {
        Layout.fillWidth: true
        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            colorScheme: root.colorScheme
            enabled: false
            LayoutMirroring.enabled: true
        }

        ComboBox {
            Layout.fillWidth: true
            model: ["First", "Second", "Third"]
            editable: true
            colorScheme: root.colorScheme
            enabled: false
            LayoutMirroring.enabled: true
        }
    }
}
