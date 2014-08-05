// wordfrequency
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readWordFromFile(filePath string, splitChars []string) []string {
	var file *os.File
	var err error
	if file, err = os.Open(filePath); err != nil {
		log.Fatal("failed to open the file: ", err)
	}
	defer file.Close()
	words := make([]string, 100)
	lines := make([]string, 50)
	reader := bufio.NewReader(file)
	reader.ReadString('\n')
	return lines
}

func main() {
	fmt.Println("Hello World!")
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	var b = make([]int, 6)
	n1 := copy(b, a[0:])
	fmt.Println(n1, b)
	n2 := copy(b, a[4:])
	fmt.Println(n2, b)
}
