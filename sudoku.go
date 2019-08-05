package main

//GridSize konstanta menunjukan jumlah maksimal matrix
const GridSize = 9

//GridStartIndex kostanta menunjukan index dimulai
const GridStartIndex = 0

//EmptyCell kontanta menunjukan bahwa angka 0 menunjukan kotak kosong
const EmptyCell = 0

//MaxPosibleNumber menunjukan maksimal kemungkinan nilai maksimal yang bisa dimasukan
const MaxPosibleNumber = 9

//MinPosibleNumber menunjukan maksimal kemungkinan nilai minimal yang bisa dimasukan
const MinPosibleNumber = 1

//SubGridSize menunjukan jumlah subsection di dalam Grid
const SubGridSize = 3

/*
 * Pada Fungsi ini akan menampilkan Grid pada console atau termina;
 */
func printToConsole(currGrid *[][]int) {
	for row := GridStartIndex; row < GridSize; row++ {
		for column := GridStartIndex; column < GridSize; column++ {
			print((*currGrid)[row][column], " ")
		}
		println("")
	}
}

/*
 * fungsi isNumberInRow akan memeriksa apakah kemungkinan angka yang dimasukan
 * sudah ada di dalam row yang sama atau tidak
 * jika ada angka yang sama di dalam row maka akan mengembalikan nilai true
 */
func isNumberInRow(row int, posibleNumber int, currGrid *[][]int) bool {
	for column := 0; column < GridSize; column++ {
		if (*currGrid)[row][column] == posibleNumber {
			return true
		}
	}
	return false
}

/*
 * fungsi isNumberInColumn akan memeriksa apakah kemungkinan angka yang dimasukan
 * sudah ada di dalam column yang sama atau tidak
 * jika ada angka yang sama di dalam column maka akan mengembalikan nilai true
 */
func isNumberInColumn(column int, posibleNumber int, currGrid *[][]int) bool {
	for row := 0; row < GridSize; row++ {
		if (*currGrid)[row][column] == posibleNumber {
			return true
		}
	}
	return false
}

/*
 * fungsi isNumberInSubGrid akan memeriksa apakah kemungkinan angka yang dimasukan
 * sudah ada di dalam subGrid yang berukuran 3x3
 * jika ada angka yang sama di dalam subGrid maka akan mengembalikan nilai true
 */
func isNumberInSubGrid(row int, column int, posibleNumber int, currGrid *[][]int) bool {
	startIndexRowSubGrid := row - row%SubGridSize
	endIndexRowSubGrid := startIndexRowSubGrid + SubGridSize

	startIndexColumnSubGrid := column - column%SubGridSize
	endIndexColumnSubGrid := startIndexColumnSubGrid + SubGridSize

	for row := startIndexRowSubGrid; row < endIndexRowSubGrid; row++ {
		for column := startIndexColumnSubGrid; column < endIndexColumnSubGrid; column++ {
			if (*currGrid)[row][column] == posibleNumber {
				return true
			}
		}
	}
	return false
}

/*
 * fungsi isNumberPosible akan memeriksa apakah kemungkinan angka yang dimasukan
 * tidak ada yang sama di row, column maupun subGrid
 * jika tidak ada angka yang sama di dalam row, column maupun subGrid
 * maka akan mengembalikan nilai true
 *
 */
func isNumberPosible(row int, column int, posibleNumber int, currGrid *[][]int) bool {
	return !isNumberInRow(row, posibleNumber, currGrid) &&
		!isNumberInColumn(column, posibleNumber, currGrid) &&
		!isNumberInSubGrid(row, column, posibleNumber, currGrid)
}

/*
 * fungsi solve digunakan untuk menyelesaikan Grid sudoku yang dimasukan
 * pada fungsi ini akan ada parameter bertipe integer array 2D.
 * Pada fungsi ini akan mengembalikan nilai true apabila nilai balikan
 * dari pemanggilan fungsi isNumberPosile bernilai true dan nilai balikan
 * dari pemamnggilan fungsi rekrusif bernilai true juga.
 *
 */
func solve(currGrid *[][]int) bool {
	for row := GridStartIndex; row < GridSize; row++ {
		for column := GridStartIndex; column < GridSize; column++ {
			if (*currGrid)[row][column] == EmptyCell {
				for posibleNumber := MinPosibleNumber; posibleNumber <= MaxPosibleNumber; posibleNumber++ {
					if isNumberPosible(row, column, posibleNumber, currGrid) {
						(*currGrid)[row][column] = posibleNumber

						if !solve(currGrid) {
							(*currGrid)[row][column] = EmptyCell
						} else {
							return true
						}
					}
				}
				return false
			}
		}
	}
	return true
}
