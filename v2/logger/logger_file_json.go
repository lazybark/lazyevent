package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/lazybark/lazyevent/v2/events"
)

type JSONFileLogger struct {
	lTypes         []events.LogType
	rotateFiles    bool
	rotateDuration time.Duration
	lastLog        time.Time
	filepath       string
	file           *os.File
}

// NewJSONtext returns logger capable of creating json-encoded records in text file.
//
// Note: current realization fills file with json-objects, but does not create array of records ([...]).
// It makes log writing faster, but the only way to read valid JSON from the file is string by string,
// not whole file at once.
func NewJSONtext(path string, truncate bool, rotateFiles int, lTypes ...events.LogType) (*JSONFileLogger, error) {
	f, err := makeLogFile(path, truncate, "json")
	if err != nil {
		return nil, fmt.Errorf("[NewJSONtext] %w", err)
	}

	return &JSONFileLogger{
		lTypes:         lTypes,
		rotateFiles:    rotateFiles > 0,
		lastLog:        time.Now(),
		rotateDuration: time.Minute * time.Duration(rotateFiles),
		filepath:       path,
		file:           f,
	}, nil
}

// Log pushes event data into default output
func (l *JSONFileLogger) Log(e events.Event, timeFormat string) error {
	//Make new file in case old one is... old
	if time.Since(l.lastLog) > l.rotateDuration && l.rotateFiles {
		l.file.Close()
		f, err := makeLogFile(l.filepath, true, "json")
		if err != nil {
			return fmt.Errorf("[JSONFileLogger][Log] %w", err)
		}
		l.file = f
	}

	l.lastLog = time.Now()
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
