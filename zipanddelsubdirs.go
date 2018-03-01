package main

import (
  "fmt"
  "path/filepath"
  "log"
  "io/ioutil"
  "io"
  "os"
  // "strings"
  "archive/zip"
)

func processDir(dirname string) error {
  // Grab directories in files
  files, err0 := ioutil.ReadDir(dirname)
  if err0 != nil {
    log.Fatal(err0)
  }

  for _, f := range files {
    if !f.IsDir() {
      continue
    }
    fmt.Println("Processing:", f.Name())
    zipname := f.Name() + ".zip"
    if _, err := os.Stat(zipname); err == nil {
      fmt.Println("Zipfile '", zipname , "' exists. Deleting...'")
      os.Remove(zipname)
    }
    zipfile, err1 := os.Create(zipname)
    if err1 != nil {
      return err1
    }
    defer zipfile.Close()
    zipWriter := zip.NewWriter(zipfile)
    defer zipWriter.Close()
    err := filepath.Walk(f.Name(),
      func(path string, fn os.FileInfo, err error) error {
        if !fn.IsDir() {
          // fmt.Println("Processing file: ", path)
          // Add files to zip
          curfile, err := os.Open(path)
          if err != nil {
              return err
          }
          defer curfile.Close()
          // Get the file information
          info, err := curfile.Stat()
          if err != nil {
              return err
          }
          header, err := zip.FileInfoHeader(info)
          if err != nil {
              return err
          }
          // Change to deflate to gain better compression
          // see http://golang.org/pkg/archive/zip/#pkg-constants
          header.Method = zip.Deflate
          writer, err := zipWriter.CreateHeader(header)
          if err != nil {
              return err
          }
          _, err = io.Copy(writer, curfile)
          if err != nil {
              return err
          }
        }
        return nil
    })
    _ = err
  }
  return nil
}

func main() {
  processDir("./")
}
