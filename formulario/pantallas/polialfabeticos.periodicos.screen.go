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

func PolialfabeticosPeriodicosGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Polialfabetico Periodicos - Vigenere")
	encryptor := encryptors.GetEncryptor(encryptors.POLIALFABETICO_PERIODICO)

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

		dataCifrado := data.Data{Message: textoInicial.Text, Clave: clave.Text}

		textoCifrado := encryptor.Cypher(dataCifrado)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()

		dataDescifrar := data.Data{EncryptedMessage: textoInicial.Text, Clave: clave.Text}

		textoDescrifrado := encryptor.Decrypt(dataDescifrar)
		textoResultante.SetText(util.Format(textoDescrifrado))
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Clave", Widget: clave, HintText: "Introduzca la clave"},
			{Text: "Resultado", Widget: textoResultante, HintText: "Aqui se mostrara el resultado"},
		},
	}

	toolTips := widgets.GenerateToolTipCopyButton(textoInicial, textoResultante)

	form.Append("", botonEncriptar)
	form.Append("", botonDesencriptar)

	return container.NewVBox(
		separadorTitulo,
		toolTips,
		form,
	)
}
