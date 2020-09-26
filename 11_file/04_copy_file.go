package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args // get args from cmd
	if len(args) != 3 {
		fmt.Println("usage: xxx sourceFile, destFile")
		return
	}
	// check source dest file name
	source := args[1]
	dest := args[2]
	if source == dest {
		fmt.Println("source file name couldn't same as dest file name")
		return
	}
	// open source
	sf, serr := os.Open(source)
	if serr != nil {
		fmt.Println(serr)
		return
	}

	// create New dest file
	df, derr := os.Create(dest)
	if derr != nil {
		fmt.Println(derr)
		return
	}
	// close file
	defer sf.Close()
	defer df.Close()

	// read write
	buf := make([]byte, 1024*4)
	for {
		n, err := sf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err=", err)
		}
		// write
		df.Write(buf[:n])
	}

}
