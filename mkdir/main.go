package main

import (
	"flag"
	"fmt"
	"github.com/hrittikhere/golang/mkdir/cmd"
)

func main() {
	var directory string
	flag.StringVar(&directory, "n", "", " File Name")
	flag.Parse()

	CreatedDirectory, _ := cmd.ExecuteMkdir(directory)

	fmt.Printf("%s Created \n", CreatedDirectory)

}
