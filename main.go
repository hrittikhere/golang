package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Test Done")
	for i := 0; i < 100; i++ {
		fmt.Println(os.Getenv("NEW_MONEY"))
	}
}
