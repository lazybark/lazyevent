package logger

import (
	"encoding/json"
	"fmt"

	"github.com/lazybark/go-helpers/cli/clf"
)

// ILogger represents general logger interface that can process
// specific event types and log events into abstract log.
// It can be CLI, text, SQL, Redis, Sentry, etc.
type ILogger interface {
	Log(Event, string) error
	Type() []LogType
}

// LogPattern is the default log pattern to transform events into text messages.
// Structure is: eventID->time->level->source->text
var (
	logPatterns = []string{
		"",
		"%s\n",
		"%s	%s\n",
		"%s	%s	%s\n",
		"%s	%s	%s	%s\n",
		"%s	%s	%s	%s	%s\n",
	}
)

// LogPatternPureText is used to create pure text messages without time, id, etc.
// Only line break is added.
var LogPatternPureText = "%s\n"

// SentryPattern is used to create records for Sentry
var SentryPattern = "[%s] %s %s %s"

// FormatOutput returns event data in string formatted accordingly to LogPattern
func FormatOutput(e Event, timeFormat string) string {
	var args []any
	if e.ID != "" {
		args = append(args, e.ID)
	}
	args = append(args, e.Time.Format(timeFormat))
	if e.Level.String() != "" {
		args = append(args, e.Level)
	}
	if e.Source.String() != "" {
		args = append(args, e.Source)
	}
	args = append(args, e.Text)

	return fmt.Sprintf(logPatterns[len(args)], args...)
}

func FormatOutputPureText(e Event) string {
	return fmt.Sprintf(LogPatternPureText, e.Text)
}

// FormatOutputSentry returns event data in string formatted accordingly to SentryPattern
func FormatOutputSentry(e Event, appID string) string {
	if e.Level.String() == "" && e.Source.String() == "" {
		return fmt.Sprintf("%s	%s\n", appID, e.Text)
	}
	if e.Level.String() == "" {
		return fmt.Sprintf("%s	%s	%s\n", appID, e.Source, e.Text)
	}
	if e.Source.String() == "" {
		return fmt.Sprintf("%s	%s	%s\n", appID, e.Level, e.Text)
	}

	return fmt.Sprintf(SentryPattern, appID, e.Level, e.Source, e.Text)
}

type LogPatternJSON struct {
	ID     string `json:"id"`
	Time   string `json:"time"`
	Level  string `json:"level"`
	Source string `json:"source"`
	Text   string `json:"text"`
}

// Format returns event data in string formatted accordingly to LogPatternJSON
func FormatJSON(e Event, timeFormat string) ([]byte, error) {
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
func FormatCSV(e Event, timeFormat string) string {
	return fmt.Sprintf(LogPatternCSV, e.ID, e.Time.Format(timeFormat), e.Level, e.Source, e.Text)
}

// FormatColors is a specific method to add ANSI escape sequences to log entries in CLI
func FormatColors(f Format, s string) string {
	if f == Red {
		return clf.Red(s)
	}
	if f == Green {
		return clf.Green(s)
	}
	if f == Yellow {
		return clf.Yellow(s)
	}
	if f == Blue {
		return clf.Blue(s)
	}
	if f == Magenta {
		return clf.Magenta(s)
	}
	if f == Cyan {
		return clf.Cyan(s)
	}
	if f == Gray {
		return clf.Gray(s)
	}
	if f == White {
		return clf.White(s)
	}
	if f == BgBlack {
		return clf.BBlack(s)
	}
	if f == BgRed {
		return clf.BRed(s)
	}
	if f == BgGreen {
		return clf.BGreen(s)
	}
	if f == BgYellow {
		return clf.BYellow(s)
	}
	if f == BgBlue {
		return clf.BBlue(s)
	}
	if f == BgMagenta {
		return clf.BMagenta(s)
	}
	if f == BgCyan {
		return clf.BCyan(s)
	}
	if f == BgWhite {
		return clf.BWhite(s)
	}

	return s
}
