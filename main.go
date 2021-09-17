package main

import "runtime"

func main() {
	c := make(chan int)
	n := 0
	for {
		n++
		go func() {
			sum := 0
			for i := 0; i < 100000000; i++ {
				sum += i
			}
			println(sum)
			c <- 0
		}()
		if n > 9 {
			break
		}
	}

	println("NumGoroutine=", runtime.NumGoroutine())
	<-c
}
