package main

func searchMatrix(matrix [][]int, target int) bool {
	// search row
	l := 0
	r := len(matrix) - 1
	colCount := len(matrix[0])

	possiblerow := -1

	for l <= r {
		mid := (l + r) / 2

		if matrix[mid][0] == target {
			return true
		} else if target > matrix[mid][0] {
			if target == matrix[mid][colCount-1] {
				return true
			}
			if target < matrix[mid][colCount-1] {
				possiblerow = mid
				break
			} else {
				l = mid + 1
			}
		} else if target < matrix[mid][0] {
			r = mid - 1
		}
	}

	if possiblerow == -1 {
		return false
	}

	// binary search on possible row
	l = 0
	r = colCount - 1

	for l <= r {
		mid := (l + r) / 2
		if target == matrix[possiblerow][mid] {
			return true
		} else if target > matrix[possiblerow][mid] {
			l = mid + 1
		} else if target < matrix[possiblerow][mid] {
			r = mid - 1
		}
	}

	return false
}
