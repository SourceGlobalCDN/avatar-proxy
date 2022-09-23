package log

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFormat(t *testing.T) {
	asserts := assert.New(t)

	logger := logrus.New()
	logger.SetFormatter(NewFormatter())
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	bytes, err := logger.Formatter.Format(logger.WithFields(logrus.Fields{"test": "test"}))
	asserts.NoError(err)
	asserts.NotEmpty(bytes)
}

func BenchmarkTestFormatter_Format(b *testing.B) {
	logger := logrus.New()
	logger.SetFormatter(NewFormatter())
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Formatter.Format(logger.WithFields(logrus.Fields{"test": "test"}))
	}
	b.StopTimer()
}
