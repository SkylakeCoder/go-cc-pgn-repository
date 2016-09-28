package cc

type chessBoard struct {
	chessInfo [][]Chess
}

func (*chessBoard) ToBytes() []byte {
	return nil
}
