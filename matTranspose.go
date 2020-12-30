package main

import "fmt"

func print(board [][]int) {
  for i := 0; i < len(board); i++ {
    for j := 0; j < len(board); j++ {
      fmt.Print(board[i][j], " ")
    }
    fmt.Println()
  }
}

func GetColumn(idx int, mat [][]int) (col []int) {
  for _, r := range mat {
    col = append(col, r[idx])
  }
  return
}

func Transpose(mat [][]int) (result [][]int) {
  for idx, _ := range mat {
    col := GetColumn(idx, mat)
    result = append(result, col)
  }
  return result
}

func main() {
  mat := [][]int{
    {1,2,3},
    {4,5,6},
    {7,8,9},
  }

  print(mat)
  ret := Transpose(mat)

  fmt.Println("Tranpose")
  print(ret)

}
