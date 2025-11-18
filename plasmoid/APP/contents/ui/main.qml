import QtQml 2.15
import QtQuick 2.0
import QtQuick.Layouts 1.0
import org.kde.plasma.components as PlasmaComponents
import org.kde.plasma.plasmoid
import org.kde.plasma.core as PlasmaCore
import org.kde.plasma.plasma5support as Plasma5Support

PlasmoidItem {
    id: widget
    
    // Plasmoid.backgroundHints: PlasmaCore.Types.ShadowBackground | PlasmaCore.Types.ConfigurableBackground
    
    preferredRepresentation: fullRepresentation
    fullRepresentation: Item {
        // applet default size
        Layout.minimumWidth: 300
        Layout.minimumHeight: 400
        Layout.preferredWidth: Layout.minimumWidth
        Layout.preferredHeight: Layout.minimumHeight
    }   

    ColumnLayout {
        id: container
        spacing: 5
        anchors.fill: parent
        anchors.margins: 5
        
        ColumnLayout {
            spacing: 10
            Rectangle {
                color: "transparent"
                height: 50
                Layout.fillWidth: true
                Text {
                    font.bold: true
                    font.pointSize: 18
                    font.family: "Helvetica"
                    text: "<h2>" + i18n("To-Do List") + "</h2>"
                    color: "#fff"
                }
            }
            Rectangle{
                color: "transparent" 
                height: 50
                Layout.fillWidth: true
                PlasmaComponents.TextField {
                    height: parent.height
                    width: parent.width-54
                    anchors.left: parent.left
                    anchors.verticalCenter: parent.verticalCenter
                    placeholderText: i18n("New Task...")
                    font.family: "Helvetica"                        
                    font.pointSize: 15
        
                }
                PlasmaComponents.Button {
                    height: parent.height
                    width: height
                    anchors.right: parent.right
                    anchors.verticalCenter: parent.verticalCenter
                    icon.name: "list-add"
                }
            }
        }   





        Rectangle {
            color: Qt.rgba(255, 255, 255, 0.1) 
            height: 230
            Layout.fillWidth: true
            PlasmaComponents.ScrollView {
                id: scrollView
                ListView {
                    model: 100
                    delegate: PlasmaComponents.CheckBox {
                    text: i18n("CheckBox #%1", index+1)
                }
            }
        }

        }
        Rectangle {
            color: "#ff0" // Yellow
            height: 40
            Layout.fillWidth: true
        }
    }
  
}

