package events

type EventLogType int

const (
	NoType EventLogType = iota
	CLI
	Stored
	InMemory
	Any
)
