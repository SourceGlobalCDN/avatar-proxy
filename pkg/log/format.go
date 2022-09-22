package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type SprintfFunc func(format string, a ...interface{}) string

type Formatter struct {
	colors map[logrus.Level]SprintfFunc
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	colorFunc := f.colors[entry.Level]

	level := entry.Level.String()
	if entry.Data["level"] != nil {
		level = entry.Data["level"].(string)
	}

	switch entry.Logger.Level {
	case logrus.DebugLevel:
		level = colorFunc("%-11s", "["+strings.ToUpper(level)+"]")
	default:
		level = colorFunc("%-7s", "["+strings.ToUpper(level)+"]")
	}

	return []byte(fmt.Sprintf(
		"%s %s | %s | %s\n",
		level,
		color.New(color.FgHiMagenta).Sprint(os.Getpid()),
		entry.Time.Format("2006-01-02 15:04:05.000"),
		entry.Message,
	)), nil
}

func NewFormatter() *Formatter {
	return &Formatter{
		colors: map[logrus.Level]SprintfFunc{
			logrus.WarnLevel:  color.New(color.FgYellow).Add(color.Bold).SprintfFunc(),
			logrus.PanicLevel: color.New(color.BgHiRed).Add(color.Bold).SprintfFunc(),
			logrus.FatalLevel: color.New(color.BgRed).Add(color.Bold).SprintfFunc(),
			logrus.ErrorLevel: color.New(color.FgRed).Add(color.Bold).SprintfFunc(),
			logrus.InfoLevel:  color.New(color.FgCyan).Add(color.Bold).SprintfFunc(),
			logrus.DebugLevel: color.New(color.FgWhite).Add(color.Bold).SprintfFunc(),
		},
	}
}
