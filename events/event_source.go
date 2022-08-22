package events

type Source struct {
	Text   string
	Format string
	Open   string
	Close  string
}

type SourceFormat int

const (
	None SourceFormat = iota
	Red
	Green
)
