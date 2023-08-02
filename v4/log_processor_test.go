package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Any will be sent into every logger
func TestLogProcessorLogAny(t *testing.T) {
	lg1 := &MockLogger{
		LogType: []LogType{Any},
	}
	lg2 := &MockLogger{
		LogType: []LogType{ErrorFlow},
	}

	e := Event{
		Text: "event1",
		Type: Any,
	}

	p := New(true, "", make(chan error), false, lg1, lg2)

	p.Log(e)

	assert.Equal(t, e.Text, lg1.LoggedData.Text)
	assert.Equal(t, true, lg1.wasCalledLog)
	assert.Equal(t, true, lg1.wasCalledType)

	assert.Equal(t, "event1", lg2.LoggedData.Text)
	assert.Equal(t, true, lg2.wasCalledLog)
	assert.Equal(t, true, lg2.wasCalledType)
}

// Main will be sent to Main or Any logger
func TestLogProcessorLogMain(t *testing.T) {
	lg1 := &MockLogger{
		LogType: []LogType{Any},
	}
	lg2 := &MockLogger{
		LogType: []LogType{ErrorFlow},
	}
	lg3 := &MockLogger{
		LogType: []LogType{Main},
	}

	e := Event{
		Text: "event1",
		Type: Main,
	}

	p := New(true, "", make(chan error), false, lg1, lg2, lg3)

	p.Log(e)

	assert.Equal(t, e.Text, lg1.LoggedData.Text)
	assert.Equal(t, true, lg1.wasCalledLog)
	assert.Equal(t, true, lg1.wasCalledType)

	assert.Equal(t, "", lg2.LoggedData.Text)
	assert.Equal(t, false, lg2.wasCalledLog)
	assert.Equal(t, true, lg2.wasCalledType)

	assert.Equal(t, e.Text, lg3.LoggedData.Text)
	assert.Equal(t, true, lg3.wasCalledLog)
	assert.Equal(t, true, lg3.wasCalledType)
}

// Should generate id only if asked to
func TestLogProcessorIDGeneration(t *testing.T) {
	lg1 := &MockLogger{
		LogType: []LogType{Any},
	}

	e := Event{
		Text: "event1",
		Type: Main,
	}

	p := New(false, "", make(chan error), false, lg1)
	p2 := New(true, "", make(chan error), false, lg1)

	p.Log(e)
	assert.Equal(t, lg1.LoggedData.ID, "")

	p2.Log(e)
	assert.Equal(t, false, lg1.LoggedData.ID == "")
}

// Should reset time if TimeFixed = true
func TestLogProcessorTimeUpdating(t *testing.T) {
	lg1 := &MockLogger{
		LogType: []LogType{Any},
	}

	e := Event{
		Text:      "event1",
		Type:      Main,
		TimeFixed: true,
		Time:      time.Now().Add(time.Hour * 15),
	}
	e2 := Event{
		Text:      "event2",
		Type:      Main,
		TimeFixed: false,
		Time:      time.Now().Add(time.Hour * -1),
	}

	p := New(false, "", make(chan error), false, lg1)

	p.Log(e)
	assert.Equal(t, true, lg1.LoggedData.Time.Equal(e.Time))

	p.Log(e2)
	assert.Equal(t, true, lg1.LoggedData.Time.After(e2.Time))
}

// Should force level & source if asked to
func TestLogProcessorForceSourceAndLevel(t *testing.T) {
	lg1 := &MockLogger{
		LogType: []LogType{Any},
	}

	e := Event{
		Text:      "event1",
		Type:      Main,
		Source:    EvsMain,
		Level:     INFO,
		TimeFixed: true,
		Time:      time.Now().Add(time.Hour * 15),
	}

	p := New(false, "", make(chan error), false, lg1)

	p.Log(e)
	assert.Equal(t, EvsMain, lg1.LoggedData.Source)
	assert.Equal(t, INFO, lg1.LoggedData.Level)

	p.ForceSource(EvsDebug)
	p.ForceLevel(WARN)

	p.Log(e)
	assert.Equal(t, EvsDebug, lg1.LoggedData.Source)
	assert.Equal(t, WARN, lg1.LoggedData.Level)

}
