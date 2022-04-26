package main

import (
	"fmt"
	"os"

	"github.com/hrittikhere/golang/ls/cmd"
)

func main() {
	wd, _ := os.Getwd()
	res, _ := cmd.ExecuteLs(wd)
	fmt.Printf("%s\n", res)

}
