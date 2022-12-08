package main

func checkIfPreviousAreLess(row []int, column int) bool {
	for i := 0; i < column; i++ {
		if row[i] >= row[column] {
			return false
		}
	}

	return true
}

func checkIfSubsequentAreLess(row []int, column int) bool {
	for i := column + 1; i < len(row); i++ {
		if row[i] >= row[column] {
			return false
		}
	}

	return true
}

func checkIfPreviousAreLessInColumn(numbers [][]int, row, column int) bool {
	for i := 0; i < row; i++ {
		if numbers[i][column] >= numbers[row][column] {
			return false
		}
	}

	return true
}

func checkIfSubsequentAreLessInColumn(numbers [][]int, row, column int) bool {
	for i := row + 1; i < len(numbers); i++ {
		if numbers[i][column] >= numbers[row][column] {
			return false
		}
	}

	return true
}

func check(numbers [][]int, row, column int) bool {
	return checkIfPreviousAreLess(numbers[row], column) ||
		checkIfSubsequentAreLess(numbers[row], column) ||
		checkIfPreviousAreLessInColumn(numbers, row, column) ||
		checkIfSubsequentAreLessInColumn(numbers, row, column)
}
