package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Test Done")

	ab := os.Getenv("NEW_MONEY")
	bc := ab+ab

	for i := 0; i < 100; i++ {
		fmt.Println(bc)
	}
}
