package service

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"strings"

	"github.com/google/uuid"
)

func formatJPEG(img image.Image) (io.Reader, error) {
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		err := jpeg.Encode(pw, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
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
