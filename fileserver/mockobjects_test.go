package fileserver

// import "github.com/stretchr/testify/mock"

// // START mockCacheServ

// type mockCacheServ struct {
// 	mock.Mock
// }

// func (m *mockCacheServ) Get(path string) ([]byte, error) {
// 	args := m.Called(path)
// 	return args.Get(0).([]byte), args.Error(1)
// }

// // END mockCacheServ

// // START mockFileServ

// type mockFileServ struct {
// 	mock.Mock
// }

// func (m *mockFileServ) Get(path string) ([]byte, error) {
// 	args := m.Called(path)
// 	return args.Get(0).([]byte), args.Error(1)
// }

// // END mockFileServ
