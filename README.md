# lazyevent
[![Test](https://github.com/lazybark/lazyevent/actions/workflows/test.yml/badge.svg)](https://github.com/lazybark/lazyevent/actions/workflows/test.yml)

LazyEvent - easy to use logger package that can be customized for almost any app. It treats events as objects with properties that can be created, modified, and logged (or not). An event is any event during runtime, and its possible to prepare record in advance and then use as template. Multiple logging of the same event will generate unique ID for every log record to avoid ambiguity. Time of event can be fixed or updated each time the template is logged.

Log records can be separated by LogType property in the event: only events with suitable type will be passed on to logger. So it's possible to pass main app flow into one log and error/debug/any other flow into separate log container.

Features:

* event-based logging to several logs via one Log processor
* support of any possible custom loggers that have Log(events.Event, string) error and Type() []events.LogType methods
* support of custom log types that can separately log events using any specific log schema
* Log processor can pass logging errors to external routine via channel
* any event can also be directed to external routine via channel (after being logged)
* support of panic & os.Exit() right after logging specific event levels (PANIC & FATAL)
* methods to await output and log error from external functions
* custom styling for records with Event.Format property
* out of the box support of Sentry, CLI, text-, JSON- & CSV-file logging (Redis & SQLite will be added in future)
* events are objects that can be stored, passed, modified and logged several times without creating new instance
* auto-rotating logfiles for plaintext, CSV & JSON loggers after desired period of time

### Event
Event is an object that can be returned by a function, created in advance and filled with function output or simply created + logged in one moment. It has only public parameters:

```
type Event struct {
	ID     string
	Level  Level
	Type   LogType
	Source Source
	Time   time.Time
	Text   string
	Format Format
	TimeFixed bool
}
```

**ID:** unique identifier of event. Log processor will generate ID (using Google's UUID) only if ID was empty at the moment Log() was called and if LP had option useID = true. If you wish to use an event as a template, then ID should stay empty. But if you need to use own IDs, then FlushID() can be called to clean it, or SetID() to set new one. ID = "..." is a good, but not very readable option (i think).

**Level:** one of predefined levels that determines how critical the event is. Higher is worse. Levels are: INFO, NOTE, WARN, ERR, CRIT, PANIC, FATAL. Events with PANIC and FATAL will cause Log processor to call panic() or exit() after logging.

**Type:** type of logger that should be used to log event. Type helps split logs by meaning, so main log will not be populated with debug or verbose info. Predefined types are Any, Main, ErrorFlow, Verbose and Debug. But you can create own types and pass to default or custom loggers.

**Source:** text representation of file/function/package or anything else we need to determine the place where event has occurred. It makes error tracing easier and logs look more user-friendly. It does not affect the event itself or logging process at all (for now). Also, Source can be used to prettify usual runtime logs (not just errors) for CLI app.

**Time:** time of event creation, which will be replaced with logging time in case TimeFixed in NOT true. Loggers use a copy of event, so Time of original event always stays the same. It's useful to make records about several events that occurred at the same second, but have different text/level/type.

**Text:** string that will tell human what happened.

**Format:** is a events.Format value that tells logger how to treat specific log entries. It could be a color or some text transformations or any other way of making event stand out. Standard formats only indicate colors for CLI, but you can create custom formats for custom loggers.

**TimeFixed:** bool value that tells Log processor to NOT update event time before logging. By default this value is updated to avoid time shift in cases when there is a latency between event creation and logging.

### Create & update event
There are two convenient ways to create an event: directly from events package and from an object called `EvDefault`.

The first case is suitable when we have only few events in a function or when all events are more or less standard. Because events created this way would have only default field values, except `Level` and `Text`. It can be done by calling `LEVEL_NAME("event text")` (possible levels: `Info`, `Note`, `Warning`, `Error`, `Critical`, `Panic`, `Fatal`). `Empty()` will create an event with INFO level, but empty text.

Second case is more convenient when we need to deploy many similar events across the app or function. We can create object of events.EvDefault as a template and then generate events with methods described above. All events will have Source, Type, Format and TimeFixed parameters similar to instance of EvDefault.

```
ed := events.EvDefault{
		Source: events.EvsMain,
		Type:   events.Debug,
		Format: events.None,
	}
	ed.Note("Event from default template")
```
We also can put an event together manually by specifying all fields values if there is any need of some custom options.

Event updating is a simple process that can happen via manual value setting or default methods: default types and formats can be set by calling methods like `e.Verbose()` or `e.Red()`, same goes for levels described above. To set source, you would need to call `e.Src()`. Text can be changed via `e.SetText()` and fixing/unfixing time is possible via `e.FixTime()/UnFixTime()`

### Log event
Event can be logged by calling `Log()` on `LogProcessor`.

Processor is created by calling `New(timeFormat string, errChan chan (error), reportErrors bool, la ...ILogger)`. timeFormat here represents template to format event time in log records. For default loggers it should be one of formats we use with time.Time.Format(), but custom loggers may use any other format. If your logger uses something special in that case, you may want to have only you custom loggers in LogProcessor to avoid problems with default ones. errChan will be used to send errors from loggers in case reportErrors = true.

Also LogProcessor has several methods to log errors only and ignore other event levels. LogErrOnly() will drop a log only in case it will receive error or event with ERR level. FatalInCaseErr() will do the same, but only with error or event of FATAL level. PanicInCaseErr() will act accordingly.

> NOTE
> 
> LogErrOnly simply sets types to events. It does not change standard log sequence. So, if you have a logger that uses events.Any logtype, it will get and log the event even if you sent specific logtype to LogErrOnly. That's why doubles may appear in case there are several loggers with crossing types.
>
> Another tricky moment: if an event has events.Any logtype it will be logged by all possible loggers. So it's ok to create strictly specific loggers, but sometimes use events that are meant for Any - those will be logged to every channel.

> NOTE 2
> 
> When creating your own logger, keep in mind that logger may not check event type. LogProcessor does that, so double-checking will just take some extra resources.

### Creating custom logger
All you need is to make it implement the ILogger interface:
```
type ILogger interface {
	//Log processes event and logs according to logger's internal rules
	Log(e Event, timeFormat string) error

	//Type returns whole list of types a logger should receive.
	//Normally you don't need to check event type in the logger as the log processor will not
	//send events that have wrong LogType
	Type() []LogType
}
```

## Tips
You can avoid creating event ID if you set `useID` parameter for `logger.New()` function to false. All events will not have IDs.

Default CLI & plaintext file loggers have pureText directive to print out pure log event text (without time, id, level, etc. - but with color formatting in case CLI).

If you leave fields ID, Source or Level empty, they will still be present on .csv and .json default loggers so log files will be available for correct parsing.

You can use custom file for default text loggers (.log, .csv, .json): just pass IFile interface wich is, for example, is `*os.File`. But if you set `rotateFiles > 0`, the file will be changed after `rotateFiles * time.Minute`. So if you want to use custom file, set rotateFiles to 0.

## Making a part of different project logic
It's actually a good idea to create small interface in your app that suits your needs. Then make a struct that holds a Lazyevent log processor and has methods to define specific events (internally using Lazyevent methods). This way you can configure logger for specific logic of the app and use with ease.

A real-world example of Lazyevent in [go-cloud-sync](https://github.com/lazybark/go-cloud-sync):

```
package lgr

import (
	"fmt"

	logger "github.com/lazybark/lazyevent/v4"
)

// LGR is a simple struct with log processor in it
type LGR struct {
	p *logger.LogProcessor
}

const ltypeEvent = logger.LogType(15) //ltypeEvent is the log type used to mark filesystem events

// Init creates new logger with hardcoded loggers
func Init() (*LGR, error) {
	var loggers []logger.ILogger

	evs, err := logger.NewPlaintext("logs/events/events", false, false, 1, ltypeEvent)
	if err != nil {
		return nil, fmt.Errorf("[LOGGER][Init] %w", err)
	}

	errs, err := logger.NewPlaintext("logs/errors/errs", false, false, 1, logger.ErrorFlow)
	if err != nil {
		return nil, fmt.Errorf("[LOGGER][Init] %w", err)
	}

	cliVerbose := logger.NewCLI(false, logger.Verbose, logger.ErrorFlow)

	loggers = append(loggers, evs, errs, cliVerbose)

	//Uncomment this to log events in CLI
	//evsCLI := logger.NewCLI(false, ltypeEvent)
	//loggers = append(loggers,  evsCLI)

	p := logger.New(false, "", make(chan error), false, loggers...)

	l := LGR{p: p}

	return &l, nil
}

// Info creates new event of INFO level
func (l *LGR) Info(t string, verbose bool) {
	e := logger.Info(t)
	if verbose {
		l.p.Log(e.Verbose())
	} else {
		l.p.Log(e.Main())
	}
}

// Error creates new event of ERROR level
func (l *LGR) Error(err error) {
	if err == nil {
		err = fmt.Errorf("[empty error log attempt]") //We want to know if our logic is bad somewhere
	}
	e := logger.Error(err.Error()).Red()
	e.Type = logger.ErrorFlow
	l.p.Log(e)
}

// Event logs filesystem event
func (l *LGR) Event(t string) {
	e := logger.Info(t)
	e.Type = ltypeEvent
	e.Level = logger.Level(0)
	l.p.Log(e)
}

// LogErrOnly simply logs any error or does nothing in case nil
func (l *LGR) LogErrOnly(e interface{}) {
	l.p.LogErrOnly(e)
}

// PanicInCaseErr logs error and panics. Nothig is done in case nil
func (l *LGR) PanicInCaseErr(e interface{}) {
	l.p.PanicInCaseErr(e)
}
```
