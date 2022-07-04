package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("File name must be informed!")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Something went wrong trying to open the specified file!", err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, file)
}
