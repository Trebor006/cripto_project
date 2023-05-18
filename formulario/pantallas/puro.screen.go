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

func PuroGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Cifrado Cesar")
	metodoCifrado := encryptors.GetEncryptor(encryptors.PURO)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()

		dataAEncriptar := data.Data{Message: textoInicial.Text}

		textoCifrado := metodoCifrado.Cypher(dataAEncriptar)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()

		dataEncriptada := data.Data{EncryptedMessage: textoInicial.Text}

		textoClaro := metodoCifrado.Decrypt(dataEncriptada)
		textoResultante.SetText(util.Format(textoClaro))
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
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
