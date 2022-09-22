package serializer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildResponse(t *testing.T) {
	asserts := assert.New(t)

	{
		response := BuildResponse(200, "OK", "test")
		asserts.Equal(200, response.Code)
		asserts.Equal("OK", response.Msg)
		asserts.Equal("test", response.Data)
	}

	{
		response := BuildResponse(200, "OK", nil)
		asserts.Equal(200, response.Code)
		asserts.Equal("OK", response.Msg)
		asserts.Nil(response.Data)
	}
}
