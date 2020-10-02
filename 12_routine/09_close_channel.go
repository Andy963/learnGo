package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if num, ok := <-ch; ok == true {
			fmt.Println("num =", num)
		} else {
			break
		}
	}
}

//if no data to send. you can close channel, otherwise no need.
// after close, you can not send data to channel
// but you can receive data after close if it have
// if nil channel no matter send or receive will blocked.