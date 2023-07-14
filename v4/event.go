package logger

import (
	"time"
)

// Event represents any event that can occur during runtime.
// All Event fields are public and it's perfectly safe to modify
// Event without using API.
//
// Event can be reused after being logged: new calling to Log()
// will log the event with new ID.
type Event struct {
	ID     string
	Level  Level
	Type   LogType
	Source Source
	Time   time.Time
	Text   string
	Format Format

	//TimeFixed should be set to true if the app must log same event instance without updating
	//event's Time value. E.g. for making several records with different text, but for same time.
	TimeFixed bool
}

func (e Event) SetText(t string) Event {
	e.Text = t
	return e
}

// FixTime marks event time as fixed, so logger will NOT use time.Now() value
// each time event is logged
func (e Event) FixTime() Event {
	e.TimeFixed = true
	return e
}

// UnFixTime marks event time as non-fixed, so logger WILL use time.Now() value
// each time event is logged
func (e Event) UnFixTime() Event {
	e.TimeFixed = false
	return e
}

// Src returns event with source = s
func (e Event) Src(s Source) Event {
	e.Source = s
	return e
}

// FlushID cleans event ID
func (e Event) FlushID() Event {
	e.ID = ""
	return e
}

// SetID sets event ID. It will be used by LogProcessor
func (e Event) SetID(id string) Event {
	e.ID = id
	return e
}

// Empty returns event with INFO level, no text and default type, source and format
func Empty() Event {
	return Event{Text: "", Time: time.Now(), Level: INFO, Type: Any, Source: EvsEmpty, Format: None}
}

// Info returns event with INFO level and default type, source and format
func Info(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: INFO, Type: Any, Source: EvsEmpty, Format: None}
}

// Note returns event with NOTE level and default type, source and format
func Note(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: NOTE, Type: Any, Source: EvsEmpty, Format: None}
}

// Warning returns event with WARN level and default type, source and format
func Warning(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: WARN, Type: Any, Source: EvsEmpty, Format: None}
}

// Error returns event with ERR level and default type, source and format
func Error(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: ERR, Type: Any, Source: EvsEmpty, Format: None}
}

// Critical returns event with CRIT level and default type, source and format
func Critical(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: CRIT, Type: Any, Source: EvsEmpty, Format: None}
}

// Panic returns event with PANIC level and default type, source and format
func Panic(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: PANIC, Type: Any, Source: EvsEmpty, Format: None}
}

// Fatal returns event with FATAL level and default type, source and format
func Fatal(t string) Event {
	return Event{Text: t, Time: time.Now(), Level: FATAL, Type: Any, Source: EvsEmpty, Format: None}
}
