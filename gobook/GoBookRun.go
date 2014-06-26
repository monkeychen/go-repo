package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initializing the GoBookRun ...")
}

func main() {
	fmt.Println("Hello Go Book!!!")
	slice := []string{"aaa", "bbb", "ccc", "ddd"}
	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))
	fmt.Println(slice)
	slice = slice[1:2]
	fmt.Printf("len = %d, cap = %d\n", len(slice), cap(slice))
	fmt.Println(slice)
}
