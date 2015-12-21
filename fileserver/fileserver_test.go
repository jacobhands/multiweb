package fileserver

import "testing"

var fs int

func init() {
	fs = 0
}
func TestGet(t *testing.T) {
	fs++
	println(fs)
}
func TestOther(t *testing.T) {
	fs++
	println(fs)
}
