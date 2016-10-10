package chess

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"github.com/axgle/mahonia"
	"os"
	"repository"
)

type ChessBoard struct {
	chessInfo [][]*Chess
}

type Point struct {
	X int
	Y int
}

var debugFlag = false

func (cb *ChessBoard) Init() {
	cb.chessInfo = [][]*Chess {}
	for i := 0; i < BOARD_ROW; i++ {
		row := []*Chess {}
		for j := 0; j < BOARD_COL; j++ {
			chess := &Chess {
				Type:CHESS_NULL,
				Color:COLOR_NULL,
			}
			row = append(row, chess)
		}
		cb.chessInfo = append(cb.chessInfo, row)
	}

	// Black.
	cb.chessInfo[0][0] = &Chess {Type:CHESS_CAR, Color:COLOR_BLACK}
	cb.chessInfo[0][1] = &Chess {Type:CHESS_HORSE, Color:COLOR_BLACK}
	cb.chessInfo[0][2] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_BLACK}
	cb.chessInfo[0][3] = &Chess {Type:CHESS_GUARD, Color:COLOR_BLACK}
	cb.chessInfo[0][4] = &Chess {Type:CHESS_KING, Color:COLOR_BLACK}
	cb.chessInfo[0][5] = &Chess {Type:CHESS_GUARD, Color:COLOR_BLACK}
	cb.chessInfo[0][6] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_BLACK}
	cb.chessInfo[0][7] = &Chess {Type:CHESS_HORSE, Color:COLOR_BLACK}
	cb.chessInfo[0][8] = &Chess {Type:CHESS_CAR, Color:COLOR_BLACK}

	cb.chessInfo[2][1] = &Chess {Type:CHESS_CANNON, Color:COLOR_BLACK}
	cb.chessInfo[2][7] = &Chess {Type:CHESS_CANNON, Color:COLOR_BLACK}

	cb.chessInfo[3][0] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cb.chessInfo[3][2] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cb.chessInfo[3][4] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cb.chessInfo[3][6] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}
	cb.chessInfo[3][8] = &Chess {Type:CHESS_PAWN, Color:COLOR_BLACK}

	// Red.
	cb.chessInfo[9][0] = &Chess {Type:CHESS_CAR, Color:COLOR_RED}
	cb.chessInfo[9][1] = &Chess {Type:CHESS_HORSE, Color:COLOR_RED}
	cb.chessInfo[9][2] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_RED}
	cb.chessInfo[9][3] = &Chess {Type:CHESS_GUARD, Color:COLOR_RED}
	cb.chessInfo[9][4] = &Chess {Type:CHESS_KING, Color:COLOR_RED}
	cb.chessInfo[9][5] = &Chess {Type:CHESS_GUARD, Color:COLOR_RED}
	cb.chessInfo[9][6] = &Chess {Type:CHESS_ELEPHANT, Color:COLOR_RED}
	cb.chessInfo[9][7] = &Chess {Type:CHESS_HORSE, Color:COLOR_RED}
	cb.chessInfo[9][8] = &Chess {Type:CHESS_CAR, Color:COLOR_RED}

	cb.chessInfo[7][1] = &Chess {Type:CHESS_CANNON, Color:COLOR_RED}
	cb.chessInfo[7][7] = &Chess {Type:CHESS_CANNON, Color:COLOR_RED}

	cb.chessInfo[6][0] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cb.chessInfo[6][2] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cb.chessInfo[6][4] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cb.chessInfo[6][6] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
	cb.chessInfo[6][8] = &Chess {Type:CHESS_PAWN, Color:COLOR_RED}
}

func (cb *ChessBoard) Reset() {
	cb.Init()
}

func (cb *ChessBoard) ToString() string {
	str := ""
	for i := 0; i < BOARD_ROW; i++ {
		for j := 0; j < BOARD_COL; j++ {
			str += cb.chessInfo[i][j].String()
		}
	}
	return str
}

func (cb *ChessBoard) ParseRecord(recordPath string) bool {
	record, err := ioutil.ReadFile(recordPath)
	if err != nil {
		log.Fatalln("load record failed...")
		return false
	}
	utf8Record := mahonia.NewDecoder("gbk").ConvertString(string(record))
	lines := strings.Split(string(utf8Record), "\n")
	lines[0] = ""
	commentOpen := false
	isRed := true
	for _, line := range(lines) {
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") {
			continue
		}
		if strings.HasPrefix(line, "{") {
			commentOpen = true
			continue
		}
		if strings.HasPrefix(line, "}") {
			commentOpen = false
			continue
		}
		if commentOpen {
			continue
		}
		if false {
			lineLen := 0
			for _, _ = range line {
				lineLen++
			}
			if lineLen < 8 {
				continue
			}
		}
		if (!strings.Contains(line, string(OP_HORIZONTAL)) &&
			!strings.Contains(line, string(OP_FORWARD)) &&
			!strings.Contains(line, string(OP_BACKWARD))) {
			continue
		}

		splitList := strings.Split(line, " ")
		for _, v := range splitList {
			if len(v) <= 3 {
				continue
			}
			key := cb.ToString()
			if debugFlag {
				fmt.Println("---------------------------------------")
				fmt.Printf("%s\n", v)
			}
			if strings.Contains(v, CN_CAR) {
				cb.moveCar(v, isRed)
			} else if (strings.Contains(v, CN_HORSE)) {
				cb.moveHorse(v, isRed)
			} else if (strings.Contains(v, CN_CANNON)) {
				cb.moveCannon(v, isRed)
			} else if (strings.Contains(v, CN_ELEPHANT_1) ||
					   strings.Contains(v, CN_ELEPHANT_2)) {
				cb.moveElephant(v, isRed)
			} else if (strings.Contains(v, CN_GUARD_1) ||
					   strings.Contains(v, CN_GUARD_2)) {
				cb.moveGuard(v, isRed)
			} else if (strings.Contains(v, CN_KING_1) ||
					   strings.Contains(v, CN_KING_2)) {
				cb.moveKing(v, isRed)
			} else if (strings.Contains(v, CN_PAWN_1) ||
					   strings.Contains(v, CN_PAWN_2)) {
				cb.movePawn(v, isRed)
			} else {
				log.Fatalln("unknown chess type..." + v)
				os.Exit(-1)
			}
			if debugFlag {
				cb.DumpForDebug()
				fmt.Println("---------------------------------------")
				fmt.Println("")
			}
			isRed = !isRed
			value := cb.ToString()
			repository.Record(key, value)
		}
	}
	// log.Printf("[path done] %s\n", recordPath)
 	return true
}

func (cb *ChessBoard) DumpForDebug() {
	for row := 0; row < BOARD_ROW; row++ {
		for col := 0; col < BOARD_COL; col++ {
			chess := cb.chessInfo[row][col]
			if chess.Type == CHESS_NULL {
				fmt.Print("　　 ")
			} else {
				strColor := ""
				if chess.Color == COLOR_RED {
					strColor = "红"
				} else {
					strColor = "黑"
				}
				strChessName := ""
				switch ChessEnum(chess.Type) {
				case CHESS_CAR:
					strChessName = CN_CAR
				case CHESS_HORSE:
					strChessName = CN_HORSE
				case CHESS_CANNON:
					strChessName = CN_CANNON
				case CHESS_ELEPHANT:
					if chess.Color == COLOR_RED {
						strChessName = CN_ELEPHANT_1
					} else {
						strChessName = CN_ELEPHANT_2
					}
				case CHESS_GUARD:
					if chess.Color == COLOR_RED {
						strChessName = CN_GUARD_1
					} else {
						strChessName = CN_GUARD_2
					}
				case CHESS_KING:
					if chess.Color == COLOR_RED {
						strChessName = CN_KING_1
					} else {
						strChessName = CN_KING_2
					}
				case CHESS_PAWN:
					if chess.Color == COLOR_RED {
						strChessName = CN_PAWN_1
					} else {
						strChessName = CN_PAWN_2
					}
				}
				fmt.Print(strColor + strChessName + " ")
			}
		}
		fmt.Println("\n")
	}
}

func (cb *ChessBoard) getRecordKey(record string) (additional, from, op, to string) {
	utf8Chars := []string {}
	for _, v := range record {
		utf8Chars = append(utf8Chars, string(v))
	}
	firstChar := utf8Chars[0]
	if firstChar == ADDITIONAL_FRONT ||
			firstChar == ADDITIONAL_BACK {
		additional = firstChar
		op = utf8Chars[2]
		to = utf8Chars[3]
	} else {
		from = utf8Chars[1]
		op = utf8Chars[2]
		to = utf8Chars[3]
	}
	return
}

func (cb *ChessBoard) getChessRowByCol(chessType ChessEnum, chessColor ChessColor, chessCol int, op string, additional string) (int, int) {
	// dirty fix (1009).
	void := -1
	if additional == ADDITIONAL_NULL {
		// eg: 象７退５. two elephants all in col 7...
		opEnum := OpEnum(op)
		if chessColor == COLOR_RED {
			if chessType != CHESS_CANNON && chessType != CHESS_CAR && chessType != CHESS_HORSE {
				if opEnum == OP_BACKWARD || opEnum == OP_HORIZONTAL {
					for row := 0; row < BOARD_ROW; row++ {
						chess := cb.chessInfo[row][chessCol]
						if chess.Type == chessType && chess.Color == chessColor {
							return row, void
						}
					}
				} else if opEnum == OP_FORWARD {
					for row := BOARD_ROW - 1; row >= 0; row-- {
						chess := cb.chessInfo[row][chessCol]
						if chess.Type == chessType && chess.Color == chessColor {
							return row, void
						}
					}
				}
			} else {
				// cannon is special..
				// dirty fix(1010)... two cannons in the same column but record is confused.
				rows := []int{}
				for row := 0; row < BOARD_ROW; row++ {
					chess := cb.chessInfo[row][chessCol]
					if chess.Type == chessType && chess.Color == chessColor {
						rows = append(rows, row)
					}
				}
				if len(rows) == 1 {
					return rows[0], void
				} else if len(rows) > 1 {
					return rows[0], rows[1]
				}
			}
		} else if chessColor == COLOR_BLACK {
			if chessType != CHESS_CANNON && chessType != CHESS_CAR && chessType != CHESS_HORSE {
				if opEnum == OP_BACKWARD {
					for row := BOARD_ROW - 1; row >= 0; row-- {
						chess := cb.chessInfo[row][chessCol]
						if chess.Type == chessType && chess.Color == chessColor {
							return row, void
						}
					}
				} else if opEnum == OP_FORWARD || opEnum == OP_HORIZONTAL {
					for row := 0; row < BOARD_ROW; row++ {
						chess := cb.chessInfo[row][chessCol]
						if chess.Type == chessType && chess.Color == chessColor {
							return row, void
						}
					}
				}
			} else {
				// cannon is special...
				// dirty fix(1010)... two cannons in the same column but record is confused.
				rows := []int{}
				for row := 0; row < BOARD_ROW; row++ {
					chess := cb.chessInfo[row][chessCol]
					if chess.Type == chessType && chess.Color == chessColor {
						rows = append(rows, row)
					}
				}
				if len(rows) == 1 {
					return rows[0], void
				} else if len(rows) > 1 {
					return rows[0], rows[1]
				}
			}
		}
	} else if additional == ADDITIONAL_FRONT {
		if chessColor == COLOR_RED {
			for row := 0; row < BOARD_ROW; row++ {
				for col := 0; col < BOARD_COL; col++ {
					chess := cb.chessInfo[row][col]
					if chess.Type == chessType && chess.Color == chessColor {
						return row, void
					}
				}
			}
		} else if chessColor == COLOR_BLACK {
			for row := BOARD_ROW - 1; row >= 0; row-- {
				for col := 0; col < BOARD_COL; col++ {
					chess := cb.chessInfo[row][col]
					if chess.Type == chessType && chess.Color == chessColor {
						return row, void
					}
				}
			}
		}
	} else if additional == ADDITIONAL_BACK {
		if chessColor == COLOR_RED {
			for row := BOARD_ROW - 1; row >= 0; row-- {
				for col := 0; col < BOARD_COL; col++ {
					chess := cb.chessInfo[row][col]
					if chess.Type == chessType && chess.Color == chessColor {
						return row, void
					}
				}
			}
		} else if chessColor == COLOR_BLACK {
			for row := 0; row < BOARD_ROW; row++ {
				for col := 0; col < BOARD_COL; col++ {
					chess := cb.chessInfo[row][col]
					if chess.Type == chessType && chess.Color == chessColor {
						return row, void
					}
				}
			}
		}
	}
	cb.DumpForDebug()
	log.Fatalln("[ChessBoard::getChessRowByCol] can't find target chess...")
	return void, void
}

func (cb *ChessBoard) getSpecialChessCol(chessType ChessEnum, chessColor ChessColor) int {
	for col := 0; col < BOARD_COL; col++ {
		count := 0
		for row := 0; row < BOARD_ROW; row++ {
			chess := cb.chessInfo[row][col]
			if chess.Type == chessType && chess.Color == chessColor {
				count++
				if count == 2 {
					return col
				}
			}
		}
	}
	log.Fatalln("ChessBoard::getSpecialChessCol error...")
	return -1
}

func (cb *ChessBoard) getChessCol(chessType ChessEnum, chessColor ChessColor, pos string) int {
	if chessColor == COLOR_RED {
		switch pos {
		case "一":
			return 8
		case "二":
			return 7
		case "三":
			return 6
		case "四":
			return 5
		case "五":
			return 4
		case "六":
			return 3
		case "七":
			return 2
		case "八":
			return 1
		case "九":
			return 0
		case "":
			return cb.getSpecialChessCol(chessType, chessColor)
		default:
			log.Fatalln("record error...")
			return -1
		}
	} else {
		switch pos {
		case "１":
			return 0
		case "２":
			return 1
		case "３":
			return 2
		case "４":
			return 3
		case "５":
			return 4
		case "６":
			return 5
		case "７":
			return 6
		case "８":
			return 7
		case "９":
			return 8
		case "":
			return cb.getSpecialChessCol(chessType, chessColor)
		default:
			log.Fatalln("record error...")
			return -1
		}
	}
}

func (cb *ChessBoard) convertCNDigitToENDigit(chessColor ChessColor, cnDigit string) int {
	if chessColor == COLOR_RED {
		switch cnDigit {
		case "一":
			return 1
		case "二":
			return 2
		case "三":
			return 3
		case "四":
			return 4
		case "五":
			return 5
		case "六":
			return 6
		case "七":
			return 7
		case "八":
			return 8
		case "九":
			return 9
		default:
			log.Fatalln("record error...")
			return -1
		}
	} else {
		switch cnDigit {
		case "１":
			return 1
		case "２":
			return 2
		case "３":
			return 3
		case "４":
			return 4
		case "５":
			return 5
		case "６":
			return 6
		case "７":
			return 7
		case "８":
			return 8
		case "９":
			return 9
		default:
			log.Fatalln("record error...")
			return -1
		}
	}
}

func (cb *ChessBoard) isRowColValid(row, col int) bool {
	if row < 0 || row >= BOARD_ROW {
		return false
	}
	if col < 0 || col >= BOARD_COL {
		return false
	}
	return true
}

func (cb *ChessBoard) checkCarMovement(chessColor ChessColor, oldRow, oldCol, newRow, newCol int, op string) bool {
	if !cb.isRowColValid(newRow, newCol) {
		return false
	}
	targetChess := cb.chessInfo[newRow][newCol]
	if targetChess.Color == chessColor {
		return false
	}
	switch OpEnum(op) {
	case OP_HORIZONTAL: {
		startCol, endCol := 0, 0
		if oldCol < newCol {
			startCol = oldCol + 1
			endCol = newCol
		} else {
			startCol = newCol + 1
			endCol = oldCol
		}
		for i := startCol; i < endCol; i++ {
			if cb.chessInfo[oldRow][i].Type != CHESS_NULL {
				return false
			}
		}
	}
	case OP_FORWARD:
		fallthrough
	case OP_BACKWARD: {
		startRow, endRow := 0, 0
		if oldRow < newRow {
			startRow = oldRow + 1
			endRow = newRow
		} else {
			startRow = newRow + 1
			endRow = oldRow
		}
		for i := startRow; i < endRow; i++ {
			if cb.chessInfo[i][oldCol].Type != CHESS_NULL {
				return false
			}
		}
	}
	default:
		log.Fatalln("[checkCarMovement] invalid car operation...")
		return false
	}
	return true
}

func (cb *ChessBoard) moveCar(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_CAR
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			var oldChess, newChess *Chess

			if !cb.checkCarMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow = oldRow2
				newCol = cb.getChessCol(chessType, chessColor, to)
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow - forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.checkCarMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow = oldRow2 - forwardRow
				newCol = oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow + forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.checkCarMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := oldRow2 + forwardRow
				newCol := oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error car movement...")
		}
	} else {
		chessType := CHESS_CAR
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			var oldChess, newChess *Chess

			if !cb.checkCarMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := oldRow2
				newCol := cb.getChessCol(chessType, chessColor, to)
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow + forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.checkCarMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := oldRow2 + forwardRow
				newCol := oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			backwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow - backwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.checkCarMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := oldRow2 - backwardRow
				newCol := oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error car movement...")
		}
	}
}

func (cb *ChessBoard) checkHorseMovement(chessColor ChessColor, oldRow, oldCol, newRow, newCol int, op string) bool {
	if !cb.isRowColValid(newRow, newCol) {
		return false
	}
	if cb.chessInfo[newRow][newCol].Color == chessColor {
		return false
	}
	blockRow, blockCol := -1, -1
	disRow := oldRow - newRow
	disCol := oldCol - newCol
	if chessColor == COLOR_RED {
		if disRow == 2 {
			blockRow = oldRow - 1
			blockCol = oldCol
		} else if disRow == -2 {
			blockRow = oldRow + 1
			blockCol = oldCol
		} else if disCol == 2 {
			blockRow = oldRow
			blockCol = oldCol - 1
		} else if disCol == -2 {
			blockRow = oldRow
			blockCol = oldCol + 1
		}
	} else {
		if disRow == -2 {
			blockRow = oldRow + 1
			blockCol = oldCol
		} else if disRow == 2 {
			blockRow = oldRow - 1
			blockCol = oldCol
		} else if disCol == 2 {
			blockRow = oldRow
			blockCol = oldCol - 1
		} else if disCol == -2 {
			blockRow = oldRow
			blockCol = oldCol + 1
		}
	}
	if cb.chessInfo[blockRow][blockCol].Type != CHESS_NULL {
		return false
	}

	return true
}

func (cb *ChessBoard) moveHorse(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_HORSE
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			targetCol := cb.getChessCol(chessType, chessColor, to)
			disCol := targetCol - oldCol
			newRow := -1
			newCol := oldCol + disCol
			if disCol == 1 || disCol == -1 {
				newRow = oldRow - 2
			} else if disCol == 2 || disCol == -2 {
				newRow = oldRow - 1
			}

			var oldChess, newChess *Chess

			if !cb.checkHorseMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := -1
				newCol := oldCol + disCol
				if disCol == 1 || disCol == -1 {
					newRow = oldRow2 - 2
				} else if disCol == 2 || disCol == -2 {
					newRow = oldRow2 - 1
				}
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			targetCol := cb.getChessCol(chessType, chessColor, to)
			disCol := targetCol - oldCol
			newRow := -1
			newCol := oldCol + disCol
			if disCol == 1 || disCol == -1 {
				newRow = oldRow + 2
			} else if disCol == 2 || disCol == -2 {
				newRow = oldRow + 1
			} else {
				log.Fatalln("error horse movement...")
			}

			var oldChess, newChess *Chess

			if !cb.checkHorseMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := -1
				newCol := oldCol + disCol
				if disCol == 1 || disCol == -1 {
					newRow = oldRow2 + 2
				} else if disCol == 2 || disCol == -2 {
					newRow = oldRow2 + 1
				}
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error horse movement...")
		}
	} else {
		chessType := CHESS_HORSE
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			targetCol := cb.getChessCol(chessType, chessColor, to)
			disCol := targetCol - oldCol
			newRow := -1
			newCol := oldCol + disCol
			if disCol == 1 || disCol == -1 {
				newRow = oldRow + 2
			} else if disCol == 2 || disCol == -2 {
				newRow = oldRow + 1
			} else {
				log.Fatalln("error horse movement...")
			}

			var oldChess, newChess *Chess

			if !cb.checkHorseMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := -1
				newCol := oldCol + disCol
				if disCol == 1 || disCol == -1 {
					newRow = oldRow2 + 2
				} else if disCol == 2 || disCol == -2 {
					newRow = oldRow2 + 1
				}
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			targetCol := cb.getChessCol(chessType, chessColor, to)
			disCol := targetCol - oldCol
			newRow := -1
			newCol := oldCol + disCol
			if disCol == 1 || disCol == -1 {
				newRow = oldRow - 2
			} else if disCol == 2 || disCol == -2 {
				newRow = oldRow - 1
			} else {
				log.Fatalln("error horse movement...")
			}

			var oldChess, newChess *Chess

			if !cb.checkHorseMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := -1
				newCol := oldCol + disCol
				if disCol == 1 || disCol == -1 {
					newRow = oldRow2 - 2
				} else if disCol == 2 || disCol == -2 {
					newRow = oldRow2 - 1
				}
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error horse movement...")
		}
	}
}

func (cb *ChessBoard) checkCannonMovement(color ChessColor, oldRow, oldCol, newRow, newCol int, op string) bool {
	if !cb.isRowColValid(newRow, newCol) {
		return false
	}
	targetChess := cb.chessInfo[newRow][newCol]
	if targetChess.Color == color {
		return false
	}
	eat := false
	if targetChess.Type != CHESS_NULL {
		eat = true
	}
	chessCount := 0
	switch OpEnum(op) {
	case OP_HORIZONTAL: {
		startCol, endCol := 0, 0
		if oldCol < newCol {
			startCol = oldCol + 1
			endCol = newCol
		} else {
			startCol = newCol + 1
			endCol = oldCol
		}
		for i := startCol; i < endCol; i++ {
			if cb.chessInfo[oldRow][i].Type != CHESS_NULL {
				chessCount++
			}
		}
	}
	case OP_FORWARD:
		fallthrough
	case OP_BACKWARD: {
		startRow, endRow := 0, 0
		if oldRow < newRow {
			startRow = oldRow + 1
			endRow = newRow
		} else {
			startRow = newRow + 1
			endRow = oldRow
		}
		for i := startRow; i < endRow; i++ {
			if cb.chessInfo[i][oldCol].Type != CHESS_NULL {
				chessCount++
			}
		}
	}
	default:
		log.Fatalln("invalid cannon operation...")
	}
	if eat {
		if chessCount != 1 {
			return false
		}
	} else {
		if chessCount != 0 {
			return false
		}
	}
	return true
}

func (cb *ChessBoard) moveCannon(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_CANNON
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			var oldChess, newChess *Chess

			if !cb.isRowColValid(newRow, newCol) || !cb.checkCannonMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := oldRow2
				newCol := cb.getChessCol(chessType, chessColor, to)
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow - forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.isRowColValid(newRow, newCol) || !cb.checkCannonMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow = oldRow2 - forwardRow
				newCol = oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow + forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.isRowColValid(newRow, newCol) || !cb.checkCannonMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow = oldRow2 + forwardRow
				newCol = oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error cannon movement...")
		}
	} else {
		chessType := CHESS_CANNON
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			var oldChess, newChess *Chess

			if !cb.isRowColValid(newRow, newCol) || !cb.checkCannonMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow := oldRow2
				newCol := cb.getChessCol(chessType, chessColor, to)
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			// dirty fix: oldRow2......chinese chess record is confused...
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow + forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.isRowColValid(newRow, newCol) || !cb.checkCannonMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow = oldRow2 + forwardRow
				newCol = oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, oldRow2 := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			forwardRow := cb.convertCNDigitToENDigit(chessColor, to)
			newRow := oldRow - forwardRow
			newCol := oldCol

			var oldChess, newChess *Chess

			if !cb.isRowColValid(newRow, newCol) || !cb.checkCannonMovement(chessColor, oldRow, oldCol, newRow, newCol, op) {
				newRow = oldRow2 - forwardRow
				newCol = oldCol
				oldChess = cb.chessInfo[oldRow2][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			} else {
				oldChess = cb.chessInfo[oldRow][oldCol]
				newChess = cb.chessInfo[newRow][newCol]
			}

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error cannon movement...")
		}
	}
}

func (cb *ChessBoard) moveElephant(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_ELEPHANT
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 2
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 2
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error elephant movement...")
		}
	} else {
		chessType := CHESS_ELEPHANT
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 2
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 2
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error elephant movement...")
		}
	}
}

func (cb *ChessBoard) moveGuard(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_GUARD
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 1
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 1
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error guard movement...")
		}
	} else {
		chessType := CHESS_GUARD
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 1
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 1
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error guard movement...")
		}
	}
}

func (cb *ChessBoard) moveKing(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_KING
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 1
			newCol := oldCol

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 1
			newCol := oldCol

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error king movement...")
		}
	} else {
		chessType := CHESS_KING
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 1
			newCol := oldCol

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_BACKWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 1
			newCol := oldCol

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error king movement...")
		}
	}
}

func (cb *ChessBoard) movePawn(record string, isRed bool) {
	additional, from, op, to := cb.getRecordKey(record)
	if isRed {
		chessType := CHESS_PAWN
		chessColor := COLOR_RED
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow - 1
			newCol := oldCol

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error pawn movement...")
		}
	} else {
		chessType := CHESS_PAWN
		chessColor := COLOR_BLACK
		switch OpEnum(op) {
		case OP_HORIZONTAL: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow
			newCol := cb.getChessCol(chessType, chessColor, to)

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		case OP_FORWARD: {
			oldCol := cb.getChessCol(chessType, chessColor, from)
			oldRow, _ := cb.getChessRowByCol(chessType, chessColor, oldCol, op, additional)
			newRow := oldRow + 1
			newCol := oldCol

			oldChess := cb.chessInfo[oldRow][oldCol]
			newChess := cb.chessInfo[newRow][newCol]

			oldChess.Type = CHESS_NULL
			oldChess.Color = COLOR_NULL
			newChess.Type = chessType
			newChess.Color = chessColor
		}
		default:
			log.Fatalln("error pawn movement...")
		}
	}
}
