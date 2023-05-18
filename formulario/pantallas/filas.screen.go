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

func FilasGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Filas")
	encryptor := encryptors.GetEncryptor(encryptors.FILAS)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	nroFilas := widgets.NewNumEntry()
	nroFilas.SetPlaceHolder("Nro Filas")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()

		nroFil, _ := strconv.Atoi(nroFilas.Text)
		dataACifrar := data.Data{Message: textoInicial.Text, NroFilas: nroFil}
		textoCifrado := encryptor.Cypher(dataACifrar)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()

		nroFil, _ := strconv.Atoi(nroFilas.Text)
		dataACifrar := data.Data{EncryptedMessage: textoInicial.Text, NroFilas: nroFil}
		textoDescifrado := encryptor.Decrypt(dataACifrar)
		textoResultante.SetText(util.Format(textoDescifrado))
	})

	//buttonsLayout := widget.FormItem{
	//	Text: "", Widget: botonEncriptar, HintText: "Introduzca el texto a Encriptar/Desencriptar",
	//	//{Text: "", Widget: botonDesencriptar, HintText: "Aqui se mostrara el resultado"},
	//}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Nro de Filas", Widget: nroFilas, HintText: "Introduzca el Nro de Filas"},
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
