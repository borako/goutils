package main
import (
  "fmt"
)

func main() {
  t := "테"
  fmt.Println("Length of ", t, ": " , len(t))
  t = "테스"
  fmt.Println("Length of ", t, ": " , len(t))
  t = "테스트"
  fmt.Println("Length of ", t, ": " , len(t))
  t = "큰테스트"
  fmt.Println("Length of ", t, ": " , len(t))
}
