package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Test Done")

	env := (os.Getenv("NEW_MONEY"))

	new, err := strconv.Atoi(env)
	
	if err != nil {
        fmt.Println(err)
    }
	var count int = new

	for i := 0; i < count; i++ {

		fmt.Println("Running")
		
	}
}
