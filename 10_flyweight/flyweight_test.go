package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChessBoard(t *testing.T) {
	board1 := NewChessBoard()
	board2 := NewChessBoard()
	assert.Equal(t, board1.chessPieces[1].pieceUnit, board2.chessPieces[1].pieceUnit)
}
