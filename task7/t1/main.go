package main

import (
	"fmt"
	"reflect"
	"strings"
)

func convValuesToSlice(s string, v ...interface{}) (string, []interface{}) {

	var result []interface{}
	f := strings.Split(s, "?")

	for i := 0; i < len(v); i++ {
		switch reflect.TypeOf(v[i]).Kind() {
		case reflect.Slice, reflect.Array:
			sv := reflect.ValueOf(v[i])
			for z := 0; z < sv.Len(); z++ {
				result = append(result, sv.Index(z).Interface())
				if z == sv.Len()-1 {
					f[i] = f[i] + "?"
				} else {
					f[i] = f[i] + "?,"
				}
			}
		default:
			f[i] = f[i] + "?"
			result = append(result, v[i])
		}
	}
	res := strings.Join([]string{f[0], f[1], f[2]}, "")

	return res, result

}

func main() {
	//"SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?",
	str, value := convValuesToSlice("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)
	fmt.Println(str, value)
}
