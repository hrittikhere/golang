package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var directory string
	flag.StringVar(&directory, "n", "", " File Name")
	flag.Parse()

	os.Mkdir(directory, 0750)

	fmt.Printf("%s Created \n", directory)

}
