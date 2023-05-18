package formulario

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var menu = []string{"ZigZag", "Columnas", "Vigenere", "Primer Orden"}

func GenerarFormulario() {
	/////MENU FYNE
	//Fyne
	a := app.NewWithID("Cripto")
	w := a.NewWindow("Cripto")
	//a.Settings().SetTheme(theme.DarkTheme())

	w.SetMainMenu(GenerarMenu(a, w))
	//w.SetFullScreen(true)

	//w.SetContent(generarContent())

	w.Resize(fyne.NewSize(1024, 720))
	w.ShowAndRun()
}

func GenerarSeparadorTitulo(nombreTitulo string) *fyne.Container {
	tituloLabel := widget.NewLabel(nombreTitulo)
	separadorTitulo := container.NewBorder(container.NewVBox(tituloLabel, widget.NewSeparator()), nil, nil, nil, nil)
	return separadorTitulo
}

type numEntry struct {
	widget.Entry
}

func NewNumEntry() *numEntry {
	e := &numEntry{}
	e.ExtendBaseWidget(e)
	e.Validator = validation.NewRegexp(`\d`, "Must contain a number")
	return e
}

func makeCell() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{128, 128, 128, 255})
	rect.SetMinSize(fyne.NewSize(30, 30))
	return rect
}

//entry := widget.NewEntry()
//entry.SetPlaceHolder("Entry")
//entryDisabled := widget.NewEntry()
//entryDisabled.SetText("Entry (disabled)")
//entryDisabled.Disable()
//entryValidated := newNumEntry()
//entryValidated.SetPlaceHolder("Must contain a number")
//entryMultiLine := widget.NewMultiLineEntry()
//entryMultiLine.SetPlaceHolder("MultiLine Entry")
//entryMultiLine.SetMinRowsVisible(4)
