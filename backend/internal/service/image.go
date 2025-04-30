package service

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"strings"

	"github.com/google/uuid"
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

func createRandomImageFileName(extension string) string {
	extension = strings.TrimPrefix(extension, ".")

	return fmt.Sprintf("%s.%s", uuid.NewString(), extension)
}
