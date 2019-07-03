package main

import (
	"os"

	svhapi "github.com/DgINC/SVH4ALL/src/api"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func main() {
	svhapi.InitLogger()
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)

	if quickcontrols2.QQuickStyle_Name() == "" {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	}

	/*var translator = core.NewQTranslator(nil)
	if translator.Load2(core.NewQLocale(), "drawer_nav_x", "_", ":/translations", ".qm") {
		core.QCoreApplication_InstallTranslator(translator)
	} else {
		println("cannot load translator", core.QLocale_System().Name(), "check content of translations.qrc")
	}*/

	var appui = initApplicationUI()
	var engine = qml.NewQQmlApplicationEngine(nil)

	// from QML we have access to ApplicationUI as myApp
	var context = engine.RootContext()
	context.SetContextProperty("SVH4ALL", appui)

	engine.Load(core.NewQUrl3("qrc:/mainwindow.qml", 0))

	gui.QGuiApplication_Exec()
}
