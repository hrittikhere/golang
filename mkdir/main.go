package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var directory string
	flag.StringVar(&directory, "n", "", " File Name")
	flag.Parse()

	err := os.Mkdir(directory, 0750)

	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	fmt.Printf("%s Created \n", directory)

}
