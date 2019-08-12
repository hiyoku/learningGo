package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Hello, 世界 " + strconv.Itoa(i))
	}
}
