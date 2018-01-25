package main

import (
  "fmt"
  "strings"
)

var list1 = []string{"abc", "def"}

func isInList(findstr string) bool {
  fmt.Println("findstr: ", findstr)
  for _, e := range list1 {
    if strings.Contains(e, findstr) {
      return true
    }
  }
  return false
}

func main() {
  fmt.Println(isInList("ab"))
  fmt.Println(isInList("cd"))
  fmt.Println(isInList("ef"))
}
