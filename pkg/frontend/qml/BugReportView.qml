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

import Proton 4.0

SettingsView {
    id: root

    fillHeight: true

    property var selectedAddress

    Label {
        text: qsTr("Report a problem")
        colorScheme: root.colorScheme
        type: Label.Heading
    }


    TextArea {
        id: description
        property int _minLength: 150
        property int _maxLength: 800

        label: qsTr("Description")
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        Layout.fillHeight: true
        Layout.minimumHeight: heightForLinesVisible(4)
        hint: description.text.length + "/" + _maxLength
        placeholderText: qsTr("Tell us what went wrong or isn't working (min. %1 characters).").arg(_minLength)

        validator: function(text) {
            if (description.text.length < description._minLength) {
                return qsTr("Enter a problem description (min. %1 characters).").arg(_minLength)
            }

            if (description.text.length > description._maxLength) {
                return qsTr("Enter a problem description (max. %1 characters).").arg(_maxLength)
            }

            return
        }

        onTextChanged: {
            // Rise max length error imidiatly while typing
            if (description.text.length > description._maxLength) {
                validate()
            }
        }

        KeyNavigation.priority: KeyNavigation.BeforeItem
        KeyNavigation.tab: address

        // set implicitHeight to explicit height because se don't
        // want TextArea implicitHeight (which is height of all text)
        // to be considered in SettingsView internal scroll view
        implicitHeight: height
    }


    TextField {
        id: address

        label: qsTr("Your contact email")
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        placeholderText: qsTr("e.g. jane.doe@protonmail.com")

        validator: function(str) {
            if (!isValidEmail(str)) {
                return qsTr("Enter valid email address")
            }
            return
        }
    }

    TextField {
        id: emailClient

        label: qsTr("Your email client (including version)")
        colorScheme: root.colorScheme
        Layout.fillWidth: true
        placeholderText: qsTr("e.g. Apple Mail 14.0")

        validator: function(str) {
            if (str.length === 0) {
                return qsTr("Enter an email client name and version")
            }
            return
        }
    }


    RowLayout {
        CheckBox {
            id: includeLogs
            text: qsTr("Include my recent logs")
            colorScheme: root.colorScheme
            checked: true
        }
        Button {
            Layout.leftMargin: 12
            text: qsTr("View logs")
            secondary: true
            colorScheme: root.colorScheme
            onClicked: Qt.openUrlExternally(root.backend.logsPath)
        }
    }

    TextEdit {
        text: {
            var address = "bridge@protonmail.com"
            var mailTo = `<a href="mailto://${address}">${address}</a>`
            return "<style>a:link { color: " + root.colorScheme.interaction_norm + "; }</style>" +qsTr(
                "These reports are not end-to-end encrypted. In case of sensitive information, contact us at %1."
            ).arg(mailTo)
        }
        onLinkActivated: Qt.openUrlExternally(link)
        textFormat: Text.RichText

        readOnly: true


        Layout.fillWidth: true
        color: root.colorScheme.text_weak
        font.family: ProtonStyle.font_family
        font.weight: ProtonStyle.fontWeight_400
        font.pixelSize: ProtonStyle.caption_font_size
        font.letterSpacing: ProtonStyle.caption_letter_spacing
        // No way to set lineHeight: Style.caption_line_height
        selectionColor: root.colorScheme.interaction_norm
        selectedTextColor: root.colorScheme.text_invert
        wrapMode: Text.WordWrap
        selectByMouse: true
    }

    Button {
        id: sendButton
        text: qsTr("Send")
        colorScheme: root.colorScheme

        onClicked: {
            description.validate()
            address.validate()
            emailClient.validate()

            if (description.error || address.error || emailClient.error) {
                return
            }

            submit()
        }

        Connections {target: root.backend; onReportBugFinished: sendButton.loading = false }
    }

    function setDefaultValue() {
        description.text = ""
        address.text = root.selectedAddress
        emailClient.text = root.backend.currentEmailClient
        includeLogs.checked = true
    }

    function isValidEmail(text){
        var reEmail = /\w+@\w+\.\w+/
        return reEmail.test(text)
    }

    function submit() {
        sendButton.loading = true
        root.backend.reportBug(
            description.text,
            address.text,
            emailClient.text,
            includeLogs.checked
        )
    }

    onVisibleChanged: {
        root.setDefaultValue()
    }
}
