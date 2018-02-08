package main

import (
  // "fmt"
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
  // err := filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
  //   fmt.Println("Current: ", path)
  //   return nil
  // })
  // _ = err

  files, err := ioutil.ReadDir(dirname)
  if err != nil {
    log.Fatal(err)
  }

  checkstr := "2부"
  // title := ""
  pos := 0
  for _, f := range files {
    // fmt.Println("Processing DIR: [", dirname, "]", f.Name() + " " +
    //   "Length: ", len(f.Name()))
    if strings.Contains(f.Name(), checkstr) {
      pos = strings.Index(f.Name(), checkstr)
      // title = f.Name()[:pos]
      // fmt.Println("*** Title: ", title[:len(title)-3])
      newdir := f.Name()[:pos] + "1" + f.Name()[pos+1:]
      fileList := []string{}
      err := filepath.Walk(f.Name(),
        func(path string, f os.FileInfo, err error) error {
          fileList = append(fileList, path)
          // os.Rename(path, newname)
          // fmt.Println("Current: ", filepath.Base(path))
          // fmt.Println("New: ", newdir + filepath.Base(path))
          return nil
      })
      _ = err
      for _, s := range fileList {
        newname := newdir + "/" + filepath.Base(s)
        os.Rename(s, newname)
        // fmt.Println("Rename '", s, "'")
        // fmt.Println("To: '", newname , "'")
      }
    }
  }
}

func moveContents(dir1 string, dir2 string) bool {
  // Move all files from dir1 to dir2
  return true
}

func find1st(dirname string) string {
  // find first part from current dirname
  return ""
}

func getTitle(dirname string, checkstr string) string {
  return dirname[:strings.Index(dirname, checkstr)]
}

func isSecond(dirname string, checkstr string) bool {
  // Check if this is "2부"
  return strings.Contains(dirname, checkstr)

}
func main() {
  processDir("./")
}
