package main

import (
  "fmt"
  "path/filepath"
  "log"
  "io/ioutil"
  "os"
  "strings"
)

func isInDelList(fname string) bool {
  var list1 = []string{"test1", "test2"}
  for _, e := range list1 {
    if strings.Contains(e, fname) {
      return true
    }
  }
  return false
}
func processDir(dirname string) {
  // fileList := []string{}
  err := filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
    fmt.Println("Current: ", path)
    return nil
  })
  _ = err

  files, err := ioutil.ReadDir(dirname)
  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    // fmt.Println("Processing DIR: [", dirname, "]", f.Name() + " " +
    //   "Length: ", len(f.Name()))
    if f.IsDir() {
      fmt.Println("Current Dir: ", f.Name())
      if isSecond(f.Name()) {
        fmt.Println("2부")
      }
      // if _, err := os.Stat(f.Name()); err == nil {
      //   processDir(f.Name())
      // }
    }
  }
}

func moveContents(dir1 string, dir2 string) bool {
  // Move all files from dir1 to dir2
}

func find1st(dirname string) string {
  // find first part from current dirname
}

func isSecond(dirname string) bool {
  // Check if this is "2부"

}
func main() {
  processDir("./")
}
