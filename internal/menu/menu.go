package menu

import (
	"fyne.io/fyne/v2"
	"github.com/linuxoid69/iknow/internal/windows"
)

func MenuMain(app fyne.App) (menu *fyne.Menu) {
	menu = fyne.NewMenu("iKnow",
		fyne.NewMenuItem("Open", windows.WindowMain(app).Show),
		fyne.NewMenuItem("About", windows.WindowAbout(app).Show),
	)
	return menu
}
