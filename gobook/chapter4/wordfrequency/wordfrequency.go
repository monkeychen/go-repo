// wordfrequency
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func CalculateWordFromFile(filePath string) (wordCountMap map[string]int) {
	var file *os.File
	var err error
	if file, err = os.Open(filePath); err != nil {
		log.Fatal("failed to open the file: ", err)
	}
	defer file.Close()
	wordCountMap = make(map[string]int)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if utf8.RuneCountInString(word) > 1 {
				wordCountMap[word] += 1
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("failed to finish reading the file: ", err)
		}
	}
	return
}

func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
	return strings.FieldsFunc(s, notALetter)
}

func reportByWords(frequencyForWord map[string]int) {
	words := make([]string, 0, len(frequencyForWord))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range frequencyForWord {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			frequencyForWord[word])
	}
}

func invertStringIntMap(intForString map[string]int) map[int][]string {
	stringsForInt := make(map[int][]string, len(intForString))
	for key, value := range intForString {
		stringsForInt[value] = append(stringsForInt[value], key)
	}
	return stringsForInt
}

func reportByFrequency(wordsForFrequency map[int][]string) {
	frequencies := make([]int, 0, len(wordsForFrequency))
	for frequency := range wordsForFrequency {
		frequencies = append(frequencies, frequency)
	}
	sort.Ints(frequencies)
	width := len(fmt.Sprint(frequencies[len(frequencies)-1]))
	fmt.Println("Frequency → Words")
	for _, frequency := range frequencies {
		words := wordsForFrequency[frequency]
		sort.Strings(words)
		fmt.Printf("%*d %s\n", width, frequency, strings.Join(words, ", "))
	}
}

func main() {
	fmt.Println("Hello World!")
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	var b = make([]int, 6)
	n1 := copy(b, a[0:])
	fmt.Println(n1, b)
	n2 := copy(b, a[4:])
	fmt.Println(n2, b)
	result := CalculateWordFromFile("ump.log")
	reportByWords(result)
	wordsForFrequency := invertStringIntMap(result)
	reportByFrequency(wordsForFrequency)
}
