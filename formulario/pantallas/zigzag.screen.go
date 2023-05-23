package pantallas

import (
	"cripto_project/main/data"
	"cripto_project/main/encryptors"
	"cripto_project/main/formulario/widgets"
	"cripto_project/main/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func ZigZagGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("ZigZag")
	encryptor := encryptors.GetEncryptor(encryptors.ZIG_ZAG)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	nroRails := widgets.NewNumEntry()
	nroRails.SetPlaceHolder("Nro Rails")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()

		nRails, e := strconv.Atoi(nroRails.Text)
		if !validarDatosIniciales(textoInicial.Text) {
			widgets.MostrarError("Revise los datos de la cadena a procesar", w)
		} else if e != nil {
			widgets.MostrarError("Debe asignar un numero de Rails", w)
		} else if nRails <= 2 {
			widgets.MostrarError("El numero de Rails debe ser de al menos 3", w)
		} else {
			dataACifrar := data.Data{Message: textoInicial.Text, NroRails: nRails}

			textoCifrado := encryptor.Cypher(dataACifrar)
			textoResultante.SetText(util.Format(textoCifrado))
		}
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()

		nRails, e := strconv.Atoi(nroRails.Text)
		if !validarDatosIniciales(textoInicial.Text) {
			widgets.MostrarError("Revise los datos de la cadena a procesar", w)
		} else if e != nil {
			widgets.MostrarError("Debe asignar un numero de Rails", w)
		} else if nRails <= 2 {
			widgets.MostrarError("El numero de Rails debe ser de al menos 3", w)
		} else {
			dataADescifrar := data.Data{EncryptedMessage: textoInicial.Text, NroRails: nRails}

			textoDescifrado := encryptor.Decrypt(dataADescifrar)
			textoResultante.SetText(util.Format(textoDescifrado))
		}
	})

	//buttonsLayout := widget.FormItem{
	//	Text: "", Widget: botonEncriptar, HintText: "Introduzca el texto a Encriptar/Desencriptar",
	//	//{Text: "", Widget: botonDesencriptar, HintText: "Aqui se mostrara el resultado"},
	//}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Nro de Rails", Widget: nroRails, HintText: "Introduzca el Nro Rails"},
			{Text: "Resultado", Widget: textoResultante, HintText: "Aqui se mostrara el resultado"},
		},
	}

	toolTips := widgets.GenerateToolTipCopyButton(textoInicial, textoResultante)
	//form.AppendItem(buttonsLayout)
	//form.AppendItem(botonEncriptar)

	form.Append("", botonEncriptar)
	form.Append("", botonDesencriptar)
	//return form

	return container.NewVBox(
		separadorTitulo,
		toolTips,
		form,
	)
}
