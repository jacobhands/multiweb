package fileserver

import "io/ioutil"

// FileServer serves web resources from files.
type FileServer struct {
	// ContentRoot is the folder containing the content
	ContentRoot string
	Files       Getter
	Cache       interface {
		Getter
		Setter
	}
	CacheLimit int
}

// NewFileServer creates a new instance with default cacheServ & fileServ Getters
func NewFileServer(contentRoot string) (fs FileServer) {
	fs = FileServer{
		ContentRoot: contentRoot,
		Files:       fileServ{},
		Cache:       cacheServ{},
	}
	if fs.ContentRoot[len(fs.ContentRoot)-1] != '/' {
		fs.ContentRoot += "/"
	}
	return
}

// Get retrieves content for the specifie path.
func (fs FileServer) Get(path string) ([]byte, error) {
	// Try retrieving from cache
	content, err := fs.Cache.Get(path)
	if err == nil {
		return content, nil
	}
	// Try getting from file
	content, err = fs.readFromFile(path)
	if err != nil {
		return nil, err
	}
	fs.Cache.Set(path, content)
	return content, nil
}

func (fs FileServer) readFromFile(path string) ([]byte, error) {
	contentPath := fs.ContentRoot + path
	if contentPath[len(contentPath)-1] == '/' {
		contentPath += "index.html"
	}
	content, err := ioutil.ReadFile(contentPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
