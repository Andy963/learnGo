package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//step 1: set seed
	//step 2: gen num

	rand.Seed(666) // if the seed is const the rand num is same every time
	for i :=0;i <5;i++{
		// gen num
		fmt.Println("rand= ",rand.Int()) //big number
	}

	rand.Seed(time.Now().UnixNano()) // use sys time as seed
	for i :=0;i<5;i++{
		fmt.Println("rand= ",rand.Intn(100)) // the rand num is smaller than 100
	}
}

