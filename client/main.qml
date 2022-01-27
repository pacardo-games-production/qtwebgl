import QtQuick
import QtQuick.Window
import QtQml
import QtWebEngine

Window {
    width: 1024
    height: 750
    visible: true
    WebEngineView {
        anchors.fill: parent
        url: "http://localhost:4000/"
    }
}
