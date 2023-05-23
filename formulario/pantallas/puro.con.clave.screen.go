package pantallas

import (
	"cripto_project/main/data"
	"cripto_project/main/encryptors"
	"cripto_project/main/formulario/widgets"
	"cripto_project/main/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func PuroConClaveGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Cifrado Cesar con Clave")
	metodoCifrado := encryptors.GetEncryptor(encryptors.PURO_CON_CLAVE)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	clave := widget.NewEntry()
	clave.SetPlaceHolder("Clave")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()
		if !validarDatosIniciales(textoInicial.Text) {
			widgets.MostrarError("Revise los datos de la cadena a procesar", w)
		} else if !validarDatosIniciales(clave.Text) {
			widgets.MostrarError("Revise los datos de la Clave", w)
		} else {
			dataAEncriptar := data.Data{Message: textoInicial.Text, Clave: clave.Text}

			textoCifrado := metodoCifrado.Cypher(dataAEncriptar)
			textoResultante.SetText(util.Format(textoCifrado))
		}
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()
		if !validarDatosIniciales(textoInicial.Text) {
			widgets.MostrarError("Revise los datos de la cadena a procesar", w)
		} else if !validarDatosIniciales(clave.Text) {
			widgets.MostrarError("Revise los datos de la Clave", w)
		} else {
			dataADesencriptar := data.Data{EncryptedMessage: textoInicial.Text, Clave: clave.Text}

			textoClaro := metodoCifrado.Decrypt(dataADesencriptar)
			textoResultante.SetText(util.Format(textoClaro))
		}
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Clave", Widget: clave, HintText: "Introduzca la clave"},
			{Text: "Resultado", Widget: textoResultante, HintText: "Aqui se mostrara el resultado"},
		},
	}

	toolTips := widgets.GenerateToolTipCopyButton(textoInicial, textoResultante)

	//form.AppendItem(buttonsLayout)
	//form.AppendItem(botonEncriptar)

	form.Append("", botonEncriptar)
	form.Append("", botonDesencriptar)

	return container.NewVBox(
		separadorTitulo,
		toolTips,
		form,
	)
}
