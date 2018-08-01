package main

import (
	"testing"
	"time"
	"sync"
	"fmt"
)

type S struct {
}

func (s *S) Close() {
	println("hello , close")
}

func Test_Defer(t *testing.T) {
	var s = &S{}

	var x  = make(chan int)
	go func(){
		time.Sleep(time.Duration(5)*time.Second)
		x <- 1
	}()
	defer s.Close()
	<-x

}


func Test_waitgroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i = i + 1 {
		wg.Add(1)
		go func(n int) {
			// defer wg.Done()
			defer wg.Done()
			EchoNumber(n)
		}(i)
	}

	wg.Wait()
}

func EchoNumber(i int) {
	time.Sleep(3e9)
	fmt.Println(i)
}