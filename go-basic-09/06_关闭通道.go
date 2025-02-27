package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(ok)
			fmt.Println(data)
		} else {
			fmt.Println(ok)
			fmt.Println(data)
			break
		}
	}
	fmt.Println("main结束")
}
