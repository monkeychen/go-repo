package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func showHelp() {
	fmt.Printf("usage: %s [-b|--bar] <whole-number>\n", filepath.Base(os.Args[0]))
	fmt.Println("-b --bar draw an underbar and an overbar")
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 || (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help")) {
		showHelp()
	}
	isShowBar := false
	stringOfDigits := os.Args[1]
	if len(os.Args) > 2 && (os.Args[1] == "-b" || os.Args[1] == "--bar") {
		stringOfDigits = os.Args[2]
		isShowBar = true
	}

	strDigitLen := len(stringOfDigits)
	if isShowBar {
		overbar := strings.Repeat("*", strDigitLen*7)
		fmt.Println(overbar)
	}
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row]
			} else {
				log.Fatal("Invalid whole number !!!")
			}
		}
		fmt.Println(line)
	}
	if isShowBar {
		underbar := strings.Repeat("*", strDigitLen*7)
		fmt.Println(underbar)
	}
}

var bigDigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{
		"   1   ",
		"  11   ",
		"   1   ",
		"   1   ",
		"   1   ",
		"   1   ",
		"  111  ",
	},
	{
		"  222  ",
		" 2   2 ",
		"    2  ",
		"   2   ",
		"  2    ",
		" 2     ",
		" 22222 ",
	},
	{
		"  333  ",
		" 3   3 ",
		"     3 ",
		"   33  ",
		"     3 ",
		" 3   3 ",
		"  333  ",
	},
	{
		"    4  ",
		"   44  ",
		"  4 4  ",
		" 4  4  ",
		"4444444",
		"    4  ",
		"    4  ",
	},
	{
		" 55555 ",
		" 5     ",
		" 5     ",
		" 5555  ",
		"     5 ",
		" 5   5 ",
		"  555  ",
	},
	{
		"  666  ",
		" 6   6 ",
		" 6     ",
		" 6666  ",
		" 6   6 ",
		" 6   6 ",
		"  666  ",
	},
	{
		" 777777",
		"      7",
		"     7 ",
		"    7  ",
		"   7   ",
		"  7    ",
		" 7     ",
	},
	{
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
	},
	{
		"  999  ",
		" 9   9 ",
		" 9   9 ",
		"  9999 ",
		"     9 ",
		" 9   9 ",
		"  999  ",
	},
}
