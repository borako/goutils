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
  t = "큰 테 스 트"
  fmt.Println("Length of ", t, ": " , len(t))
  t = "원펀맨 리메이크 127화 _ 원펀맨 (리메이크) _ 미니툰 - 만화는 언제나 미니툰 ^._.html"
  fmt.Println("Length of ", t, ": " , len(t))
  for pos, _ := range t {
    fmt.Println("Len: ", pos, ": ", t[:pos])
  }
  fmt.Println(t[:3])
}
