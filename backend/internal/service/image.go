package service

import (
	"image"
	"image/jpeg"
	"io"

	"github.com/nfnt/resize"
)

func resizeImage(img image.Image, width, height uint) (io.Reader, error) {
	resizedImage := resize.Resize(width, height, img, resize.Lanczos2)

	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		err := jpeg.Encode(pw, resizedImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			pw.CloseWithError(err)
		}
	}()

	return pr, nil
}
