package fileserver

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FileServerSuite struct {
	suite.Suite
	fs FileServer
}

func TestFileServerSuite(t *testing.T) {
	suite.Run(t, new(FileServerSuite))
}

func (s *FileServerSuite) SetupTest() {
	s.fs = FileServer{
		ContentRoot: "/",
		Getters: []Getter{
			&mockCacheServ{},
			&mockFileServ{},
		},
	}
}

// START TESTS

func (s FileServerSuite) TestGet() {
	// content, err := s.fs.Get("/test")

	mockFS := new(mockFileServ)
	mockCS := new(mockCacheServ)
	s.fs.Getters = []Getter{
		mockFS,
		mockCS,
	}
	mockFS.On("Get", "/abc/").Return([]byte{}, nil)

}
