package canvas

import (
	"github.com/tdewolff/canvas"
	"image/png"
	"os"
)

func drawImage(ctx *canvas.Context, path string, x float64, y float64, dpmm float64) {
	image, err := os.Open(path)
	if err != nil {
		println(err)
	}
	pngImage, decodeErr := png.Decode(image)
	if decodeErr != nil {
		panic(err)
	}
	defer image.Close()
	ctx.DrawImage(x, y, pngImage, canvas.DPMM(dpmm))
}
