package events

import "fmt"

// Source is just a text representation of any possible event source
// in the app. Its desirable to create your own sources to determine
// runtime problems specific to your app.
type Source struct {
	Text  string
	Open  string
	Close string
}

func (s Source) String() string {
	return fmt.Sprintf("%s%s%s", s.Open, s.Text, s.Close)
}

var (
	//EvsEmpty is an empty source to create log records with no source record
	EvsEmpty = Source{}

	//EvsDebug is a default event source to mark debug messages
	EvsDebug = Source{
		Text:  "DEBUG",
		Open:  "[",
		Close: "]",
	}

	//EvsMain represents main() function as an event source
	EvsMain = Source{
		Text:  "MAIN",
		Open:  "[",
		Close: "]",
	}
)
