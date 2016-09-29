package cc

type ChessEnum byte
type ChessColor byte
type OpEnum string

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

const (
	OP_NULL OpEnum = ""
	OP_FORWARD OpEnum = "进"
	OP_BACKWARD OpEnum = "退"
	OP_HORIZONTAL OpEnum = "平"
)

const (
	CN_CAR = "车"
	CN_HORSE = "马"
	CN_CANNON = "炮"
	CN_ELEPHANT_1 = "相"
	CN_ELEPHANT_2 = "象"
	CN_GUARD_1 = "仕"
	CN_GUARD_2 = "士"
	CN_KING_1 = "帅"
	CN_KING_2 = "将"
	CN_PAWN_1 = "兵"
	CN_PAWN_2 = "卒"
)
