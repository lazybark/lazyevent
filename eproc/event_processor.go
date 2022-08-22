package eproc

import (
	"fmt"
	"os"

	"github.com/lazybark/lazyevent/events"
	"github.com/lazybark/lazyevent/logger"
)

type EventProcessor struct {
	reportErrors bool
	loggers      []logger.Logger
	evChan       chan (events.Event)
	errChan      chan (error)
}

func New(errChan chan (error), reportErrors bool, la ...logger.Logger) *EventProcessor {
	evChan := make(chan (events.Event))
	p := &EventProcessor{evChan: evChan, errChan: errChan, reportErrors: reportErrors, loggers: la}
	go p.start()

	return p
}

func (ep *EventProcessor) AddLoggers(la ...logger.Logger) {
	ep.loggers = append(ep.loggers, la...)
}

func (ep *EventProcessor) Log(e events.Event) {
	logType := events.EventLogType(0)
	for _, lg := range ep.loggers {
		logType = lg.Type()
		if e.Type == logType || logType == events.Any {
			if err := lg.Log(e); err != nil {
				ep.errChan <- fmt.Errorf("error making log record: %w", err)
			}
		}
	}
}

// start launches routine to log events
func (ep *EventProcessor) start() {
	for e := range ep.evChan {
		ep.Log(e)
		if e.Level == events.Panic {
			panic(e.Text)
		}
		if e.Level == events.Fatal {
			os.Exit(2)
		}
	}
}
