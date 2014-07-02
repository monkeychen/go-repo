package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ex1() {
	æs := ""
	for _, char := range []rune{'æ', 0xE6, 0346, 230, '\xE6', '\u00E6'} {
		fmt.Printf("[0x%X '%c'] ", char, char)
		æs += string(char)
	}
	fmt.Println(æs)
}

func ex2() {
	phrase := "vætt og tёгддt"
	fmt.Printf("string:\"%s\"\n", phrase)
	fmt.Println("index rune char bytes")
	for index, char := range phrase {
		fmt.Printf("%-2d	%U	'%c'	%X\n", index, char, char, []byte(string(char)))
	}
}

func ex3() {
	line := "aæbc gæo gæo	gooægle"
	i := strings.Index(line, " ")
	firstWorld := line[:i]
	j := strings.LastIndex(line, " ")
	lastWorld := line[j+1:]
	fmt.Println(firstWorld, lastWorld)
}

func ex4() {
	line := "aæbc gæo gæo	gooægle"
	i := strings.IndexFunc(line, unicode.IsSpace)
	firstWorld := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	_, size := utf8.DecodeRuneInString(line[j:])
	lastWorld := line[j+size:]
	fmt.Println(firstWorld, lastWorld)
}

func IntForBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}
	return s
}

func ex5() {
	fmt.Printf("%t %t \n", true, false)
	fmt.Printf("%d %d \n", IntForBool(true), IntForBool(false))
	fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)
	fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41, 41, 41, 41, -41)
	i := 3931
	fmt.Printf("|%x|%X|%8x|%8X|%#04x|%#04X|0x%04X|\n", i, i, i, i, i, i, i)
	i = 569
	fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i, i, i, Pad(i, 6, '*'))
}

func main() {
	fmt.Println("string base examples ...")
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
}
