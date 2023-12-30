package canvas

import (
	"github.com/tdewolff/canvas"
	"image/png"
	"io"
	"log"
	"net/http"
)

func downloadImage(filename string) io.ReadCloser {
	res, err := http.Get("https://duolingo.s3.fr-par.scw.cloud/" + filename)

	if err != nil {
		log.Fatalf("Error while downloading %s: %s", filename, err)
	}

	return res.Body
}

func drawImage(ctx *canvas.Context, path string, x float64, y float64, dpmm float64) {
	image := downloadImage(path)
	pngImage, decodeErr := png.Decode(image)
	if decodeErr != nil {
		log.Fatalf("Error while decoding %s: %s", path, decodeErr)
	}
	defer image.Close()
	ctx.DrawImage(x, y, pngImage, canvas.DPMM(dpmm))
}
