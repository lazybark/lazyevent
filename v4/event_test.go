package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventFormats(t *testing.T) {
	e := Event{}

	e = e.Red()
	assert.Equal(t, e.Format, Red)
	e = e.Green()
	assert.Equal(t, e.Format, Green)
	e = e.Yellow()
	assert.Equal(t, e.Format, Yellow)
	e = e.Blue()
	assert.Equal(t, e.Format, Blue)
	e = e.Magenta()
	assert.Equal(t, e.Format, Magenta)
	e = e.Cyan()
	assert.Equal(t, e.Format, Cyan)
	e = e.Gray()
	assert.Equal(t, e.Format, Gray)
	e = e.White()
	assert.Equal(t, e.Format, White)
	e = e.BgBlack()
	assert.Equal(t, e.Format, BgBlack)
	e = e.BgRed()
	assert.Equal(t, e.Format, BgRed)
	e = e.BgGreen()
	assert.Equal(t, e.Format, BgGreen)
	e = e.BgGreen()
	assert.Equal(t, e.Format, BgGreen)
	e = e.BgYellow()
	assert.Equal(t, e.Format, BgYellow)
	e = e.BgBlue()
	assert.Equal(t, e.Format, BgBlue)
	e = e.BgMagenta()
	assert.Equal(t, e.Format, BgMagenta)
	e = e.BgCyan()
	assert.Equal(t, e.Format, BgCyan)
	e = e.BgWhite()
	assert.Equal(t, e.Format, BgWhite)
	e = e.None()
	assert.Equal(t, e.Format, None)

}

func TestEventLevel(t *testing.T) {
	e := Event{}

	e = e.Info()
	assert.Equal(t, e.Level, INFO)
	e = e.Note()
	assert.Equal(t, e.Level, NOTE)
	e = e.Warning()
	assert.Equal(t, e.Level, WARN)
	e = e.Error()
	assert.Equal(t, e.Level, ERR)
	e = e.Critical()
	assert.Equal(t, e.Level, CRIT)
	e = e.Panic()
	assert.Equal(t, e.Level, PANIC)
	e = e.Fatal()
	assert.Equal(t, e.Level, FATAL)
}

func TestEventType(t *testing.T) {
	e := Event{}

	e = e.Any()
	assert.Equal(t, e.Type, Any)
	e = e.Main()
	assert.Equal(t, e.Type, Main)
	e = e.Verbose()
	assert.Equal(t, e.Type, Verbose)
	e = e.Debug()
	assert.Equal(t, e.Type, Debug)
}
