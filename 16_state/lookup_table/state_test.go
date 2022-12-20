package statelookuptable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallGotMushroom(t *testing.T) {
	s := NewMarioStateMachine(0, SMALL)
	s.obtainMushroom()
	assert.Equal(t, SUPER, s.currentState)
	assert.Equal(t, 100, s.score)
}

func TestSuperGotFire(t *testing.T) {
	s := NewMarioStateMachine(0, SUPER)
	s.obtainFireFlower()
	assert.Equal(t, FIRE, s.currentState)
	assert.Equal(t, 300, s.score)
}
