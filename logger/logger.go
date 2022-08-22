package logger

import (
	"github.com/lazybark/lazyevent/events"
)

type Logger interface {
	Log(events.Event) error
	Type() events.EventLogType
}
