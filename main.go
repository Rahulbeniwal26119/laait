package main

import (
	"fmt"
	"laait/repl"
	"laait/repl/laait_parser"
	"os"
	"os/user"
)

func main() {
	printUser()
	if len(os.Args) < 2 {
		lexer()
	} else {
		cmd := os.Args[1]
		if cmd == "lexer" {
			lexer()
		} else if cmd == "parser" {
			parser()
		}
	}
	fmt.Println("\nThanks to use LAAIT")
}

func printUser() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s (Current Logged User )\nWelcome to LAAIT.\n",
		user.Username)
}

func parser() {
	fmt.Println(" ____        ")
	fmt.Println("|  _ \\ __ _ _ __ ___  ___ _ __ ")
	fmt.Println("| |_) / _` | '__/ __|/ _ \\ '__|")
	fmt.Println("|  __/ (_| | |  \\__ \\  __/ |  ")
	fmt.Println("|_|   \\__,_|_|  |___/\\___|_|  ")
	laait_parser.Parser_Start(os.Stdin, os.Stdout)
}

func lexer() {
	fmt.Println(" _____     _      ")
	fmt.Println("|_   _|__ | | _____ _ __  ___ ")
	fmt.Println("  | |/ _ \\| |/ / _ \\ '_ \\/ __|")
	fmt.Println("  | | (_) |   <  __/ | | \\__ \\")
	fmt.Println("  |_|\\___/|_|\\_\\___|_| |_|___/")
	repl.Start(os.Stdin, os.Stdout)
}
