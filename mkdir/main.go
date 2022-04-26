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

	fmt.Printf("%s Created \n", directory)

}
