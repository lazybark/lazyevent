package logger

import (
	"fmt"
	"os"

	"github.com/lazybark/go-helpers/fsw"
	"github.com/lazybark/lazyevent/v2/events"
)

type CSVFileLogger struct {
	lTypes []events.LogType
	file   *os.File
}

// NewCSVtext returns logger capable of creating csv file records.
func NewCSVtext(path string, truncate bool, lTypes ...events.LogType) (*CSVFileLogger, error) {
	f, err := fsw.MakePathToFile(path, truncate)
	if err != nil {
		return nil, fmt.Errorf("[CSVFileLogger] error making log path: %w", err)
	}

	csv := &CSVFileLogger{lTypes: lTypes, file: f}

	if truncate {
		_, err := f.WriteString(CSVHead)
		if err != nil {
			return nil, fmt.Errorf("[CSVFileLogger] error making CSV head entry: %w", err)
		}
	}

	return csv, nil
}

// Log pushes event data into default output
func (l CSVFileLogger) Log(e events.Event, timeFormat string) error {
	_, err := l.file.WriteString(FormatCSV(e, timeFormat))
	if err != nil {
		return fmt.Errorf("[CSVFileLogger] error making log entry: %w", err)
	}

	return nil
}

// Type returns set of types supported by the logger
func (l CSVFileLogger) Type() []events.LogType { return l.lTypes }
