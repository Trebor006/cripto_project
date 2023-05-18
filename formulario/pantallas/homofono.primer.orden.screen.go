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

func HomofonoPrimerOrdenGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Primer Orden")
	homofonoPrimerOrden := encryptors.GetEncryptor(encryptors.PRIMER_ORDEN)
	random := util.ObtenerGeneradorRandom()
	diccionario := util.InicializarDiccionarioPrimerOrden(100, random)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()
		dataToCypherByPrimerOrden := data.Data{Message: textoInicial.Text, Diccionario: diccionario, RandomGenerator: random}
		encryptedMessageByPrimerOrden := homofonoPrimerOrden.Cypher(dataToCypherByPrimerOrden)
		textoResultante.SetText(encryptedMessageByPrimerOrden)
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()
		dataToDecryptByPrimerOrden := data.Data{EncryptedMessage: textoInicial.Text, Diccionario: diccionario, RandomGenerator: random}
		textoDesencriptado := homofonoPrimerOrden.Decrypt(dataToDecryptByPrimerOrden)

		textoResultante.SetText(util.Format(textoDesencriptado))
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