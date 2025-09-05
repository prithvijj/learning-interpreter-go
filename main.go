package main

import (
	"fmt"
	"os"
	"os/user"

	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! This is monkey language\n", user.Username)
	fmt.Printf("type any commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
