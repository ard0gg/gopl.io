package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Println("Arg:", v, "Index:", i)
	}
}
