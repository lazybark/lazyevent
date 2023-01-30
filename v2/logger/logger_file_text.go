package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/lazybark/lazyevent/v2/events"
)

type PlaintextFileLogger struct {
	//pureText means logger will print out only log text itself
	pureText       bool
	rotateFiles    bool
	rotateDuration time.Duration
	lastLog        time.Time
	lTypes         []events.LogType
	filepath       string
	file           *os.File
}

// NewPlaintext returns logger capable of appending strings to text file
func NewPlaintext(path string, pureText bool, truncate bool, rotateFiles int, lTypes ...events.LogType) (*PlaintextFileLogger, error) {
	f, err := makeLogFile(path, truncate, "log")
	if err != nil {
		return nil, fmt.Errorf("[NewPlaintext] %w", err)
	}

	return &PlaintextFileLogger{
		pureText:       pureText,
		rotateFiles:    rotateFiles > 0,
		rotateDuration: time.Minute * time.Duration(rotateFiles),
		lastLog:        time.Now(),
		lTypes:         lTypes,
		filepath:       path,
		file:           f,
	}, nil
}

// Log pushes event data into default output
func (l PlaintextFileLogger) Log(e events.Event, timeFormat string) error {
	//Make new file in case old one is... old
	if time.Since(l.lastLog) > l.rotateDuration && l.rotateFiles {
		l.file.Close()
		f, err := makeLogFile(l.filepath, true, "log")
		if err != nil {
			return fmt.Errorf("[PlaintextFileLogger][Log] %w", err)
		}
		l.file = f
	}

	log := ""
	if l.pureText {
		log = FormatPureText(e)
	} else {
		log = Format(e, timeFormat)
	}
	_, err := l.file.WriteString(log)
	if err != nil {
		return err
	}

	return nil
}

// Type returns set of types supported by the logger
func (l PlaintextFileLogger) Type() []events.LogType { return l.lTypes }
