package logger

import (
	"fmt"
	"os"

	"github.com/lazybark/go-helpers/fsw"
	"github.com/lazybark/lazyevent/v2/events"
)

type JSONFileLogger struct {
	lTypes []events.LogType
	file   *os.File
}

// NewJSONtext returns logger capable of creating json-encoded records in text file.
//
// Note: current realization fills file with json-objects, but does not create array of records ([...]).
// It makes log writing faster, but the only way to read valid JSON from the file is string by string,
// not whole file at once.
func NewJSONtext(path string, truncate bool, lTypes ...events.LogType) (*JSONFileLogger, error) {
	f, err := fsw.MakePathToFile(path, truncate)
	if err != nil {
		return nil, fmt.Errorf("[JSONFileLogger] error making log path: %w", err)
	}

	return &JSONFileLogger{lTypes: lTypes, file: f}, nil
}

// Log pushes event data into default output
func (l JSONFileLogger) Log(e events.Event, timeFormat string) error {
	js, err := FormatJSON(e, timeFormat)
	if err != nil {
		return fmt.Errorf("[JSONFileLogger] error formatting event to JSON: %w", err)
	}
	js = append(js, '\n')
	_, err = l.file.Write(js)
	if err != nil {
		return fmt.Errorf("[JSONFileLogger] error making log entry: %w", err)
	}

	return nil
}

// Type returns set of types supported by the logger
func (l JSONFileLogger) Type() []events.LogType { return l.lTypes }
