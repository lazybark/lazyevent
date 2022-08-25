package events

//Level determines how critical the event is
type Level int

const (
	el_start Level = iota

	//INFO informs reader about any event in app
	INFO

	//NOTE is like info, but with higher priority (e.g. for reader to make notes)
	NOTE

	//WARN warns about potentially dangreous situation
	WARN

	//ERR reports that something bad has happened
	ERR

	//CRIT reports that something REALLY bad has happened
	CRIT

	//Panic makes app panic after event is logged
	PANIC

	//FATAL makes app exit after event is logged
	FATAL

	el_end
)

var eLevelNames = [...]string{
	"",
	"INFO",
	"NOTE",
	"WARNING",
	"ERROR",
	"CRITICAL",
	"PANIC",
	"FATAL",
}

func (l Level) String() string {
	if l <= el_start || l >= el_end {
		return ""
	}
	return eLevelNames[l]
}

//IsError returns true in case level of event is high enough
//to be a problem right now.
func (l Level) IsError() bool {
	return l > WARN
}
