package logger

// Format represents any possible format to output log strings.
// Default formats are for colored CLI ouput, but logger developers
// can create new formats to suit specific logger options.
type Format int

const (
	// None should represent default log format for the logger
	None Format = iota

	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	Gray
	White
	BgBlack
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)
