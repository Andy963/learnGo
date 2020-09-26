package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLine(path string) {
	// open
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//close
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err=", err)
		}
		fmt.Printf("buf=#%s#", string(buf))
	}

}
func main() {
	path := "./demo.txt"
	ReadLine(path)
}
