package object

import (
	"context"
	"io"
)

type Store interface {
	// Uploads the input reader stream to the specified path, returning the
	// location of where the object was uploaded to.
	Upload(context.Context, string, io.Reader) (string, error)

	// Gets the object body given its path.
	Get(context.Context, string) (io.ReadCloser, error)

	//Deletes an object given its path.
	Delete(context.Context, string) error
}
