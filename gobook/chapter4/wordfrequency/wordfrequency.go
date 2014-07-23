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
	words := make([]string)
	lines := make([]string)
	reader := bufio.NewReader(file)
	reader.ReadString('\n')
}

func main() {
	fmt.Println("Hello World!")
}
