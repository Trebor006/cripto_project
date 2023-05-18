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

func GruposGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Grupos")
	metodoCifrado := encryptors.GetEncryptor(encryptors.GRUPOS)

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

		dataAEncriptar := data.Data{Message: textoInicial.Text, Clave: clave.Text}

		textoCifrado := metodoCifrado.Cypher(dataAEncriptar)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()
		dataADesencriptar := data.Data{EncryptedMessage: textoInicial.Text, Clave: clave.Text}

		textoClaro := metodoCifrado.Decrypt(dataADesencriptar)
		textoResultante.SetText(util.Format(textoClaro))
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

//func columnasScreen(w fyne.Window) fyne.CanvasObject {
//	separadorTitulo := widgets.GenerarSeparadorTitulo("Columnas")
//	encryptor := encryptors.GetEncryptor(encryptors.COLUMNAS)
//
//	textoInicial := widget.NewMultiLineEntry()
//	textoInicial.SetPlaceHolder("Texto a Inicial")
//	textoInicial.SetMinRowsVisible(5)
//
//	nroColumnas := widgets.NewNumEntry()
//	nroColumnas.SetPlaceHolder("Nro columnas")
//
//	textoResultante := widget.NewMultiLineEntry()
//	textoResultante.SetPlaceHolder("Resultado")
//	textoResultante.SetMinRowsVisible(5)
//
//	botonEncriptar := widget.NewButton("Encriptar", func() {
//		data := data.Data{}
//		data.Message = textoInicial.Text
//		data.NroColumns, _ = strconv.Atoi(nroColumnas.Text)
//
//		textoCifrado := encryptor.Cypher(data)
//		textoResultante.SetText(util.Format(textoCifrado))
//	})
//
//	botonDesencriptar := widget.NewButton("Desencriptar", func() {
//		data := data.Data{}
//		data.EncryptedMessage = textoInicial.Text
//		data.NroColumns, _ = strconv.Atoi(nroColumnas.Text)
//
//		textoCifrado := encryptor.Decrypt(data)
//		textoResultante.SetText(textoCifrado)
//	})
//
//	return container.NewVBox(
//		separadorTitulo,
//		textoInicial,
//		nroColumnas,
//		textoResultante,
//		botonEncriptar,
//		botonDesencriptar,
//	)
//}
