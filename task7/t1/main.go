package main

import (
	"fmt"
	"reflect"
)

func t1(sql string, v ...interface{}) {
	fmt.Println(v)
	var vv []interface{}
	for i := 0; i < len(v); i++ {
		rv := reflect.TypeOf(v[i]).Kind()
		if rv == reflect.Bool {
			vv = append(vv, v[i])
		}
		if rv == reflect.Int {
			vv = append(vv, v[i])
		}
		if rv == reflect.Slice {

		}
	}
	fmt.Println(vv)

}

func main() {

	t1("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)

}
