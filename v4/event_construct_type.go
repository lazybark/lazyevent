package logger

//Any sets event Type to Any
func (e Event) Any() Event {
	e.Type = Any

	return e
}

//Main sets event Type to Main
func (e Event) Main() Event {
	e.Type = Main

	return e
}

//Verbose sets event Type to Verbose
func (e Event) Verbose() Event {
	e.Type = Verbose

	return e
}

//Debug sets event Type to Debug
func (e Event) Debug() Event {
	e.Type = Debug

	return e
}
