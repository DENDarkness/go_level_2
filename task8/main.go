package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var list = map[string]int64{}
var double = []string{}
var mu sync.RWMutex

// func ReadChan(d chan string) {
// 	for v := range d {
// 		fmt.Println(v)
// 		double = append(double, v)
// 	}
// }

func searchFile(dir string, wg *sync.WaitGroup) {

	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() { // Если это дериктория вызываем в go рутине еще одну функцию для этого каталога.
			wg.Add(1)
			go searchFile(dir+"/"+f.Name(), wg)
		} else {
			//mu.RLock()
			if _, ok := list[f.Name()]; ok { // Проверяем наличие записи в map с таким именем.
				if list[f.Name()] == f.Size() { // Если запись с таким именем есть проверяем размер.
					//fmt.Printf("%s/%s\n", dir, f.Name())
					d := fmt.Sprintf("%s/%s", dir, f.Name())
					fmt.Println(d)
					double = append(double, d) // Если это дубль, отправляем полный путь к дублю в канал.

				}
				//	mu.RUnlock()
			} else {
				// Если записи в map не найдено, записываем в мапу название файл и размер.
				mu.Lock()
				list[f.Name()] = f.Size()
				mu.Unlock()
			}
		}
	}
	wg.Done()
}

func main() {

	var name string
	var wg sync.WaitGroup

	dir := flag.String("p", "~/", "Directory in which to search for duplicates")
	flag.Parse()

	wg.Add(1)
	go searchFile(*dir, &wg)

	wg.Wait()

	for {
		fmt.Print("Удалить дубли?: [Yn] ")
		fmt.Fscan(os.Stdin, &name)
		if name != "Y" || name != "n" || name != "y" || name != "yes" || name != "no" {
			fmt.Print("Введено не коректное значение\n")
			continue
		}
		if name == "Y" || name == "yes" || name == "Yes" || name == "y" {
			for _, v := range double {
				os.Remove(v)
			}
			break
		}
		if name == "n" || name == "no" || name == "No" || name == "N" {
			break
		}
		break
	}

}
