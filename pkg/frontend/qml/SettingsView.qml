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
import QtQuick.Controls 2.13
import QtQuick.Controls.impl 2.13

import Proton 4.0

Item {
    id: root

    property var colorScheme
    property var backend
    default property alias items: content.children

    signal back()

    property int _leftMargin: 64
    property int _rightMargin: 64
    property int _topMargin: 32
    property int _bottomMargin: 32
    property int _spacing: 20

    // fillHeight indicates whether the SettingsView should fill all available explicit height set
    property bool fillHeight: false

    ScrollView {
        id: scrollView
        clip: true

        anchors.fill: parent

        Item {
            // can't use parent here because parent is not ScrollView (Flickable inside contentItem inside ScrollView)
            width: scrollView.availableWidth
            height: scrollView.availableHeight

            implicitHeight: children[0].implicitHeight + children[0].anchors.topMargin + children[0].anchors.bottomMargin
            // do not set implicitWidth because implicit width of ColumnLayout will be equal to maximum implicit width of
            // internal items. And if one of internal items would be a Text or Label - implicit width of those is always
            // equal to non-wrapped text (i.e. one line only). That will lead to enabling horizontal scroll when not needed
            implicitWidth: width

            ColumnLayout {
                anchors.fill: parent
                spacing: 0

                ColumnLayout {
                    id: content
                    spacing: root._spacing

                    Layout.fillWidth: true

                    Layout.topMargin: root._topMargin
                    Layout.bottomMargin: root._bottomMargin
                    Layout.leftMargin: root._leftMargin
                    Layout.rightMargin: root._rightMargin
                }

                Item {
                    id: filler
                    Layout.fillHeight: true
                    visible: !root.fillHeight
                }
            }
        }
    }

    Button {
        id: backButton
        anchors {
            top: parent.top
            left: parent.left
            topMargin: root._topMargin
            leftMargin: (root._leftMargin-backButton.width) / 2
        }
        colorScheme: root.colorScheme
        onClicked: root.back()
        icon.source: "icons/ic-arrow-left.svg"
        secondary: true
        horizontalPadding: 8
    }
}
