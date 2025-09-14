package psql

import "fmt"

func NewQueryBuildError(queryName string, err error) error {
	return fmt.Errorf("Could not build query %q: %v", queryName, err.Error())
}
