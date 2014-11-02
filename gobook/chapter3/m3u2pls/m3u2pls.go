package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func apiDemo() {
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
	type FileMode uint32

	const (
		ModeDir FileMode = 1 << (32 - 1 - iota)
		TestB
		TestC
		TestD
		TestE FileMode = 15
	)
	fmt.Printf("%33b\n", ModeDir)
	fmt.Printf("%33b\n", TestB)
	fmt.Printf("%33b\n", TestC)
	fmt.Printf("%33b\n", TestD)
	fmt.Printf("%33b\n", TestE)
	var i uint8 = 2
	fmt.Printf("%b\n", i)
	i = 100
	fmt.Printf("%b,%#x\n", i, i)
}

type M3uEntry struct {
	Index  uint32
	Name   string
	Path   string
	Length uint32
}

func (entry M3uEntry) String() string {
	return fmt.Sprintf("File%d=%s\nTitle%d=%s\nLength%d=%d\n",
		entry.Index, entry.Path, entry.Index, entry.Name, entry.Index, entry.Length)
}

type M3uFile struct {
	fileHead  string
	entryHead string
	entries   []M3uEntry
}

func readByBufioApi() {
	inArgs := os.Args
	if len(inArgs) < 2 || !strings.HasSuffix(inArgs[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(inArgs[0]))
		os.Exit(1)
	}
	m3uFileName := inArgs[1]
	m3uFile, err := os.OpenFile(m3uFileName, os.O_RDWR, 666)
	//apiDemo()
	if err != nil {
		fmt.Println("fail to open file ...")
		os.Exit(1)
	}
	reader := bufio.NewReader(m3uFile)
	n, err := reader.ReadString('\n')
	fmt.Println("n = ", n, ", err = ", err)
}

func main() {
	inArgs := os.Args
	if len(inArgs) < 2 || !strings.HasSuffix(inArgs[1], ".m3u") {
		fmt.Printf("usage: %s <file.m3u>\n", filepath.Base(inArgs[0]))
		os.Exit(1)
	}
	m3uFileName := inArgs[1]
	if rawBytes, err := ioutil.ReadFile(m3uFileName); err != nil {
		log.Fatal(err)
	} else {
		var entries []*M3uEntry
		//isEntryHead := false
		var entry M3uEntry
		count := 0
		for _, line := range strings.Split(string(rawBytes), "\n") {
			if line == "" || strings.HasPrefix(line, "#EXTM3U") {
				continue
			}
			if strings.HasPrefix(line, "#EXTINF") {
				//isEntryHead = true

				line = strings.TrimLeft(line, "#EXTINF:")
				idx := strings.Index(line, ",")
				length, err := strconv.ParseUint(string(line[0:idx]), 10, 32)
				if err != nil {
					continue
				}
				count++
				name := line[idx+1:]
				entry = M3uEntry{Index: uint32(count), Name: name, Length: uint32(length)}
				entries = append(entries, &entry)
			} else {
				//isEntryHead = false
				entry.Path = line
			}
		}
		fmt.Println(entries)
	}

}
