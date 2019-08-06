package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for i, v := range os.Args {
		fmt.Println("Arg:", v, "Index:", i)
	}
}

func joinArgsCustom(a []strings) string {
	s, sep := "", ""
	for _, arg := range a {
		s += sep + arg
		sep := " "
	}
	fmt.Println(s)
	return s
}

func joinArgs(a []strings) string {
	return strings.Join(a, " ")
}
