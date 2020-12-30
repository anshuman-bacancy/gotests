package main

import "fmt"

func Max(a, b uint8) uint8 {
  if a > b { return a } else { return b }
}

func jump(jumps []uint8) bool {
  var i, nextReachableIdx uint8 = 0, 0
  var jumpsLength uint8 = uint8(len(jumps));

  for ; i < jumpsLength; i++ {
    if i > nextReachableIdx {
      return false
    } else {
      nextReachableIdx = Max(nextReachableIdx, i+jumps[i])
    }
  }
  return true
}

func main() {
  jumps := []uint8 {1,1,2,3,2,1,0,0,1,3}
  ret := jump(jumps)
  fmt.Println(ret)
}
