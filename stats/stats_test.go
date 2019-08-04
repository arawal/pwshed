package stats

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUpdateStats(t *testing.T) {
	Init()
	currentStats := stats

	UpdateStats(5003.0)
	newStats := stats

	assert.Equal(t, currentStats.Count+1, newStats.Count)
}

func TestGetCurrentStats(t *testing.T) {
	Init()
	_, err := GetCurrentStats()

	assert.Equal(t, nil, err)
}
