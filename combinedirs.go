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
  // Grab directories in files
  files, err := ioutil.ReadDir(dirname)
  if err != nil {
    log.Fatal(err)
  }

  checkstr := "2부"
  pos := 0
  for _, f := range files {
    // If "2부"
    if strings.Contains(f.Name(), checkstr) {
      pos = strings.Index(f.Name(), checkstr)
      newdir := f.Name()[:pos] + "1" + f.Name()[pos+1:]
      shortdir := f.Name()[:pos-3]
      fileList := []string{}
      err := filepath.Walk(f.Name(),
        func(path string, f os.FileInfo, err error) error {
          if !f.IsDir() {
            fileList = append(fileList, path)
            newname := newdir + "/" + filepath.Base(path)
            os.Rename(path, newname)
          }
          return nil
      })
      os.Remove(f.Name())
      fmt.Println("Rename '", newdir, "' to: '", shortdir, "'")
      os.Rename(newdir, shortdir)
      _ = err
    }
  }
}

func getTitle(dirname string, checkstr string) string {
  return dirname[:strings.Index(dirname, checkstr)]
}

func main() {
  processDir("./")
}
