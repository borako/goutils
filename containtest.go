package main

import (
  "fmt"
  // "path/filepath"
  // "log"
  // "io/ioutil"
  // "os"
  "strings"
)

func main() {
  cs := "2부"
  fmt.Println(strings.Contains("테스트", "테"))
  fmt.Println(strings.Contains("테스트", "나"))
  fmt.Println(strings.Contains("우리의 사랑 2부", cs))
  cs2 := "2부"
  ts := "바라카몬 13권 - 2부 _ 바라카"
  fmt.Println(strings.Contains("바라카몬 13권 - 2부 _ 바라카", cs2))
  fmt.Println(strings.Index("바라카몬 13권 - 2부 _ 바라카", cs2))
  fmt.Println(ts[:strings.Index(ts, cs2)])
  fmt.Println(len(""))
}
