package main

import (
	"fmt"
	"time"
)

func main()  {

	start := time.Now().Unix()

	fmt.Printf("hello go %v\n", start)
}
