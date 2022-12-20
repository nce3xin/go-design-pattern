package finitestatemachine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallGotMushroom(t *testing.T) {
	s := NewMarioStateMachine(0)
	s.ObtainMushroom()
	assert.Equal(t, "super", s.currentState.GetName())
	assert.Equal(t, 100, s.score)
}

func TestSuperGotFire(t *testing.T) {
	s := NewMarioStateMachine(0)
	s.currentState = &SuperMario{stateMachine: s}
	s.ObtainFireFlower()
	assert.Equal(t, "fire", s.currentState.GetName())
	assert.Equal(t, 300, s.score)
}
