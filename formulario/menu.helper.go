package formulario

import "cripto_project/main/formulario/pantallas"

func GenerarMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	cifradosDeDesplazamiento := fyne.NewMenuItem("Cifra de Desplazamiento", nil)
	cifradosDeDesplazamiento.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Puro", func() {
			w.SetContent(pantallas.PuroGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Puro con Palabra Clave", func() {
			w.SetContent(pantallas.PuroConClaveGenerarPantalla(w))
		}),
	)

	cifradosPorTraspocicion := fyne.NewMenuItem("Cifra por Transposición", nil)
	cifradosPorTraspocicion.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Grupos", func() {
			w.SetContent(pantallas.GruposGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Series", func() {
			w.SetContent(pantallas.SeriesGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Columnas", func() {
			w.SetContent(pantallas.ColumnasGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Filas", func() {
			w.SetContent(pantallas.FilasGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Zig-Zag", func() {
			w.SetContent(pantallas.ZigZagGenerarPantalla(w))
		}),
	)

	cifradosPorSustitucion := fyne.NewMenuItem("Cifra por Sustitución", nil)
	sustitucionMonoAlfabetica := fyne.NewMenuItem("Mono alfabética", nil)

	sustitucionMonoAlfabetica.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Decimación Pura", func() {
			w.SetContent(pantallas.DecimacionPuraGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Transformación Afín", func() {
			w.SetContent(pantallas.TransformacionAfinGenerarPantalla(w))
		}),
	)

	sustitucionPoliAlfabetica := fyne.NewMenuItem("Poli alfabética", nil)
	sustitucionPoliAlfabetica.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Homófonos de Primer Orden", func() {
			w.SetContent(pantallas.HomofonoPrimerOrdenGenerarPantalla(w))
		}),
		fyne.NewMenuItem("Homófonos de Orden Mayor", func() {
			w.SetContent(pantallas.HomofonoOrdenMayorGenerarPantalla(w))
		}),
	)

	cifradosPorSustitucion.ChildMenu = fyne.NewMenu("",
		sustitucionMonoAlfabetica,
		fyne.NewMenuItemSeparator(),
		sustitucionPoliAlfabetica,
	)

	cifradosPorSustitucionMonografica := fyne.NewMenuItem("Sustitucion Monogramica Polialfabeto", nil)
	cifradosPorSustitucionMonografica.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Polialfabeticos periodicos", func() {
			w.SetContent(pantallas.PolialfabeticosPeriodicosGenerarPantalla(w))
		}),
	)

	metodosDeCifrado := fyne.NewMenu("Metodos de Cifrado",
		cifradosDeDesplazamiento,
		fyne.NewMenuItemSeparator(),
		cifradosPorTraspocicion,
		fyne.NewMenuItemSeparator(),
		cifradosPorSustitucion,
		fyne.NewMenuItemSeparator(),
		cifradosPorSustitucionMonografica,
	)

	main := fyne.NewMainMenu(
		metodosDeCifrado,
	)

	return main
}
