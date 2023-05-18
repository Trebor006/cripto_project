package formulario

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func GenerarFormulario() {
	a := app.NewWithID("Cripto")
	w := a.NewWindow("Cripto")
	a.Settings().SetTheme(theme.DarkTheme())

	w.SetMainMenu(GenerarMenu(a, w))
	//w.SetFullScreen(true)

	w.Resize(fyne.NewSize(1024, 720))
	w.ShowAndRun()
}
