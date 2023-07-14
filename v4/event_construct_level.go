package logger

//Info sets event Level to INFO
func (e Event) Info() Event {
	e.Level = INFO
	return e
}

//Note sets event Level to NOTE
func (e Event) Note() Event {
	e.Level = NOTE
	return e
}

//Warning sets event Level to WARN
func (e Event) Warning() Event {
	e.Level = WARN
	return e
}

//Error sets event Level to ERR
func (e Event) Error() Event {
	e.Level = ERR
	return e
}

//Critical sets event Level to CRIT
func (e Event) Critical() Event {
	e.Level = CRIT
	return e
}

//Panic sets event Level to PANIC
func (e Event) Panic() Event {
	e.Level = PANIC
	return e
}

//Fatal sets event Level to FATAL
func (e Event) Fatal() Event {
	e.Level = FATAL
	return e
}
