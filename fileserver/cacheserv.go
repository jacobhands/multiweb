package fileserver

type cacheServ struct {
	pages map[string][]byte
}

func (c cacheServ) Get(path string) (content []byte, err error) {
	content = c.pages[path]
	if content == nil || len(content) == 0 {
		return nil, ErrInvalidContentPath{path}
	}
	return content, nil
}

func (c cacheServ) Set(path string, content []byte) (err error) {

	return
}
