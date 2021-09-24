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
	fmt.Printf("Hello %s (Current Logged User )\nThis is LAAIT(lite) Token Synthesizer.\n",
		user.Username)
	fmt.Println("|_   _|__ | | _____ _ __  ___")
	fmt.Println("  | |/ _ \\| |/ / _ \\ '_ \\/ __|")
	fmt.Println("  | | (_) |   <  __/ | | \\__ \\")
	fmt.Println("  |_|\\___/|_|\\_\\___|_| |_|___/")

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)

	fmt.Println("\nThanks to use LAAIT")
}
