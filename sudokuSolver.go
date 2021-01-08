package main

import "fmt"

//HELPER FUNCTIONS
func hasUniqueSubSquares(board *[][]int) bool {
  subSquares := getSubSquares(board)
  for _, square := range subSquares {
    //fmt.Println("Checking sub board .....", square)
    isDup := hasDuplicates(square)
    if isDup == true { return false }
  }
  return true
}

func getSubSquares(board *[][]int) [][]int {
  subSquares := make([][]int, 0)
  endIdx := 3
  for i := 0; i < len(*board); i += 3 {
    limit := 9
    col := getCols(*board, i, endIdx)
    for j := 0; j < len(col); j += 9 {
      subSquares = append(subSquares, col[j:limit])
      limit += 9
    }
    endIdx += 3
  }
  return subSquares
}

func getCols(board[][]int, startIdx, endIdx int) (cols []int) {
  cols = make([]int, 0)
  for _, row := range board {
    cols = append(cols, row[startIdx:endIdx]...)
  }
  return
}

func getColumn(board [][]int, columnIndex int) (column []int) {
  column = make([]int, 0)
  for _, row := range board {
    column = append(column, row[columnIndex])
  }
  return
}

func hasDuplicates(row []int) bool {
  for i := 0; i < len(row)-1; i++ {
    for j := i+1; j < len(row); j++ {
      if row[i] == row[j] {
        return true
      }
    }
  }
  return false
}

/*
func isValidBoard(board *[][]int) bool {
  var validBoard bool

  if len(*board) == 0 { return false }

  allRows := make([]int, 0)
  allCols := make([]int, 0)
  for idx, row := range *board {
    allRows = append(allRows, row...)
    IsRowDuplicate := hasDuplicates(row)
    col := getColumn(*board, idx)
    IsColumnDuplicate := hasDuplicates(col)
    allCols = append(allCols, col...)

    if IsRowDuplicate == true || IsColumnDuplicate == true {
      validBoard = false
    } else {
      validBoard = true
    }
  }

  UniqueSubBoards := hasUniqueSubSquares(board)
  if UniqueSubBoards == true {
    validBoard = true
  } else {
    validBoard = false
  }

  return validBoard
}
*/


func isValidBoard(board *[][]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := make([]int, 10, 10)
		for col := 0; col < 9; col++ {
			counter[(*board)[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := make([]int, 10, 10)
		//counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[(*board)[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
      counter := make([]int, 10, 10)
			//counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[(*board)[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}


func containsZero(board *[][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*board)[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

//SOLVE
func solve(board *[][]int) bool {
  if !containsZero(board) {
    return true
	}
  for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*board)[i][j] == 0 {
				for valToAdd := 9; valToAdd >= 1; valToAdd-- {
					(*board)[i][j] = valToAdd
					if isValidBoard(board) {
						if solve(board) {
							return true
						}
						(*board)[i][j] = 0
					} else {
						(*board)[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}


// CHEATING!
func processBoard(board map[int][]int) (newBoard[][]int) {
  newBoard = make([][]int, 0)
  for i := 1; i <= len(board); i++ {
    newBoard = append(newBoard, board[i])
  }
  return
}

//DISPLAY
func print(board *[][]int) {
  for i := 0; i < len((*board)); i++ {
    if i == 3 || i == 6 {
      fmt.Println("--------------------")
    }
    for j := 0; j < len((*board)); j++ {
      if j == 3 || j == 6 {
        fmt.Print("|")
      }
        fmt.Print((*board)[i][j], " ")
    }
    fmt.Println()
  }
}

func main() {
  board := map[int][]int {
    1: {0,0,0,1,0,5,0,6,8},
    2: {0,0,0,1,0,5,7,0,1},
    3: {9,0,1,0,0,0,0,3,0},
    4: {0,0,7,0,2,6,0,0,0},
    5: {5,0,0,0,0,0,0,0,3},
    6: {0,0,0,8,7,0,4,0,0},
    7: {0,3,0,0,0,0,8,0,5},
    8: {1,0,5,0,0,0,0,6,8},
    9: {7,9,0,4,0,1,0,0,0},
  }

  betterBoard := processBoard(board)
  fmt.Println("Sudoku Board \n0's -> Empty spaces... \n")
  print(&betterBoard)

  if solve(&betterBoard) {
    fmt.Println("Solved... :)")
    print(&betterBoard)
  } else {
    fmt.Println("Not solved... :( \n")
  }
}
