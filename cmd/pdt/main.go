package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("pdt v0.1.0-dev")
		return
	}
	fmt.Println("PDT - The Personal Data Toolkit (scaffold ready)")
}
