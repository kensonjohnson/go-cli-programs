package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		file, _ := os.ReadDir("./")
		for _, file := range file {
			fmt.Print(file.Name(), " ")
		}
	} else {
		file, err := os.ReadDir(os.Args[1])
		if err != nil {
			panic(err)
		}
		for _, file := range file {
			fmt.Print(file.Name(), " ")
		}
	}

	fmt.Println()
}
