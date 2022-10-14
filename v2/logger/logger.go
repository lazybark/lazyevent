package logger

import (
	"encoding/json"
	"fmt"

	"github.com/lazybark/go-helpers/cli/clf"
	"github.com/lazybark/lazyevent/v2/events"
)

// ILogger represents general logger interface that can process
// specific event types and log events into abstract log.
// It can be CLI, text, SQL, Redis, Sentry, etc.
type ILogger interface {
	Log(events.Event, string) error
	Type() []events.LogType
}

// LogPattern is the default log pattern to transform events into text messages.
// Structure is: eventID->time->level->source->text
var LogPattern = "%s	%s	%s	%s	%s\n"

// LogPatternPureText is used to create pure text messages without time, id, etc.
// Only line break is added.
var LogPatternPureText = "%s\n"

// Format returns event data in string formatted accordingly to LogPattern
func Format(e events.Event, timeFormat string) string {
	return fmt.Sprintf(LogPattern, e.ID, e.Time.Format(timeFormat), e.Level, e.Source, e.Text)
}

func FormatPureText(e events.Event) string {
	return fmt.Sprintf(LogPatternPureText, e.Text)
}

type LogPatternJSON struct {
	ID     string `json:"id"`
	Time   string `json:"time"`
	Level  string `json:"level"`
	Source string `json:"source"`
	Text   string `json:"text"`
}

// Format returns event data in string formatted accordingly to LogPatternJSON
func FormatJSON(e events.Event, timeFormat string) ([]byte, error) {
	return json.Marshal(&LogPatternJSON{
		ID:     e.ID,
		Time:   e.Time.Format(timeFormat),
		Level:  fmt.Sprint(e.Level),
		Source: e.Source.String(),
		Text:   e.Text,
	},
	)
}

// LogPatternCSV is the default log pattern to transform events into csv records.
// Structure is: eventID;time;level;source;text
var LogPatternCSV = "%s;%s;%s;%s;%s;\n"

var CSVHead = "Event ID;Time;Level;Source;Text;\n"

// FormatCSV returns event data in string formatted accordingly to LogPatternCSV
func FormatCSV(e events.Event, timeFormat string) string {
	return fmt.Sprintf(LogPatternCSV, e.ID, e.Time.Format(timeFormat), e.Level, e.Source, e.Text)
}

// FormatColors is a specific method to add ANSI escape sequences to log entries in CLI
func FormatColors(f events.Format, s string) string {
	if f == events.Red {
		return clf.Red(s)
	}
	if f == events.Green {
		return clf.Green(s)
	}
	if f == events.Yellow {
		return clf.Yellow(s)
	}
	if f == events.Blue {
		return clf.Blue(s)
	}
	if f == events.Magenta {
		return clf.Magenta(s)
	}
	if f == events.Cyan {
		return clf.Cyan(s)
	}
	if f == events.Gray {
		return clf.Gray(s)
	}
	if f == events.White {
		return clf.White(s)
	}
	if f == events.BgBlack {
		return clf.BBlack(s)
	}
	if f == events.BgRed {
		return clf.BRed(s)
	}
	if f == events.BgGreen {
		return clf.BGreen(s)
	}
	if f == events.BgYellow {
		return clf.BYellow(s)
	}
	if f == events.BgBlue {
		return clf.BBlue(s)
	}
	if f == events.BgMagenta {
		return clf.BMagenta(s)
	}
	if f == events.BgCyan {
		return clf.BCyan(s)
	}
	if f == events.BgWhite {
		return clf.BWhite(s)
	}

	return s
}
