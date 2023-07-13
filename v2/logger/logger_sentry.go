package logger

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/lazybark/lazyevent/v2/events"
)

// SentryLogger uses Go-native Sentry package to log events. If you need to adjust parameters, use functions defined by
// github.com/getsentry/sentry-go (ConfigureScope, for example) after logger.NewSentry has been called.
type SentryLogger struct {
	lTypes []events.LogType
	dsn    string
}

// NewSentry returns SentryLogger and configures Go-native Sentry package to use specified parameters
func NewSentry(dsn, env, rel string, debug bool, tsr float64, lTypes ...events.LogType) (*SentryLogger, error) {
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: env,
		Release:     rel,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: debug,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: tsr,
	})
	if err != nil {
		return nil, fmt.Errorf("[NewSentry] %w", err)
	}

	return &SentryLogger{dsn: dsn, lTypes: lTypes}, nil

}

func (l *SentryLogger) Log(e events.Event, timeFormat string) error {
	//timeFormat is unused here and is left to compatibility with the interface

	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage(FormatPureText(e))

	return nil
}

// Type returns set of types supported by the logger
func (l *SentryLogger) Type() []events.LogType { return l.lTypes }
