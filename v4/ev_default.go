package logger

//EvDefault is a helper that can be used to create events with pre-defined
//properties. It will create events with defined Source, Type & Format
//in case these props are set to non-default.
//
//Several EvDefaults can be created in main() to cover whole app needs
//or one by one in methods where similar events are used.
type EvDefault struct {
	Source    Source
	Type      LogType
	Format    Format
	TimeFixed bool
}

func (ed EvDefault) setProps(e Event) Event {
	if ed.Source != EvsEmpty {
		e.Source = ed.Source
	}
	if ed.Type != Any {
		e.Type = ed.Type
	}
	if ed.Format != None {
		e.Format = ed.Format
	}
	if ed.TimeFixed {
		e.TimeFixed = true
	}
	return e
}

// Empty returns event with INFO level, no text and default type, source and format
func (ed EvDefault) Empty() Event {
	return ed.setProps(Empty())
}

// Info returns event with INFO level and type, source, format set to ed parameters
func (ed EvDefault) Info(t string) Event {
	return ed.setProps(Info(t))
}

// Note returns event with NOTE level and type, source, format set to ed parameters
func (ed EvDefault) Note(t string) Event {
	return ed.setProps(Note(t))
}

// Warning returns event with WARN level and type, source, format set to ed parameters
func (ed EvDefault) Warning(t string) Event {
	return ed.setProps(Warning(t))
}

// Error returns event with ERR level and type, source, format set to ed parameters
func (ed EvDefault) Error(t string) Event {
	return ed.setProps(Error(t))
}

// Critical returns event with CRIT level and type, source, format set to ed parameters
func (ed EvDefault) Critical(t string) Event {
	return ed.setProps(Critical(t))
}

// Panic returns event with PANIC level and type, source, format set to ed parameters
func (ed EvDefault) Panic(t string) Event {
	return ed.setProps(Panic(t))
}

// Fatal returns event with FATAL level and type, source, format set to ed parameters
func (ed EvDefault) Fatal(t string) Event {
	return ed.setProps(Fatal(t))
}
