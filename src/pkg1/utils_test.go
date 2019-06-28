package pkg1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeBrokerStorageAuthenticationAccountKey(t *testing.T) {
	assertor := assert.New(t)

	assertor.Equal("Hello World", GetStuff(), "GetStuff()")
}
