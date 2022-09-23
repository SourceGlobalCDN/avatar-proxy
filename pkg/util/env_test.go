package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvStr(t *testing.T) {
	asserts := assert.New(t)

	os.Setenv("TEST", "test")
	asserts.Equal("test", EnvStr("TEST", "default"))
	asserts.Equal("default", EnvStr("TEST1", "default"))

	os.Unsetenv("TEST")
	asserts.Equal("default", EnvStr("TEST", "default"))
}

func TestEnvInt(t *testing.T) {
	asserts := assert.New(t)

	os.Setenv("TEST", "1")
	asserts.Equal(1, EnvInt("TEST", 2))
	asserts.Equal(2, EnvInt("TEST1", 2))

	os.Unsetenv("TEST")
	asserts.Equal(2, EnvInt("TEST", 2))

	os.Setenv("TEST", "a")
	asserts.Equal(2, EnvInt("TEST", 2))

	os.Unsetenv("TEST")
	asserts.Equal(2, EnvInt("TEST", 2))
}

func TestEnvBool(t *testing.T) {
	asserts := assert.New(t)

	os.Setenv("TEST", "true")
	asserts.Equal(true, EnvBool("TEST", false))
	asserts.Equal(false, EnvBool("TEST1", false))

	os.Unsetenv("TEST")
	asserts.Equal(false, EnvBool("TEST", false))

	os.Setenv("TEST", "a")
	asserts.Equal(false, EnvBool("TEST", false))

	os.Unsetenv("TEST")
	asserts.Equal(false, EnvBool("TEST", false))
}
