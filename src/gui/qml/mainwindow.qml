import QtQuick 2.12
import QtQuick.Window 2.12

Window {
    id: window
    visible: true
    width: 640
    height: 480
    property alias windowTitle: window.title
    title: qsTr("Hello World")

    PathView {
        id: pathView
        anchors.bottomMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        anchors.right: parent.right
        anchors.rightMargin: 0
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.top: parent.top
        path: Path {
            startX: 107
            startY: 223

            PathCubic {
                x: 165
                y: 42
                control2X: 385.333
                control1X: 306.333
                control2Y: 46.333
                control1Y: 162.333
            }

            PathCubic {
                x: 107
                y: 223
                control2X: 26.667
                control1X: 57.667
                control2Y: 83.333
                control1Y: 56.333
            }
        }
        delegate: Column {
            spacing: 5
            Rectangle {
                width: 40
                height: 40
                color: colorCode
                anchors.horizontalCenter: parent.horizontalCenter
            }

            Text {
                x: 5
                text: name
                font.bold: true
                anchors.horizontalCenter: parent.horizontalCenter
            }
        }
        model: ListModel {
            ListElement {
                name: "Grey"
                colorCode: "grey"
            }

            ListElement {
                name: "Red"
                colorCode: "red"
            }

            ListElement {
                name: "Blue"
                colorCode: "blue"
            }

            ListElement {
                name: "Green"
                colorCode: "green"
            }
        }
    }
}

/*##^## Designer {
    D{i:1;anchors_height:278;anchors_width:393;anchors_x:113;anchors_y:46}
}
 ##^##*/
