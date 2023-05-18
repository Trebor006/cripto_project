package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
	"os/exec"
)

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

func GenerateToolTipCopyButton(textoInicial *widget.Entry, textoResultante *widget.Entry) *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			textoInicial.SetText(textoResultante.Text)
			textoResultante.SetText("")
		}))
}

func LimpiarConsola() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
