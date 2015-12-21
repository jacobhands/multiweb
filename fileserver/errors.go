package fileserver

import "fmt"

// ErrInvalidContentPath occurs when no content is found for the given path.
type ErrInvalidContentPath struct {
	path string
}

func (e ErrInvalidContentPath) Error() string {
	return fmt.Sprintf("Invalid content path: %s", e.path)
}
