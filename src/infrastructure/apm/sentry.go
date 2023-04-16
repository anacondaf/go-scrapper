package apm

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/kainguyen/go-scrapper/src/config"
	"time"
)

func SetupSentry(config *config.Config) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.Sentry.DSN,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("sentry.Init: %s", err))
	}

	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")

	return nil
}
