package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type JSONFileLogger struct {
	lTypes         []LogType
	rotateFiles    bool
	rotateDuration time.Duration
	lastLog        time.Time
	llMutex        *sync.RWMutex
	filepath       string
	file           *os.File
	fileMutex      *sync.Mutex
	firstLine      bool
}

// NewJSONtext returns logger capable of creating json-encoded records in text file.
//
// Note: current realization fills file with json-objects, but does not create array of records ([...]).
// It makes log writing faster, but the only way to read valid JSON from the file is to frame its contents by [] before parsing.
func NewJSONtext(path string, truncate bool, rotateFiles int, lTypes ...LogType) (*JSONFileLogger, error) {
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
		llMutex:        &sync.RWMutex{},
		fileMutex:      &sync.Mutex{},
		firstLine:      true,
	}, nil
}

// Log pushes event data into default output
func (l *JSONFileLogger) Log(e Event, timeFormat string) error {
	l.SetLastLog(time.Now())
	l.fileMutex.Lock()
	defer l.fileMutex.Unlock()
	//Make new file in case old one is... old
	if time.Since(l.LastLog()) > l.rotateDuration && l.rotateFiles {
		l.file.Close()
		f, err := makeLogFile(l.filepath, true, "json")
		if err != nil {
			return fmt.Errorf("[JSONFileLogger][Log] %w", err)
		}
		l.file = f
		l.firstLine = true
	}

	var text []byte
	if !l.firstLine {
		text = append(text, ',', '\n')
	}

	js, err := FormatJSON(e, timeFormat)
	if err != nil {
		return fmt.Errorf("[JSONFileLogger] error formatting event to JSON: %w", err)
	}
	text = append(text, js...)
	_, err = l.file.Write(text)
	if err != nil {
		return fmt.Errorf("[JSONFileLogger] error making log entry: %w", err)
	}
	l.firstLine = false

	return nil
}

// Type returns set of types supported by the logger
func (l *JSONFileLogger) Type() []LogType { return l.lTypes }

// LastLog returns last time the logger was used
func (l *JSONFileLogger) LastLog() time.Time {
	l.llMutex.RLock()
	defer l.llMutex.RUnlock()
	return l.lastLog
}

// SetLastLog sets last time the logger was used
func (l *JSONFileLogger) SetLastLog(t time.Time) {
	l.llMutex.Lock()
	defer l.llMutex.Unlock()
	l.lastLog = t
}
