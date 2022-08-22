package events

import "time"

//Event represents any event that can occur during runtime.
//All Event fields are public and it's perfectly safe to modify
//Event without using API
type Event struct {
	ID     string
	Level  EventLevel
	Type   EventLogType
	Source Source
	Time   time.Time
	Text   string
	Format SourceFormat
}
