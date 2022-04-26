package cmd

import (
	"fmt"
	"os"
)

func ExecuteLs(path string) (string, error) {

	entries, _ := os.ReadDir(path)

	output := fmt.Sprintf("Files in %s\n", path)
	output += "Name\t\tDirectory\t\n"

	for _, e := range entries {
		output += fmt.Sprintf("%s\t\t%v\n", e.Name(), e.IsDir())
	}

	return output, nil
}
