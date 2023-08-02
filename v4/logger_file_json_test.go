package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoggerFileJSONType(t *testing.T) {
	m := &MockFile{}
	lg1, err := NewJSONtext("some", false, 1, m, Any, ErrorFlow)
	require.NoError(t, err)

	types := lg1.Type()

	assert.Equal(t, Any, types[0])
	assert.Equal(t, ErrorFlow, types[1])
}

func TestLoggerFileJSONLog(t *testing.T) {
	m := &MockFile{}
	lg1, err := NewJSONtext("some", false, 1, m, Any)
	require.NoError(t, err)

	now := time.Now()
	lg1.lastLog = now
	time.Sleep(time.Second * 1)

	e := Event{
		Text:      "event1",
		Type:      Any,
		Time:      time.Now(), //Time should be set manually, so we know it
		TimeFixed: true,       //And fixed to not be updated by processor
	}

	p := New(false, "", make(chan error), false, lg1)

	p.Log(e)

	result, err := FormatJSON(e, time.UnixDate) //time.UnixDate is default here
	require.NoError(t, err)

	assert.Equal(t, string(result), m.Text[0])
	assert.Equal(t, true, lg1.lastLog.After(now))
}
