package main

import(
	"sync"
	"fmt"
	"time"
)

func Merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
  
	out := make(chan int)
  
	send := func(c <-chan int) {
	  for n := range c {
		out <- n
	  }
	  wg.Done()
	}
  
	wg.Add(len(ch))
	for _, c := range ch {
	  go send(c)
	}
  
	go func() {
	  wg.Wait()
	  close(out)
	}()
	return out
  }

func makeChannel(a ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, i := range a {
			ch <- i
			time.Sleep(time.Duration(100) * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

func main() {
	a := makeChannel(1, 2, 3, 4, 5, 6, 7, 8)
	b := makeChannel(11, 12, 13, 14, 15)
	for i := range Merge(a, b) {
		fmt.Println(i)
	}
}