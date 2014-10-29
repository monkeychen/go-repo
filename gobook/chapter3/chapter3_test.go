// chapter3_test.go
package chapter3

import (
	"bufio"
	"fmt"
	"gobook/chapter3/stringbase"
	"io"
	_ "io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestIntForBool(t *testing.T) {
	res := main.IntForBool(true)
	fmt.Println("Hello World!", res)
}

func TestSplit(t *testing.T) {
	names := "hello.wor ld.hello.go.simiam.work"
	fmt.Print("|")
	for idx, name := range strings.Split(names, ".") {
		fmt.Printf("[%d->%s]|", idx, name)
	}
	fmt.Println()
	fmt.Print("|")
	for idx, name := range strings.Split(names, "") {
		fmt.Printf("[%d->%s]|", idx, name)
	}
	fmt.Println()
}

func TestSplitN_After(t *testing.T) {
	names := "hi*go,hello.wor ld.hello.go...simiam..work"
	fmt.Print("|")
	for idx, name := range strings.SplitN(names, ".", 3) {
		fmt.Printf("[%d->%s]|", idx, name)
	}
	fmt.Println()
	fmt.Print("|")
	for idx, name := range strings.SplitN(names, ".", 0) {
		fmt.Printf("[%d->%s]|", idx, name)
	}
	fmt.Println()
	fmt.Print("|")
	for idx, name := range strings.SplitN(names, ".", -1) {
		fmt.Printf("[%d->%s]|", idx, name)
	}
	fmt.Println()
	fmt.Print("|")
	for idx, name := range strings.SplitAfter(names, ".") {
		fmt.Printf("[%d->%s]|", idx, name)
	}
	fmt.Println()
	// same as afterN

	for idx, name := range strings.SplitAfter(names, "hello") {
		fmt.Printf("#[%d->%s]#\n", idx, name)
		var sli interface{}
		sli = strings.FieldsFunc(name, func(char rune) bool {
			switch char {
			case '.', '*':
				return true
			}
			return false
		})
		if sli, ok := sli.([]string); ok {
			fmt.Print("类型断言》|")
			for idx, name := range sli {
				fmt.Printf("[%d->%s]|", idx, name)
			}
			fmt.Println()
		}

	}
	fmt.Println()

}

func TestReplace(t *testing.T) {
	name := "hello\t\tworld\t,hi golan\t!"
	fmt.Println("before:", name)
	after := strings.Replace(name, "\t", "_", 2)
	fmt.Println("1.st->name:", name)
	fmt.Println("1.st->after:", after)
	after = strings.Replace(name, "\t", "_", -1)
	fmt.Println("2.nd->name:", name)
	fmt.Println("2.nd->after:", after)
}

func TestStringApi(t *testing.T) {
	str := "khello golang, hello world!"
	res := strings.Contains(str, "hello")
	fmt.Println("res = ", res)
	fmt.Println(strings.EqualFold("aAa", "Aaa"))
	fmt.Println("bbb" == "Bbb")
	var sli interface{}
	sli = strings.Fields("a b	c  d")
	arr, ok := sli.([]string)
	fmt.Println("Is slice:", ok)
	fmt.Println(ok, sli, arr[0], arr[1], arr[2], arr[3])
	fmt.Println(len("aaa"))
	fmt.Println(strings.Index(str, "golang"))
	fmt.Println(strings.IndexAny(str, "goang"))
	fmt.Println(strings.IndexRune(str, 'w'))
	fmt.Println(strings.Join([]string{"a", "b", "c", "d"}, "___"))
	mf := func(old rune) rune {
		if old == 65 {
			return 97
		}
		return old
	}
	fmt.Println(strings.Map(mf, "AbcdA"))
	fmt.Println("[" + main.SimplifyWhitespace("   aaa bbb ccc    dddd     ") + "]")
	asciiOnly := func(char rune) rune {
		if char > 127 {
			return -1
		}
		return char
	}
	fmt.Println(strings.Map(asciiOnly, "そそstそそそrings.Map() test!!!せせぬめつ"))
	reader := strings.NewReader("そそstそそそrings")
	for {
		char, size, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Printf("%U '%c' %d: % X\n", char, char, size, []byte(string(char)))
	}
}

func TestStrconvApi(t *testing.T) {
	res := "test"
	fmt.Println(res)
	res = strconv.Quote(res)
	fmt.Println(res)
}

func TestUtf8Api(t *testing.T) {
	fmt.Println("中国，len=", len("中国"))
	fmt.Println([]byte("中国"), len([]byte("中国")))
	fmt.Println(utf8.RuneCount([]byte("中国")))
	fmt.Println(unicode.IsDigit('c'))
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit, 'a'))
	fmt.Println(unicode.Is(unicode.ASCII_Hex_Digit, 'k'))
	names := []string{"Mic.Jordon.steve", "steve.job.davi"}
	fmt.Println(names)
	nameRex := regexp.MustCompile(`(\pL+\.?(?:\s+\pL+\.?)*)\s+(\pL+)`)
	for i := 0; i < len(names); i++ {
		names[i] = nameRex.ReplaceAllString(names[i], "${2}, ${1}")
	}
	fmt.Println(names)
	lines := "language : golang "
	valueForKey := make(map[string]string)
	keyValueReg := regexp.MustCompile(`\s*([[:alpha:]]\w*)\s*:\s*(.+)`)
	if matches := keyValueReg.FindAllStringSubmatch(lines, -1); matches != nil {
		for _, match := range matches {
			valueForKey[match[2]] = strings.TrimRight(match[1], "\t ")
		}
	}
	fmt.Println(valueForKey)
}

func TestLog(t *testing.T) {
	if !FileExists("logfile") {
		CreateFile("logfile")
	}
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	defer f.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// attempt #1
	log.SetOutput(io.MultiWriter(os.Stderr, f))
	log.Println("hello, logfile")

	// attempt #2
	log.SetOutput(io.Writer(f))
	log.Println("hello, logfile")

	// attempt #3
	log.SetOutput(f)
	log.Println("hello, logfile")
}

func TestLogger(t *testing.T) {
	if !FileExists("loggerfile") {
		CreateFile("loggerfile")
	}
	f, err := os.OpenFile("loggerfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	defer f.Close()
	var buf []byte
	n, err := f.Read(buf)
	reader := bufio.NewReader(f)
	line, err := reader.ReadString('\n')
	fmt.Printf("Read from file[%s],content=[%s]\n", f.Name(), line)
	fmt.Println("n=", n, "err = ", err)
	writter := bufio.NewWriter(f)
	writter.WriteString("This line was written by bufio.write()...\n")
	writter.Flush()

	logger1 := log.New(f, "Logger.test.", log.LstdFlags|log.Lshortfile)
	logger1.Println("log test ...")
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
