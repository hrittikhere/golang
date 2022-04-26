package cmd

import (
	"os"
)

func ExecuteMkdir(directory string) (string, error) {

err :=os.Mkdir(directory, 0750)

return directory,err


}
