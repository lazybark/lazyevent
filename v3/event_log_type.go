package logger

//LogType represents type of logger that should log the event.
//It can be used in cases when some events have specific
//meaning and should not appear in some logs. E.g. debug messages
//in CLI or request logs
type LogType int

const (
	//Any-logtyped event will be logged by all available loggers
	Any LogType = iota

	//Main represents main event sequence that stores all app events
	Main

	//ErrorFlow represents error log in case errors should not be
	//in main event sequence
	ErrorFlow

	//Verbose indicates that this specific event should be considered as verbose
	//and app may want to use specific logger to treat that event.
	Verbose

	//Debug event typically means info that should be read by developer or QA only
	Debug
)
