package logger

import (
	"fmt"
	"sync"
	"time"
)

type PlaintextFileLogger struct {
	//pureText means logger will print out only log text itself
	pureText       bool
	rotateFiles    bool
	rotateDuration time.Duration
	lastLog        time.Time
	llMutex        *sync.RWMutex
	lTypes         []LogType
	filepath       string
	file           IFile
	fileMutex      *sync.Mutex
}

// NewPlaintext returns logger capable of appending strings to text file.
//
// By passing IFile interface as f you can set the initial object to write logs to. Otherwise path & truncate
// will be used to create new file.
// Note: if rotateFiles > 0, file will be changed after this period of time any way
func NewPlaintext(path string, pureText bool, truncate bool, rotateFiles int, f IFile, lTypes ...LogType) (*PlaintextFileLogger, error) {
	var err error
	if f == nil {
		f, err = makeLogFile(path, truncate, "log")
		if err != nil {
			return nil, fmt.Errorf("[NewPlaintext] %w", err)
		}
	}

	return &PlaintextFileLogger{
		pureText:       pureText,
		rotateFiles:    rotateFiles > 0,
		rotateDuration: time.Minute * time.Duration(rotateFiles),
		lastLog:        time.Now(),
		lTypes:         lTypes,
		filepath:       path,
		file:           f,
		llMutex:        &sync.RWMutex{},
		fileMutex:      &sync.Mutex{},
	}, nil
}

// Log pushes event data into default output
func (l *PlaintextFileLogger) Log(e Event, timeFormat string) error {
	l.SetLastLog(time.Now())
	l.fileMutex.Lock()
	defer l.fileMutex.Unlock()
	//Make new file in case old one is... old
	if time.Since(l.LastLog()) > l.rotateDuration && l.rotateFiles {
		l.file.Close()
		f, err := makeLogFile(l.filepath, true, "log")
		if err != nil {
			return fmt.Errorf("[PlaintextFileLogger][Log] %w", err)
		}
		l.file = f
	}

	log := ""
	if l.pureText {
		log = FormatOutputPureText(e)
	} else {
		log = FormatOutput(e, timeFormat)
	}
	_, err := l.file.WriteString(log)
	if err != nil {
		return err
	}

	return nil
}

// Type returns set of types supported by the logger
func (l *PlaintextFileLogger) Type() []LogType { return l.lTypes }

// LastLog returns last time the logger was used
func (l *PlaintextFileLogger) LastLog() time.Time {
	l.llMutex.RLock()
	defer l.llMutex.RUnlock()

	return l.lastLog
}

// SetLastLog sets last time the logger was used
func (l *PlaintextFileLogger) SetLastLog(t time.Time) {
	l.llMutex.Lock()
	defer l.llMutex.Unlock()
	l.lastLog = t
}
