package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func wrapError(e interface{}) error {
	t := time.Now()
	return fmt.Errorf("%v error: %v", t.Format("2006-01-02 15:04:05"), e)
}

func genNumPassword(l int, s []byte) string {

	if l > 24 {
		panic("too many characters")
	}

	var r []byte

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < l; i++ {
		r = append(r, s[rand.Intn(len(s))])
	}

	return string(r)
}

func writePasswordToFile(p string) error {
	f, err := os.Create("password.txt")
	if err != nil {
		return wrapError(err)
	}
	defer f.Close()

	f.WriteString(p)

	return nil
}

func main() {

	s := flag.Int("s", 24, "Password size.")
	// TODO
	// t := flag.Int("t", 1, "Password type.\n 1 - only numbers.\n 2 - numbers and letters of upper and lower case.")

	flag.Parse()

	num := []byte{
		48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
		97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
		65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	}

	defer func(n []byte) {
		if v := recover(); v != nil {
			fmt.Println(wrapError(v))
			fmt.Println("reduce the number of characters to 24")
			fmt.Println(genNumPassword(24, n[:]))
		}
	}(num)

	pass := genNumPassword(*s, num[:])

	fmt.Println(pass)

	if err := writePasswordToFile(pass); err != nil {
		wrapError(err)
	}

}
