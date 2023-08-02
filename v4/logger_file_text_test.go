package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoggerFileTextType(t *testing.T) {
	m := &MockFile{}
	lg1, err := NewPlaintext("some", false, false, 1, m, Any, ErrorFlow)
	require.NoError(t, err)

	types := lg1.Type()

	assert.Equal(t, Any, types[0])
	assert.Equal(t, ErrorFlow, types[1])
}

func TestLoggerFileTextLog(t *testing.T) {
	m := &MockFile{}
	lg1, err := NewPlaintext("some", false, false, 1, m, Any)
	require.NoError(t, err)

	m2 := &MockFile{}
	lg2, err := NewPlaintext("some", true, false, 1, m2, Any)
	require.NoError(t, err)

	now := time.Now()
	lg1.lastLog = now
	lg2.lastLog = now
	time.Sleep(time.Second * 1)

	e := Event{
		Text:      "event1",
		Type:      Any,
		Time:      time.Now(), //Time should be set manually, so we know it
		TimeFixed: true,       //And fixed to not be updated by processor
	}

	p := New(false, "", make(chan error), false, lg1, lg2)

	p.Log(e)
	assert.Equal(t, FormatOutput(e, time.UnixDate), m.Text[0]) //time.UnixDate is default here
	assert.Equal(t, FormatOutputPureText(e), m2.Text[0])
	assert.Equal(t, true, lg1.lastLog.After(now))
	assert.Equal(t, true, lg2.lastLog.After(now))
}
