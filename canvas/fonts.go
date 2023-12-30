package canvas

import (
	"github.com/tdewolff/canvas"
	"io"
)

func loadFontFamily() *canvas.FontFamily {
	feather := canvas.NewFontFamily("feather")
	font := downloadImage("feather.ttf")
	data, err := io.ReadAll(font)
	if err != nil {
		panic(err)
	}
	if err := feather.LoadFont(data, 0, canvas.FontRegular); err != nil {
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
