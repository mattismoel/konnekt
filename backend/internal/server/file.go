package server

import (
	"fmt"
	"path"

	"github.com/google/uuid"
)

func generateRandomFileName(fileName string) string {
	ext := path.Ext(fileName)
	return fmt.Sprintf("%s%s", uuid.NewString(), ext)
}
