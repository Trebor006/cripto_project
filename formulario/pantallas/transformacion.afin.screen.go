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

func TransformacionAfinGenerarPantalla(w fyne.Window) fyne.CanvasObject {
	separadorTitulo := widgets.GenerarSeparadorTitulo("Transformacion Afin")
	transformacionAfin := encryptors.GetEncryptor(encryptors.TRANSFORMACION_AFIN)

	textoInicial := widget.NewMultiLineEntry()
	textoInicial.SetPlaceHolder("Texto a Encriptar/Desencriptar")
	textoInicial.SetMinRowsVisible(5)

	decimacion := widgets.NewNumEntry()
	decimacion.SetPlaceHolder("Decimacion")

	desplazamiento := widgets.NewNumEntry()
	desplazamiento.SetPlaceHolder("Desplazamiento")

	textoResultante := widget.NewMultiLineEntry()
	textoResultante.SetPlaceHolder("Resultado")
	textoResultante.SetMinRowsVisible(5)

	botonEncriptar := widget.NewButton("Encriptar", func() {
		widgets.LimpiarConsola()

		nroDecimacion, _ := strconv.Atoi(decimacion.Text)
		nroDesplazamiento, _ := strconv.Atoi(desplazamiento.Text)
		dataTransformacionAfin := data.Data{Message: textoInicial.Text, NroDecimacion: nroDecimacion, Desplazamiento: nroDesplazamiento}

		textocifradoportransformacionafin := transformacionAfin.Cypher(dataTransformacionAfin)
		textoResultante.SetText(util.Format(textocifradoportransformacionafin))
	})

	botonDesencriptar := widget.NewButton("Desencriptar", func() {
		widgets.LimpiarConsola()
		nroDecimacion, _ := strconv.Atoi(decimacion.Text)
		nroDesplazamiento, _ := strconv.Atoi(desplazamiento.Text)
		dataDescifradoTransformacionAfin := data.Data{EncryptedMessage: textoInicial.Text, NroDecimacion: nroDecimacion, Desplazamiento: nroDesplazamiento}

		textoDescrifrado := transformacionAfin.Decrypt(dataDescifradoTransformacionAfin)
		textoResultante.SetText(util.Format(textoDescrifrado))
	})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Texto Inicial", Widget: textoInicial, HintText: "Introduzca el texto a Encriptar/Desencriptar"},
			{Text: "Decimacion", Widget: decimacion, HintText: "Introduzca la decimacion"},
			{Text: "Desplazamiento", Widget: desplazamiento, HintText: "Introduzca el desplazamiento"},
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
