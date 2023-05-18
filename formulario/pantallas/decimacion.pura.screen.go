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

func DecimacionPuraGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Decimacion Pura")
	metodoCifrado := encryptors.GetEncryptor(encryptors.DECIMACION_PURA)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	nroDecimacion := widgets.NewNumEntry()
	nroDecimacion.SetPlaceHolder("Nro Decimacion")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()

		decimacion, _ := strconv.Atoi(nroDecimacion.Text)
		dataToCypherDecimacionPura := data.Data{Message: textoInicial.Text, NroDecimacion: decimacion}

		textoCifradoPorDecimacionPura := metodoCifrado.Cypher(dataToCypherDecimacionPura)
		textoResultante.SetText(util.Format(textoCifradoPorDecimacionPura))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()

		decimacion, _ := strconv.Atoi(nroDecimacion.Text)
		dataToDecryptDecimacionPura := data.Data{EncryptedMessage: textoInicial.Text, NroDecimacion: decimacion}
		textoDesencriptadoPorDecimacionPura := metodoCifrado.Decrypt(dataToDecryptDecimacionPura)
		textoResultante.SetText(util.Format(textoDesencriptadoPorDecimacionPura))
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Nro Decimacion", Widget: nroDecimacion, HintText: "Introduzca la Decimacion"},
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
