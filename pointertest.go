package main

import (
	"fmt"
	"strings"
)
var nodelist = []string{".JPG", ".JPEG", ".PY", ".ZIP"}
var jpglist = []string{".JPG", ".JPEG"}

func isInList(f string, l []string) bool {
	for _, e := range l {
		if strings.Contains(e, f) {
			return true
		}
	}
	return false
}

func main() {
	var st = "JPG"
	fmt.Println("Is '", st, "' in list? ", isInList(st, nodelist))
	st = "jpg"
	fmt.Println("Is '", st, "' in list? ", isInList(st, nodelist))
}
