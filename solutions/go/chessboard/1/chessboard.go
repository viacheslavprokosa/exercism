package chessboard

type File []bool
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	squares, exists := cb[file]
	var count int
	if !exists {
		return 0
	}
	for _, square := range squares {
		if square {
			count++
		}
	}
	return count
}

func CountInRank(cb Chessboard, rank int) int {
	var count int
	if rank < 1 || rank > 8 {
		return 0
	}
	for file := range cb {
		if cb[file][rank-1] {
			count++
		}
	}
	return count
}

func CountAll(cb Chessboard) int {
	for file := range cb {
		return len(cb[file]) * len(cb)
	}
	return 0
}

func CountOccupied(cb Chessboard) int {
	var count int
	for file := range cb {
		count += CountInFile(cb, file)
	}
	return count
}