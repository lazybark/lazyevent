package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventLevelError(t *testing.T) {
	assert.Equal(t, false, INFO.IsError())
	assert.Equal(t, false, NOTE.IsError())
	assert.Equal(t, false, WARN.IsError())
	assert.Equal(t, true, ERR.IsError())
	assert.Equal(t, true, CRIT.IsError())
	assert.Equal(t, true, PANIC.IsError())
	assert.Equal(t, true, FATAL.IsError())
}

func TestEventLevelNames(t *testing.T) {
	assert.Equal(t, "INFO", INFO.String())
	assert.Equal(t, "NOTE", NOTE.String())
	assert.Equal(t, "WARNING", WARN.String())
	assert.Equal(t, "ERROR", ERR.String())
	assert.Equal(t, "CRITICAL", CRIT.String())
	assert.Equal(t, "PANIC", PANIC.String())
	assert.Equal(t, "FATAL", FATAL.String())

	assert.Equal(t, "", Level(0).String())
	assert.Equal(t, "", Level(11).String())
	assert.Equal(t, "", el_start.String())
	assert.Equal(t, "", el_end.String())
}
