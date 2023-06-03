package logservice

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

func NewLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}

	logger := zerolog.
		New(output).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Logger()

	return &logger
}
