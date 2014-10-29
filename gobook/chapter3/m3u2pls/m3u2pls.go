package main

import (
	"bufio"
	"fmt"
	"io"
	_ "io/ioutil"
	_ "log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inArgs := os.Args
	fmt.Println(inArgs)
	if len(inArgs) < 2 || !strings.HasSuffix(inArgs[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(inArgs[0]))
		//fmt.Printf("usage: %s <file.m3u>\n", inArgs[0])
		os.Exit(1)
	}
	for idx, val := range inArgs {
		fmt.Printf("inArgs[%d] = %s\n", idx, val)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please input song name:")
		// You had better not use ReadLine function...
		//buf, isPrefix, err := reader.ReadLine()
		//line := string(buf)
		//fmt.Println("isPrefix = ", isPrefix)
		line, err := reader.ReadString('\n')
		if err == nil || err != io.EOF {
			fmt.Printf("You have input:%s\n", line)
			if strings.TrimSpace(line) == "exit" {
				break
			}
		} else {
			break
		}
	}

	fmt.Println("----------------")

}
