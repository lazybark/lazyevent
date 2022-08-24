package lproc

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lazybark/lazyevent/v2/events"
	"github.com/lazybark/lazyevent/v2/logger"
)

// LogProcessor manages all available loggers, processes log errors
// and sends events to external routine if needed
type LogProcessor struct {
	timeFormat   string
	loggers      []logger.ILogger
	useChan      bool
	evChan       chan (events.Event)
	reportErrors bool
	errChan      chan (error)
	force        Force
}

type Force struct {
	forceSource bool
	source      events.Source
	forceLevel  bool
	level       events.Level
}

// New creates new LogProcessor with selected parameters.
//
// If timeFormat is an empty string, time.UnixDate will be used
func New(timeFormat string, errChan chan (error), reportErrors bool, la ...logger.ILogger) *LogProcessor {
	evChan := make(chan (events.Event))
	if timeFormat == "" {
		timeFormat = time.UnixDate
	}
	p := &LogProcessor{timeFormat: timeFormat, evChan: evChan, errChan: errChan, reportErrors: reportErrors, loggers: la}

	return p
}

// SetErrChan sets new error channel to send log process errors.
func (ep *LogProcessor) SetErrChan(evChan chan (events.Event)) {
	ep.useChan = true
	ep.evChan = evChan
}

// UnSetErrChan prevents EP from sending log errors to outer routine and
// closes the error channel
func (ep *LogProcessor) UnSetErrChan() {
	ep.useChan = false
	close(ep.errChan)
}

// ForceSource makes EP forsibly change source of every event to s
func (ep *LogProcessor) ForceSource(s events.Source) {
	ep.force.forceSource = true
	ep.force.source = s
}

// ForceLevel makes EP forsibly change leve of every event to l
func (ep *LogProcessor) ForceLevel(l events.Level) {
	ep.force.forceLevel = true
	ep.force.level = l
}

// AddLoggers adds list of loggers to EP's pool
func (ep *LogProcessor) AddLoggers(la ...logger.ILogger) {
	ep.loggers = append(ep.loggers, la...)
}

// Log logs event according to it's type and level.
//
// It panics after logging PANIC-level and calls to
// exit(2) after logging FATAL event.
func (ep *LogProcessor) Log(e events.Event) {
	//Set ID for event to avoid ambiguity in logs.
	//Skip in case it has custom ID
	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	if !e.TimeFixed {
		e.Time = time.Now()
	}

	if ep.force.forceSource {
		e.Source = ep.force.source
	}
	if ep.force.forceLevel {
		e.Level = ep.force.level
	}

	var logTypes []events.LogType
	for _, lg := range ep.loggers {
		logTypes = lg.Type()
		for _, lt := range logTypes {
			if e.Type == lt || lt == events.Any {
				if err := lg.Log(e, ep.timeFormat); err != nil && ep.reportErrors {
					go func(err error) { ep.errChan <- fmt.Errorf("error making log record: %w", err) }(err)
				}
				break
			}
		}

	}

	//In case we use chan, we do not panic or exit - it will be job of
	//external routine
	if ep.useChan {
		go func(e events.Event) { ep.evChan <- e }(e)
	} else {
		if e.Level == events.PANIC {
			panic(e.Text)
		}
		if e.Level == events.FATAL {
			os.Exit(2)
		}
	}
}

// PanicInCaseErr does nothing if nil is provided, but panics in case error is not nil.
//
// lt specifies logtype to use for this specific log, but only first one will be used
// as panic will cause app to stop after first log entry.
func (ep *LogProcessor) PanicInCaseErr(err error, lt ...events.LogType) {
	if err == nil {
		return
	}
	e := events.Panic(err.Error())

	if len(lt) > 0 {
		e.Type = lt[0]
		ep.Log(e)
	} else {
		ep.Log(e)
	}
}

// FatalInCaseErr does nothing if nil is provided, but panics in case error is not nil.
//
// lt specifies logtype to use for this specific log, but only first one will be used
// as exit() will cause app to stop after first log entry.
func (ep *LogProcessor) FatalInCaseErr(err error, lt ...events.LogType) {
	if err == nil {
		return
	}
	e := events.Fatal(err.Error())

	if len(lt) > 0 {
		e.Type = lt[0]
		ep.Log(e)
	} else {
		ep.Log(e)
	}
}

// LogErrOnly simply logs any error or does nothind in case nil.
//
// May generate doubles in case same logger is used for several types
// from lt. Doubles ID will stay the same.
func (ep *LogProcessor) LogErrOnly(err error, lt ...events.LogType) {
	if err == nil {
		return
	}
	e := events.Error(err.Error())
	//Setting ID in case there are few logtypes we should use
	e.ID = uuid.New().String()

	if len(lt) > 0 {
		for _, l := range lt {
			e.Type = l
			ep.Log(e)
		}
	} else {
		ep.Log(e)
	}
}