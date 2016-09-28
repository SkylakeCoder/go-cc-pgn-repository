package cc

type ChessEnum byte
type ChessColor byte

const (
	CHESS_NULL ChessEnum = iota
	CHESS_CAR
	CHESS_HORSE
	CHESS_CANNON
	CHESS_ELEPHANT
	CHESS_GUARD
	CHESS_KING
	CHESS_PAWN
)

const (
	COLOR_NULL ChessColor = iota
	COLOR_RED
	COLOR_BLACK
)
