package chap3

import (
	"fmt"
	"math/rand"
	"sync"
)

var keyGen func() string

type IMyMap interface {
	push(key string, e interface{}) interface{}
	pop(key string) interface{}
}

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

type myMap struct {
	m map[string]interface{}
	sync.Mutex
}

func (m *myMap) push(key string, e interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	defer fmt.Println("push:exe m.Unlock.")
	if v, exist := m.m[key]; exist {
		return v
	}
	m.m[key] = e
	return nil
}

func (m *myMap) pop(key string) interface{} {
	m.Lock()
	defer m.Unlock()
	defer fmt.Println("pop:exe m.Unlock.")
	if v, exist := m.m[key]; exist {
		m.m[key] = nil
		return v
	}
	return nil
}

func newMap() *myMap {
	return &myMap{m: make(map[string]interface{})}
}

func RunSection3_1() {
	m := newMap()
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.pop("hello"))
	fmt.Println(m.pop("hello"))
}

func RunSection3() {
	fmt.Println(keyGen())
	fmt.Println(keyGen())
	fmt.Println(keyGen())
}
