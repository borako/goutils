package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"path/filepath"
)

var dellist = []string{"CM_B", "comm", "luna", "zzan", "dudd", "jinj", "kyj0",
		"aksg", "maru", "tpgh", "djas", "wjdg", "cogu", "appl", "kiyo", "wlsr",
		"admi", "womw", "afte", "gogo", "brea", "siro", "taka", "hyou", "dudd",
		"icek", "quiz", "rige", "Aki9", "lecl", "mp34", "yong", "k534", "hell",
		"toma", "wodn", "bijo", "tmdg", "lees", "wate", "o2o1", "ffff", "dpsh",
		"afte", "thum", "e692", "ferm", "fire", "ouje", "dark", "sung", "9668",
		"mool", "plqa", "zizo", "flys", "rkdd", "1223", "pkg5", "cust", "leeh",
		"wkdv", "f292", "pare", "wwbs", "wooa", "neob", "hanu", "toad", "dldn",
		"zmwk", "jjya", "star", "saku", "_B", "V1", "L300", "616d", "name", "78_0",
		"22_0", "B_V6", "0_V3", "B_V4", "0_V9", "B_V3", "0_V2", "_V14"}
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

func procDir(dirname string) {
	fmt.Println("##### DIR: ", dirname , " #########")
	err := filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
		base := filepath.Base(path)
		ext := strings.ToUpper(filepath.Ext(f.Name()))
		fmt.Println("Base: ", base, "Ext:", ext)
		if f.IsDir() == false {
			if isInList(base, dellist) {
				fmt.Println("Current (dellist): ", base)
				fmt.Println("Attempting to delete", path)
				err := os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
			}
			if isInList(strings.ToUpper(filepath.Ext(f.Name())), nodelist) == false {
				fmt.Println("Current (nodelist): ", base)
				fmt.Println("Attempting to delete", path)
				err := os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
			}
			if ext == "" {
				fmt.Println("Current (nil ext): ", base)
				fmt.Println("Attempting to delete", path)
				err := os.Remove(path)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	// processDir("./")
	procDir("./")
}
