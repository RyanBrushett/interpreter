package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ryanbrushett/interpreter/repl"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	user, err := user.Current()
	checkError(err)
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in some commands.")
	repl.Start(os.Stdin, os.Stdout)
}
