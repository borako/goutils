package main
import (
  "fmt"
  "strings"
  "golang.org/x/text/search"
  "golang.org/x/text/language"
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
  if strings.Contains(t, "화") {
    fmt.Println("Yes")
  }
  m := search.New(language.Korean, search.IgnoreCase)
  i, _ := m.IndexString(t, "화")
  fmt.Println("Index: ", i)
  fmt.Println(len("화"))
  fmt.Println(t[:i+6])
}
