package logger

//None sets event Format to None
func (e Event) None() Event {
	e.Format = None

	return e
}

//Red sets event Format to Red
func (e Event) Red() Event {
	e.Format = Red

	return e
}

//Green sets event Format to Green
func (e Event) Green() Event {
	e.Format = Green

	return e
}

//Yellow sets event Format to Yellow
func (e Event) Yellow() Event {
	e.Format = Yellow

	return e
}

//Blue sets event Format to Blue
func (e Event) Blue() Event {
	e.Format = Blue

	return e
}

//Magenta sets event Format to Magenta
func (e Event) Magenta() Event {
	e.Format = Magenta

	return e
}

//Cyan sets event Format to Cyan
func (e Event) Cyan() Event {
	e.Format = Cyan

	return e
}

//Gray sets event Format to Gray
func (e Event) Gray() Event {
	e.Format = Gray

	return e
}

//White sets event Format to White
func (e Event) White() Event {
	e.Format = White

	return e
}

//BgBlack sets event Format to BgBlack
func (e Event) BgBlack() Event {
	e.Format = BgBlack

	return e
}

//BgRed sets event Format to BgRed
func (e Event) BgRed() Event {
	e.Format = BgRed

	return e
}

//BgGreen sets event Format to BgGreen
func (e Event) BgGreen() Event {
	e.Format = BgGreen

	return e
}

//BgYellow sets event Format to BgYellow
func (e Event) BgYellow() Event {
	e.Format = BgYellow

	return e
}

//BgBlue sets event Format to BgBlue
func (e Event) BgBlue() Event {
	e.Format = BgBlue

	return e
}

//BgMagenta sets event Format to BgMagenta
func (e Event) BgMagenta() Event {
	e.Format = BgMagenta

	return e
}

//BgCyan sets event Format to BgCyan
func (e Event) BgCyan() Event {
	e.Format = BgCyan

	return e
}

//BgWhite sets event Format to BgWhite
func (e Event) BgWhite() Event {
	e.Format = BgWhite

	return e
}
