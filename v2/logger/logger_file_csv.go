package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/lazybark/lazyevent/v2/events"
)

type CSVFileLogger struct {
	lTypes         []events.LogType
	rotateFiles    bool
	rotateDuration time.Duration
	lastLog        time.Time
	filepath       string
	file           *os.File
}

// NewCSVtext returns logger capable of creating csv file records.
func NewCSVtext(path string, truncate bool, rotateFiles int, lTypes ...events.LogType) (*CSVFileLogger, error) {
	f, err := makeLogFile(path, truncate, "csv")
	if err != nil {
		return nil, fmt.Errorf("[NewCSVtext] %w", err)
	}

	csv := &CSVFileLogger{
		lTypes:         lTypes,
		rotateFiles:    rotateFiles > 0,
		rotateDuration: time.Minute * time.Duration(rotateFiles),
		lastLog:        time.Now(),
		file:           f,
		filepath:       path,
	}

	if truncate {
		_, err := f.WriteString(CSVHead)
		if err != nil {
			return nil, fmt.Errorf("[CSVFileLogger] error making CSV head entry: %w", err)
		}
	}

	return csv, nil
}

// Log pushes event data into default output
func (l *CSVFileLogger) Log(e events.Event, timeFormat string) error {
	//Make new file in case old one is... old
	if time.Since(l.lastLog) > l.rotateDuration && l.rotateFiles {
		l.file.Close()
		f, err := makeLogFile(l.filepath, true, "csv")
		if err != nil {
			return fmt.Errorf("[CSVFileLogger][Log] %w", err)
		}
		l.file = f

		_, err = f.WriteString(CSVHead)
		if err != nil {
			return fmt.Errorf("[CSVFileLogger][Log] error making CSV head entry: %w", err)
		}
	}

	l.lastLog = time.Now()
	_, err := l.file.WriteString(FormatCSV(e, timeFormat))
	if err != nil {
		return fmt.Errorf("[CSVFileLogger] error making log entry: %w", err)
	}

	return nil
}

// Type returns set of types supported by the logger
func (l *CSVFileLogger) Type() []events.LogType { return l.lTypes }
