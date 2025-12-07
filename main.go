package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/linuxoid69/iknow/internal/menu"
	"github.com/linuxoid69/iknow/internal/windows"
)

func main() {
	application := app.NewWithID("com.linuxoid69.iknow")

	if desk, ok := application.(desktop.App); ok {
		mainIcon, _ := menu.GetIcon("icons/main.png")
		desk.SetSystemTrayIcon(fyne.NewStaticResource("mainIcon", mainIcon))
		desk.SetSystemTrayMenu(menu.MenuMain(application))
	}

	
	windows.WindowMain(application).ShowAndRun()
}
