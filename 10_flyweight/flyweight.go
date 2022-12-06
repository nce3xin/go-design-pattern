package flyweight

import (
	"fmt"
	"k8s.io/klog/v2"
)

var pieceUnits = map[int]*ChessPieceUnit{
	1: NewChessPieceUnit(1, "车", "red"),
	2: NewChessPieceUnit(2, "马", "red"),
	3: NewChessPieceUnit(3, "炮", "red"),
	4: NewChessPieceUnit(4, "车", "black"),
	5: NewChessPieceUnit(5, "马", "black"),
	6: NewChessPieceUnit(6, "炮", "black"),
	//...省略摆放其他棋子的代码
}

// ChessPieceUnit 棋子享元类，包含棋子类中不变的部分
type ChessPieceUnit struct {
	id    int
	text  string
	color string
}

func NewChessPieceUnit(id int, text string, color string) *ChessPieceUnit {
	return &ChessPieceUnit{
		id:    id,
		text:  text,
		color: color,
	}
}

func GetPieceUnit(id int) *ChessPieceUnit {
	pu, ok := pieceUnits[id]
	if !ok {
		klog.Errorf("Id = %v is not in cache!", id)
		panic(fmt.Sprintf("Id = %v is not in cache!", id))
	}
	return pu
}

// ChessPiece 棋子类
type ChessPiece struct {
	pieceUnit *ChessPieceUnit
	positionX int
	positionY int
}

func NewChessPiece(pu *ChessPieceUnit, posX, posY int) *ChessPiece {
	return &ChessPiece{
		pieceUnit: pu,
		positionX: posX,
		positionY: posY,
	}
}

// ChessBoard 棋盘类
type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func (c *ChessBoard) init() {
	c.chessPieces = make(map[int]*ChessPiece)
	c.chessPieces[1] = NewChessPiece(GetPieceUnit(1), 0, 0)
	c.chessPieces[2] = NewChessPiece(GetPieceUnit(2), 1, 0)
	// ...省略摆放其他棋子的代码
}

func (c *ChessBoard) move(chessPieceId int, posX, posY int) {
	// TODO
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{}
	board.init()
	return board
}
