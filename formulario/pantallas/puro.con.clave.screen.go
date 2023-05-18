package pantallas

import (
	"cripto_project/main/data"
	"cripto_project/main/encryptors"
	"cripto_project/main/formulario"
)

func PuroConClaveGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := formulario.GenerarSeparadorTitulo("Cifrado Cesar con Clave")
	encryptor := encryptors.GetEncryptor(encryptors.PURO_CON_CLAVE)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		data := data.Data{}
		data.Message = textoInicial.Text
		textoCifrado := encryptor.Cypher(data)
		textoResultante.SetText(textoCifrado)
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		data := data.Data{}
		data.EncryptedMessage = textoInicial.Text
		textoCifrado := encryptor.Decrypt(data)
		textoResultante.SetText(textoCifrado)
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Resultado", Widget: textoResultante, HintText: "Aqui se mostrara el resultado"},
		},
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

	return container.NewVBox(
		separadorTitulo,
		//toolTips,
		form,
	)
}
