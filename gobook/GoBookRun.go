package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func init() {
	fmt.Println("Initializing the GoBookRun ...")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[i:]
	}
	return file
}

func main() {
	fmt.Println("Hello Go Book!!!")
	slice := []string{"aaa", "bbb", "ccc", "ddd"}
	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))
	fmt.Println(slice)
	slice = slice[1:2]
	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))
	fmt.Println(slice)
	var i interface{} = 99
	var s interface{} = []string{"left", "right"}
	j := i.(int)
	fmt.Printf("%T -> %d\n", j, j)
	if i, ok := i.(int); ok {
		fmt.Printf("%T -> %d\n", i, i)
		i = 100
		fmt.Printf("in if:%T -> %d\n", i, i)
		fmt.Printf("before: ok = %t\n", ok)
		ok = false
		fmt.Printf("after: ok = %t\n", ok)
	}
	fmt.Printf("out if:%T -> %d\n", i, i)
	if s, ok := s.([]string); ok {
		fmt.Printf("ok = %t\n", ok)
		fmt.Printf("%T -> %q\n", s, s)
	}
	suffix := Suffix("test.tar.gz")
	fmt.Println(suffix)
}
