package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fail(errors.New("String empty"))
	}
	str := os.Args[1]

	if str == "" {
		fmt.Println(0)
		return
	}
	fmt.Println(len(strings.Split(str, " ")))
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
