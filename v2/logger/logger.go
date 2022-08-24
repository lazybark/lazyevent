package logger

import (
	"encoding/json"
	"fmt"

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

// Format returns event data in string formatted accordingly to LogPattern
func Format(e events.Event, timeFormat string) string {
	return fmt.Sprintf(LogPattern, e.ID, e.Time.Format(timeFormat), e.Level, e.Source, e.Text)
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
