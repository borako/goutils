package main

import (
	// "fmt"
	"os"
)

func main() {
	t2 := "t2/test2"
	t1 := "t1/test2"
	os.Rename(t2, t1)
}
