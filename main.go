package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aaazlkm/go-interpreter/repl"
)

func main() {
	fmt.Println("hello world")
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello world %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
