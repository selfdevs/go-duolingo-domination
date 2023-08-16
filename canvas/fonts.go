package canvas

import "github.com/tdewolff/canvas"

func loadFontFamily() *canvas.FontFamily {
	feather := canvas.NewFontFamily("feather")
	if err := feather.LoadFontFile("assets/feather.ttf", canvas.FontRegular); err != nil {
		panic(err)
	}
	return feather
}

func createTitleFontFace() *canvas.FontFace {
	feather := loadFontFamily()
	return feather.Face(150, green, canvas.FontNormal, canvas.FontNormal)
}

func createEntryFontFace() *canvas.FontFace {
	feather := loadFontFamily()
	return feather.Face(100, canvas.White, canvas.FontNormal, canvas.FontNormal)
}
