package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type CSVFileLogger struct {
	lTypes         []LogType
	rotateFiles    bool
	rotateDuration time.Duration
	lastLog        time.Time
	llMutex        *sync.RWMutex
	filepath       string
	file           *os.File
	fileMutex      *sync.Mutex
}

// NewCSVtext returns logger capable of creating csv file records.
func NewCSVtext(path string, truncate bool, rotateFiles int, lTypes ...LogType) (*CSVFileLogger, error) {
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
		llMutex:        &sync.RWMutex{},
		fileMutex:      &sync.Mutex{},
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
func (l *CSVFileLogger) Log(e Event, timeFormat string) error {
	l.SetLastLog(time.Now())
	l.fileMutex.Lock()
	defer l.fileMutex.Unlock()
	//Make new file in case old one is... old
	if time.Since(l.LastLog()) > l.rotateDuration && l.rotateFiles {
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

	_, err := l.file.WriteString(FormatCSV(e, timeFormat))
	if err != nil {
		return fmt.Errorf("[CSVFileLogger] error making log entry: %w", err)
	}

	return nil
}

// Type returns set of types supported by the logger
func (l *CSVFileLogger) Type() []LogType { return l.lTypes }

// LastLog returns last time the logger was used
func (l *CSVFileLogger) LastLog() time.Time {
	l.llMutex.RLock()
	defer l.llMutex.RUnlock()
	return l.lastLog
}

// SetLastLog sets last time the logger was used
func (l *CSVFileLogger) SetLastLog(t time.Time) {
	l.llMutex.Lock()
	defer l.llMutex.Unlock()
	l.lastLog = t
}
