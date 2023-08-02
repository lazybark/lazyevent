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
* out of the box support of Sentryб CLI, text-, JSON- & CSV-file logging (Redis & SQLite will be added in future)
* events are objects that can be stored, passed, modified and logged several times without creating new instance
* auto-rotating logfiles for plaintext, CSV & JSON loggers after desired period of time

## Tips
You can avoid creating event ID if you set useID parameter for logger.New() function to false. All events will not have IDs.
