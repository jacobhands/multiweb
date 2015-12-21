package fileserver

// FileServer serves web resources from files.
type FileServer struct {
	// ContentRoot is the folder containing the content
	ContentRoot string
	Getters     []Getter
}

// NewFileServer creates a new instance with default cacheServ & fileServ Getters
func NewFileServer(contentRoot string) (fs FileServer) {
	fs = FileServer{
		ContentRoot: contentRoot,
		Getters: []Getter{
			cacheServ{},
			fileServ{},
		},
	}
	return
}

// Get retrieves content for the specifie path.
func (fs FileServer) Get(path string) (content []byte, err error) {
	for i, getter := range fs.Getters {
		content, err = getter.Get(path)
		if err == nil {
			return content, nil
		}
		if i == len(fs.Getters)-1 { // All getters returned errors
			return nil, ErrInvalidContentPath{path}
		}
	}
	return
}
