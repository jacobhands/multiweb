package fileserver

// Getter gets a byte array containing a web resource.
type Getter interface {
	// Get retrieves a web resource.
	// path is the path to the resource.
	Get(path string) ([]byte, error)
}
