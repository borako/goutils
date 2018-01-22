package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
)

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
