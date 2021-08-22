package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	c := B()
	fmt.Println(c == nil)
	fmt.Println(reflect.ValueOf(c).IsNil())

	w := make(map[int]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer func() {
				wg.Done()
			}()
			w[idx] = idx
		}(i)
	}
	wg.Wait()

	fmt.Println(w)
}

type T struct {
}

func B() interface{} {
	return A()
}

func A() map[string]string {
	return nil
}
