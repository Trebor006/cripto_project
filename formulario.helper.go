package main

import (
	"criptograms/main/data"
	"criptograms/main/encryptors"
	"criptograms/main/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
)

var menu = []string{"ZigZag", "Columnas", "Vigenere", "Primer Orden"}

func GenerarFormulario() {
	/////MENU FYNE
	//Fyne
	a := app.NewWithID("Cripto")
	w := a.NewWindow("Cripto")
	//a.Settings().SetTheme(theme.DarkTheme())

	w.SetMainMenu(makeMenu(a, w))
	//w.SetFullScreen(true)

	//w.SetContent(generarContent())

	w.Resize(fyne.NewSize(1024, 720))
	w.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	cifradosDeDesplazamiento := fyne.NewMenuItem("Cifra de Desplazamiento", nil)
	cifradosDeDesplazamiento.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Puro", func() {
			w.SetContent(puroScreen(w))
		}),
		fyne.NewMenuItem("Puro con Palabra Clave", func() {
			w.SetContent(puroConClaveScreen(w))
		}),
	)

	cifradosPorTraspocicion := fyne.NewMenuItem("Cifra por Transposición", nil)
	cifradosPorTraspocicion.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Grupos", func() {
			w.SetContent(gruposScreen(w))
		}),
		fyne.NewMenuItem("Series", func() {
			w.SetContent(seriesScreen(w))
		}),
		fyne.NewMenuItem("Columnas", func() {
			w.SetContent(columnasScreen(w))
		}),
		fyne.NewMenuItem("Filas", func() {
			w.SetContent(filasScreen(w))
		}),
		fyne.NewMenuItem("Zig-Zag", func() {
			w.SetContent(zigzagScreen(w))
		}),
	)

	cifradosPorSustitucion := fyne.NewMenuItem("Cifra por Sustitución", nil)
	sustitucionMonoAlfabetica := fyne.NewMenuItem("Mono alfabética", nil)

	sustitucionMonoAlfabetica.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Decimación Pura", func() {
			w.SetContent(decimacionPuraScreen(w))
		}),
		fyne.NewMenuItem("Transformación Afín", func() {
			w.SetContent(transformacionAfinScreen(w))
		}),
	)

	sustitucionPoliAlfabetica := fyne.NewMenuItem("Poli alfabética", nil)
	sustitucionPoliAlfabetica.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Homófonos de Primer Orden", func() {
			w.SetContent(homofonoPrimerOrdenScreen(w))
		}),
		fyne.NewMenuItem("Homófonos de Orden Mayor", func() {
			w.SetContent(homofonoOrdenMayorScreen(w))
		}),
	)

	cifradosPorSustitucion.ChildMenu = fyne.NewMenu("",
		sustitucionMonoAlfabetica,
		fyne.NewMenuItemSeparator(),
		sustitucionPoliAlfabetica,
	)

	cifradosPorSustitucionMonografica := fyne.NewMenuItem("Sustitucion Monogramica Polialfabeto", nil)
	cifradosPorSustitucionMonografica.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Polialfabeticos periodicos", func() {
			w.SetContent(polialfabeticoPeriodicosScreen(w))
		}),
	)

	metodosDeCifrado := fyne.NewMenu("Metodos de Cifrado",
		cifradosDeDesplazamiento,
		fyne.NewMenuItemSeparator(),
		cifradosPorTraspocicion,
		fyne.NewMenuItemSeparator(),
		cifradosPorSustitucion,
		fyne.NewMenuItemSeparator(),
		cifradosPorSustitucionMonografica,
	)

	main := fyne.NewMainMenu(
		metodosDeCifrado,
	)

	return main
}

func puroScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Cifrado Cesar")
	encryptor := GetEncryptor(PURO)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func puroConClaveScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Cifrado Cesar con Clave")
	encryptor := GetEncryptor(PURO_CON_CLAVE)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func homofonoOrdenMayorScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Orden Mayor")
	encryptor := GetEncryptor(ORDEN_MAYOR)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func homofonoPrimerOrdenScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Primer Orden")
	encryptor := GetEncryptor(PRIMER_ORDEN)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func decimacionPuraScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Decimacion Pura")
	encryptor := GetEncryptor(DECIMACION_PURA)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func transformacionAfinScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Transformacion Afin")
	encryptor := GetEncryptor(TRANSFORMACION_AFIN)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func polialfabeticoPeriodicosScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Polialfabetico Periodicos - Vigenere")
	encryptor := GetEncryptor(POLIALFABETICO_PERIODICO)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func columnasScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Columnas")
	encryptor := GetEncryptor(COLUMNAS)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Inicial")
	textoInicial.SetMinRowsVisible(5)

	nroColumnas := newNumEntry()
	nroColumnas.SetPlaceHolder("Nro columnas")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		data := data.Data{}
		data.Message = textoInicial.Text
		data.NroColumns, _ = strconv.Atoi(nroColumnas.Text)

		textoCifrado := encryptor.Cypher(data)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		data := data.Data{}
		data.EncryptedMessage = textoInicial.Text
		data.NroColumns, _ = strconv.Atoi(nroColumnas.Text)

		textoCifrado := encryptor.Decrypt(data)
		textoResultante.SetText(textoCifrado)
	})

	return container.NewVBox(
		separadorTitulo,
		textoInicial,
		nroColumnas,
		textoResultante,
		botonEncriptar,
		botonDesencriptar,
	)
}

func filasScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Filas")
	encryptor := GetEncryptor(FILAS)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func gruposScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Grupos")
	encryptor := GetEncryptor(GRUPOS)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func seriesScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("Series")
	encryptor := GetEncryptor(SERIES)

	return obetnerFormularioGenerico(encryptor, separadorTitulo)
}

func zigzagScreen(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := generarSeparadorTitulo("ZigZag")
	encryptor := GetEncryptor(ZIG_ZAG)

	return obetnerFormularioGenerico2(encryptor, separadorTitulo)
}

func obetnerFormularioGenerico(
	encryptor encryptors.EncryptorInterface,
	separadorTitulo *fyne.Container,
) fyne.CanvasObject {

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		data := data.Data{}
		data.Message = textoInicial.Text
		textoCifrado := encryptor.Cypher(data)
		textoResultante.SetText(textoCifrado)
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		data := data.Data{}
		data.EncryptedMessage = textoInicial.Text
		textoCifrado := encryptor.Decrypt(data)
		textoResultante.SetText(textoCifrado)
	})

	//buttonsLayout := widget.FormItem{
	//	Text: "", Widget: botonEncriptar, HintText: "Introduzca el texto a Encriptar/Desencriptar",
	//	//{Text: "", Widget: botonDesencriptar, HintText: "Aqui se mostrara el resultado"},
	//}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Resultado", Widget: textoResultante, HintText: "Aqui se mostrara el resultado"},
		},
		//OnCancel: func() {
		//	//fmt.Println("Cancelled")
		//},
		//OnSubmit: func() {
		//	//fmt.Println("Form submitted")
		//	//fyne.CurrentApp().SendNotification(&fyne.Notification{
		//	//	Title:   "Form for: " + name.Text,
		//	//	Content: largeText.Text,
		//	//})
		//},
	}

	//toolTips := widget.NewToolbar(
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.MoveUpIcon(), func() {
	//		fmt.Println("Copy")
	//	}))

	//form.AppendItem(buttonsLayout)
	//form.AppendItem(botonEncriptar)

	form.Append("", botonEncriptar)
	form.Append("", botonDesencriptar)
	//return form

	return container.NewVBox(
		separadorTitulo,
		//toolTips,
		form,
	)
}

func obetnerFormularioGenerico2(
	encryptor encryptors.EncryptorInterface,
	separadorTitulo *fyne.Container,
) fyne.CanvasObject {

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	nroColumnas := newNumEntry()
	nroColumnas.SetPlaceHolder("Nro columnas")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		data := data.Data{}
		data.Message = textoInicial.Text
		data.NroColumns, _ = strconv.Atoi(nroColumnas.Text)

		textoCifrado := encryptor.Cypher(data)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		data := data.Data{}
		data.EncryptedMessage = textoInicial.Text
		data.NroColumns, _ = strconv.Atoi(nroColumnas.Text)

		textoCifrado := encryptor.Decrypt(data)
		textoResultante.SetText(textoCifrado)
	})

	//buttonsLayout := widget.FormItem{
	//	Text: "", Widget: botonEncriptar, HintText: "Introduzca el texto a Encriptar/Desencriptar",
	//	//{Text: "", Widget: botonDesencriptar, HintText: "Aqui se mostrara el resultado"},
	//}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Nro", Widget: nroColumnas, HintText: "Introduzca el Nro"},
			{Text: "Resultado", Widget: textoResultante, HintText: "Aqui se mostrara el resultado"},
		},
		//OnCancel: func() {
		//	//fmt.Println("Cancelled")
		//},
		//OnSubmit: func() {
		//	//fmt.Println("Form submitted")
		//	//fyne.CurrentApp().SendNotification(&fyne.Notification{
		//	//	Title:   "Form for: " + name.Text,
		//	//	Content: largeText.Text,
		//	//})
		//},
	}

	//toolTips := widget.NewToolbar(
	//	widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.MoveUpIcon(), func() {
	//		fmt.Println("Copy")
	//	}))

	//form.AppendItem(buttonsLayout)
	//form.AppendItem(botonEncriptar)

	form.Append("", botonEncriptar)
	form.Append("", botonDesencriptar)
	//return form

	return container.NewVBox(
		separadorTitulo,
		//toolTips,
		form,
	)
}

func generarSeparadorTitulo(nombreTitulo string) *fyne.Container {
	tituloLabel := widget.NewLabel(nombreTitulo)
	separadorTitulo := container.NewBorder(container.NewVBox(tituloLabel, widget.NewSeparator()), nil, nil, nil, nil)
	return separadorTitulo
}

type numEntry struct {
	widget.Entry
}

func newNumEntry() *numEntry {
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
