package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strings"
)

func isInDelList(fname string) bool {
  var list1 = []string{"test1", "test2"}
  for _, e := range list1 {
    if strings.Contains(e, findstr) {
      return true
    }
  }
  return false
}
func processDir(dirname string) {
  files, err := ioutil.ReadDir(dirname)
  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    fmt.Println("Processing DIR: [", dirname, "]", f.Name() + " " +
      "Length: ", len(f.Name()))
    if f.IsDir() {
      if _, err := os.Stat(f.Name()); err == nil {
        processDir(f.Name())
      }
    }
  }
}

func main() {
  processDir("./")
}
