package pantallas

import (
	"cripto_project/main/data"
	"cripto_project/main/encryptors"
	"cripto_project/main/formulario"
	"cripto_project/main/util"
	"strconv"
)

func ZigZagGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := formulario.GenerarSeparadorTitulo("ZigZag")
	encryptor := encryptors.GetEncryptor(encryptors.ZIG_ZAG)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	nroRails := formulario.NewNumEntry()
	nroRails.SetPlaceHolder("Nro Rails")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		data := data.Data{}
		data.Message = textoInicial.Text
		data.NroRails, _ = strconv.Atoi(nroRails.Text)

		textoCifrado := encryptor.Cypher(data)
		textoResultante.SetText(util.Format(textoCifrado))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		data := data.Data{}
		data.EncryptedMessage = textoInicial.Text
		data.NroRails, _ = strconv.Atoi(nroRails.Text)

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
			{Text: "Nro", Widget: nroRails, HintText: "Introduzca el Nro"},
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

//func columnasScreen(w fyne.Window) fyne.CanvasObject {
//	separadorTitulo := formulario.GenerarSeparadorTitulo("Columnas")
//	encryptor := encryptors.GetEncryptor(encryptors.COLUMNAS)
//
//	textoInicial := widget.NewMultiLineEntry()
//	textoInicial.SetPlaceHolder("Texto a Inicial")
//	textoInicial.SetMinRowsVisible(5)
//
//	nroColumnas := formulario.NewNumEntry()
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
