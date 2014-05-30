package chap3

import (
	"fmt"
	"math/rand"
)

var keyGen func() string

func init() {
	keys := make(chan string)
	go func() {
		cnt := 0
		for {
			var buf [8]byte
			for i := 0; i < 8; i++ {
				buf[i] = byte(rand.Intn(26)) + byte('A')
			}
			cnt++
			fmt.Printf("before the %d time.\n", cnt)
			keys <- string(buf[:])
			fmt.Printf("after the %d time.\n", cnt)
		}
	}()
	keyGen = func() string {
		return <-keys
	}
}

func RunSection3() {
	fmt.Println(keyGen())
	fmt.Println(keyGen())
	fmt.Println(keyGen())
}
