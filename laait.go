package main

import (
	"fmt"
	"laait/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s This is LAAIT(lite) Progrmming Language.\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

	fmt.Println("\nThanks to use LAAIT")
}
