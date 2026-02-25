package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	titleAbout string = "About"
	titleMain  string = "iKnow"
)

func WindowAbout(app fyne.App) (w fyne.Window) {

	w = app.NewWindow(titleAbout)

	githubLink := widget.NewHyperlink("", nil)

	githubLink.SetText("Github")
	githubLink.Move(fyne.NewPos(100, 100))

	w.SetContent(
		container.NewVBox(
			githubLink,
		),
	)

	w.Resize(fyne.NewSize(400, 200))
	w.SetCloseIntercept(w.Hide)
	w.CenterOnScreen()
	return w
}

func WindowMain(app fyne.App) (w fyne.Window) {
	w = app.NewWindow(titleMain)
	w.SetContent(widget.NewLabel(titleMain))
	w.SetCloseIntercept(w.Hide)
	w.Resize(fyne.NewSize(400, 400))
	w.CenterOnScreen()
	return w
}
