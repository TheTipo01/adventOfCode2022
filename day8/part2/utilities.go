package main

func checkIfPreviousAreLess(row []int, column int) int {
	cont := 0

	for i := column - 1; i >= 0; i-- {
		cont++

		if row[i] >= row[column] {
			break
		}
	}

	return cont
}

func checkIfSubsequentAreLess(row []int, column int) int {
	cont := 0

	for i := column + 1; i < len(row); i++ {
		cont++

		if row[i] >= row[column] {
			break
		}
	}

	return cont
}

func checkIfPreviousAreLessInColumn(numbers [][]int, row, column int) int {
	cont := 0

	for i := row - 1; i >= 0; i-- {
		cont++

		if numbers[i][column] >= numbers[row][column] {
			break
		}
	}

	return cont
}

func checkIfSubsequentAreLessInColumn(numbers [][]int, row, column int) int {
	cont := 0

	for i := row + 1; i < len(numbers); i++ {
		cont++

		if numbers[i][column] >= numbers[row][column] {
			break
		}

	}

	return cont
}

func check(numbers [][]int, row, column int) int {
	left, right, up, down := checkIfPreviousAreLess(numbers[row], column), checkIfSubsequentAreLess(numbers[row], column),
		checkIfPreviousAreLessInColumn(numbers, row, column), checkIfSubsequentAreLessInColumn(numbers, row, column)

	// fmt.Printf("Tree at position %d,%d: %d, %d, %d, %d\n", row, column, left, right, up, down)

	return left * right * up * down
}
