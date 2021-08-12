package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var list = map[string]int64{}
var dub = []string{}
var mu sync.RWMutex

//var list = map[string]map[string]int64{}

type Files struct {
	Path string
	Name string
	Size int64
}

//var list []Files

func searchFile(dir string) {

	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			go searchFile(dir + "/" + f.Name())
		} else {

			if _, ok := list[f.Name()]; ok {
				if list[f.Name()] == f.Size() {
					fmt.Printf("%s/%s\n", dir, f.Name())

				}
			} else {
				mu.Lock()
				list[f.Name()] = f.Size()
				mu.Unlock()
			}
		}
	}

}

func main() {
	dir := "/home/denys/Public"

	//list := make(map[string]int64)

	searchFile(dir)
	time.Sleep(time.Second * 1)
	//fmt.Println(list)

	//	dulicat(list)

}
