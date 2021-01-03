package main

import "fmt"

func hasUniqueSubSquares(board [][]int) bool {
  subSquares := getSubSquares(board)
  for _, square := range subSquares {
    fmt.Println("Checking sub board .....", square)
    isDup := checkDuplicate(square)
    if isDup == true { return false }
  }
  return true
}

func getSubSquares(board [][]int) [][]int {
  subSquares := make([][]int, 0)
  endIdx := 3
  for i := 0; i < len(board); i += 3 {
    limit := 9
    col := getCols(board, i, endIdx)
    for j := 0; j < len(col); j += 9 {
      subSquares = append(subSquares, col[j:limit])
      limit += 9
    }
    endIdx += 3
  }
  return subSquares
}

func checkDuplicate(arr []int) bool {
  for idx := 0; idx < len(arr)-1; idx++ {
    for i := idx+1; i < len(arr); i++ {
      if arr[idx] == arr[i] {
        return true
      }
    }
  }
  return false
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

func sudokuValid(board [][]int) bool {
  var validBoard bool

  // check if empty 
  if len(board) == 0 { return false }

  allRows := make([]int, 0)
  allCols := make([]int, 0)
  for idx, row := range board {
    allRows = append(allRows, row...)
    IsRowDuplicate := checkDuplicate(row)
    col := getColumn(board, idx)
    IsColumnDuplicate := checkDuplicate(col)
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

func print(board [][]int) {
  for i := 0; i < len(board); i++ {
    if i == 3 || i == 6 {
      fmt.Println(" ")
    }
    for j := 0; j < len(board); j++ {
      if (j == 3) || (j == 6) {
        fmt.Print(" ")
      }
      fmt.Print(board[i][j], " ")
    }
    fmt.Println()
  }
}

func main() {
  board := [][]int{
    {5,3,4,6,7,8,9,1,2},
    {6,7,2,1,9,5,3,4,8},
    {1,9,8,3,4,2,5,6,7},
    {8,5,9,7,6,1,4,2,3},
    {4,2,6,8,5,3,7,9,1},
    {7,1,3,9,2,4,8,5,6},
    {9,6,1,5,3,7,2,8,4},
    {2,8,7,4,1,9,6,3,5},
    {3,4,5,2,8,6,1,7,9},
  }

  /*
  board := [][]int{
    {6,3,4,6,7,8,9,1,2},
    {6,7,2,1,9,5,3,4,8},
    {1,9,8,3,4,2,5,6,7},
    {8,5,9,7,6,1,4,2,3},
    {4,2,6,8,5,3,7,9,1},
    {7,1,3,9,2,4,8,5,6},
    {9,6,1,5,3,7,2,8,4},
    {2,8,7,4,1,9,6,3,5},
    {3,4,5,2,8,6,1,7,9},
  }
  */

  fmt.Println("Sudoku Board")
  print(board)
  fmt.Println()
  isValid := sudokuValid(board)
  if isValid == true {
    fmt.Println("Valid Sudoku Board")
  } else {
    fmt.Println("Invalid Sudoku Board")
  }
}

