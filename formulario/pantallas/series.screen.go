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

func SeriesGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Series")
	encryptor := encryptors.GetEncryptor(encryptors.SERIES)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()

		dataACifrar := data.Data{Message: textoInicial.Text}

		textoCifrado := encryptor.Cypher(dataACifrar)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()

		dataADescifrar := data.Data{EncryptedMessage: textoInicial.Text}

		textoDescifrado := encryptor.Decrypt(dataADescifrar)
		textoResultante.SetText(util.Format(textoDescifrado))
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
