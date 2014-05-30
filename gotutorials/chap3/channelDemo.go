package chap3

import (
	"fmt"
)

type myMapPair struct {
	key   string
	value interface{}
}

type mapChan struct {
	push_req chan *myMapPair
	push_rep chan interface{}
	pop_req  chan string
	pop_rep  chan interface{}
}

func (c *mapChan) push(key string, e interface{}) interface{} {
	c.push_req <- &myMapPair{key, e}
	return <-c.push_rep
}

func (c *mapChan) pop(key string) interface{} {
	c.pop_req <- key
	return <-c.pop_rep
}

func newMyMap() IMyMap {
	c := mapChan{
		push_req: make(chan *myMapPair),
		push_rep: make(chan interface{}),
		pop_req:  make(chan string),
		pop_rep:  make(chan interface{}),
	}
	m := make(map[string]interface{})
	go func() {
		for {
			select {
			case r := <-c.push_req:
				if v, exist := m[r.key]; exist {
					c.push_rep <- v
				} else {
					m[r.key] = r.value
					c.push_rep <- nil
				}
			case r := <-c.pop_req:
				if v, exist := m[r]; exist {
					m[r] = nil
					c.pop_rep <- v
				} else {
					c.pop_rep <- nil
				}
			}
		}
	}()
	return &c
}

func RunChannelDemo() {
	m := newMyMap()
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.pop("hello"))
	fmt.Println(m.pop("hello"))
}
