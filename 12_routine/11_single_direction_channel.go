package main

func main() {
	ch := make(chan int)
	// double direction channel can transfer to single direction channel implicitly
	var writeCh chan<- int = ch // only write, no read
	var readCh <-chan int = ch  // only read, no write

	writeCh <- 666
	<-readCh
}
