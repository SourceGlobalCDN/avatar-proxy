package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	asserts := assert.New(t)
	asserts.NotNil(Log())
}

func TestNewLogger(t *testing.T) {
	asserts := assert.New(t)
	logger := NewLogger()
	asserts.NotNil(logger)
	asserts.Equal(Level, logger.Level)
	asserts.NotNil(logger.Formatter)
	asserts.NotNil(logger.Out)
}
