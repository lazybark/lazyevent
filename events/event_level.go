package events

type EventLevel int

const (
	el_start EventLevel = iota

	//All represents all levels. It will be empty in logs and passed to every logger
	//according to loggertype only
	All
	Info
	Note
	Warning
	Error
	Critical
	Panic
	Fatal

	el_end
)

var eLevelNames = [...]string{
	"",
	"",
	"CLI",
	"Stored",
	"InMemory",
	"Any",
	"",
}

func (l EventLevel) String() string {
	if l <= el_start || l >= el_end {
		return ""
	}
	return eLevelNames[l]
}
