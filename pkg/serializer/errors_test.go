package serializer

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHttpError(t *testing.T) {
	asserts := assert.New(t)

	{
		err := HttpError(200)
		asserts.Equal(200, err.Code)
		asserts.Equal(http.StatusText(200), err.Msg)
		asserts.Nil(err.Data)
	}

	{
		err := HttpError(404)
		asserts.Equal(404, err.Code)
		asserts.Equal(http.StatusText(404), err.Msg)
		asserts.Nil(err.Data)
	}
}
