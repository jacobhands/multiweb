package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGET(t *testing.T) {

}

func TestGetSubDomain(t *testing.T) {
	assert.Equal(t, "aaa.bbb.ccc", getSubDomain("aaa.bbb.ccc.ddd.fff.ggg", "ddd.fff.ggg"))
}
