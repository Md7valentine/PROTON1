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

ColumnLayout {
    id: root
    property ColorScheme colorScheme

    spacing: 10

    TextArea {
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        Layout.preferredHeight: 100

        placeholderText: "Placeholder"
        label: "Label"
        hint: "Hint"
        assistiveText: "Assistive text"

        wrapMode: TextInput.Wrap
    }

    TextArea {
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        Layout.preferredHeight: 100

        text: "Value"
        placeholderText: "Placeholder"
        label: "Label"
        hint: "Hint"
        assistiveText: "Assistive text"

        wrapMode: TextInput.Wrap
    }


    TextArea {
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        Layout.preferredHeight: 100

        error: true

        text: "Value"
        placeholderText: "Placeholder"
        label: "Label"
        hint: "Hint"
        errorString: "Error message"

        wrapMode: TextInput.Wrap
    }


    TextArea {
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        Layout.preferredHeight: 100

        enabled: false

        text: "Value"
        placeholderText: "Placeholder"
        label: "Label"
        hint: "Hint"
        assistiveText: "Assistive text"

        wrapMode: TextInput.Wrap
    }

    TextArea {
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        Layout.preferredHeight: 100

        placeholderText: "Type 42 here"
        label: "42 Validator"
        hint: "Accepts only \"42\""
        assistiveText: "Type sometihng here, preferably 42"

        wrapMode: TextInput.Wrap

        validator: function(str) {
            if (str === "42") {
                return
            }

            return "Not 42"
        }
    }
}

