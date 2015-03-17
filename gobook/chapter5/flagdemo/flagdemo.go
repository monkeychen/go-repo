// flagdemo.go
package main

import (
	"flag"
	"fmt"
	"os"
	path "path/filepath"
)

func main() {
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	var flagvar int
	flag.IntVar(&flagvar, "flagname2", 1234, "help message for flagname2")
	flag.Parse()
	fmt.Println(*ip)
	fmt.Println(flag.Args())
	fmt.Println(flag.Arg(0))
	fmt.Println(flagvar)
	fmt.Println(os.Args)
	fmt.Println(os.Args[0])
	fmt.Println(os.Getwd())
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(path.SplitList(os.Getenv("GOPATH")))
	fmt.Println(os.Stat("."))
	fmt.Printf("% x\n", uint32(os.ModeAppend))
}
