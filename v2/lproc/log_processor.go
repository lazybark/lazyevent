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
func (lp *LogProcessor) AddLoggers(la ...logger.ILogger) {
	lp.loggers = append(lp.loggers, la...)
}

// Log logs event according to it's type and level.
//
// It panics after logging PANIC-level and calls to
// exit(2) after logging FATAL event.
func (lp *LogProcessor) Log(e events.Event) {
	//Set ID for event to avoid ambiguity in logs.
	//Skip in case it has custom ID
	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	if !e.TimeFixed {
		e.Time = time.Now()
	}

	if lp.force.forceSource {
		e.Source = lp.force.source
	}
	if lp.force.forceLevel {
		e.Level = lp.force.level
	}

	var logTypes []events.LogType
	for _, lg := range lp.loggers {
		logTypes = lg.Type()
		for _, lt := range logTypes {
			if e.Type == lt || lt == events.Any || e.Type == events.Any {
				if err := lg.Log(e, lp.timeFormat); err != nil && lp.reportErrors {
					go func(err error) { lp.errChan <- fmt.Errorf("error making log record: %w", err) }(err)
				}
				break
			}
		}

	}

	//In case we use chan, we do not panic or exit - it will be job of
	//external routine
	if lp.useChan {
		go lp.SendEventToChan(e)
	} else {
		if e.Level == events.PANIC {
			panic(e.Text)
		}
		if e.Level == events.FATAL {
			os.Exit(2)
		}
	}
}

func (lp *LogProcessor) SendEventToChan(e events.Event) {
	lp.evChan <- e
}

// PanicInCaseErr does nothing if nil or event with level<ERR is provided,
// but panics in case error is not nil.
//
// May generate doubles in case same logger is used for several types
// presenting in lt. Doubles will have same ID
func (lp *LogProcessor) PanicInCaseErr(err interface{}, lt ...events.LogType) {
	if err == nil {
		return
	}

	doLog := false

	e := events.Empty()
	e.Level = events.PANIC

	if er, ok := err.(error); ok {
		e.Text = er.Error()
		doLog = true
	}

	if ev, ok := err.(events.Event); ok {
		//We can not simply make e = ev, need to investigate
		if ev.Level > events.WARN {
			doLog = true
			if ev.ID != "" {
				e.ID = ev.ID
			}
			e.Level = events.PANIC
			e.Source = ev.Source
			e.Time = ev.Time
			e.Text = ev.Text
			e.TimeFixed = ev.TimeFixed
			e.Format = ev.Format
		}
	}

	if !doLog {
		return
	}

	if len(lt) > 0 {
		e.Type = lt[0]
		lp.Log(e)
	} else {
		lp.Log(e)
	}
}

// FatalInCaseErr does nothing if nil or event with level<ERR is provided,
// but calls os.Exit() in case error is not nil.
//
// May generate doubles in case same logger is used for several types
// presenting in lt. Doubles will have same ID
func (lp *LogProcessor) FatalInCaseErr(err interface{}, lt ...events.LogType) {
	if err == nil {
		return
	}

	doLog := false

	e := events.Empty()
	e.Level = events.FATAL

	if er, ok := err.(error); ok {
		e.Text = er.Error()
		doLog = true
	}

	if ev, ok := err.(events.Event); ok {
		//We can not simply make e = ev, need to investigate
		if ev.Level > events.WARN {
			doLog = true
			if ev.ID != "" {
				e.ID = ev.ID
			}
			e.Level = events.FATAL
			e.Source = ev.Source
			e.Time = ev.Time
			e.Text = ev.Text
			e.TimeFixed = ev.TimeFixed
			e.Format = ev.Format
		}
	}

	if !doLog {
		return
	}

	if len(lt) > 0 {
		e.Type = lt[0]
		lp.Log(e)
	} else {
		lp.Log(e)
	}
}

// LogErrOnly simply logs any error or does nothind in case nil.
//
// May generate doubles in case same logger is used for several types
// presenting in lt. Doubles will have same ID
func (lp *LogProcessor) LogErrOnly(err interface{}, lt ...events.LogType) {
	if err == nil {
		return
	}

	doLog := false

	e := events.Empty()
	//Setting ID in case there are few logtypes we should use
	e.ID = uuid.New().String()
	e.Level = events.ERR

	if er, ok := err.(error); ok {
		e.Text = er.Error()
		doLog = true
	}

	if ev, ok := err.(events.Event); ok {
		//We can not simply make e = ev, need to investigate
		if ev.Level > events.WARN {
			doLog = true
			if ev.ID != "" {
				e.ID = ev.ID
			}
			e.Level = ev.Level
			e.Source = ev.Source
			e.Time = ev.Time
			e.Text = ev.Text
			e.TimeFixed = ev.TimeFixed
			e.Format = ev.Format
		}

	}

	if !doLog {
		return
	}

	if len(lt) > 0 {
		for _, l := range lt {
			e.Type = l
			lp.Log(e)
		}
	} else {
		lp.Log(e)
	}
}

func (lp LogProcessor) LogRed(e events.Event) {
	lp.Log(e.Red())
}
