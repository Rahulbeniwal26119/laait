package main

import (
	"fmt"
	"laait/repl"
	"laait/repl/evaluator"
	"laait/repl/laait_parser"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) < 2 {
		printUser()
		evalutor()
		fmt.Println("\nThanks to use LAAIT")
	} else if len(os.Args) == 3 {
		evaluator.Start_notebook("read_text.txt", "output_.txt")
	} else {
		printUser()
		cmd := os.Args[1]
		if cmd == "lexer" || cmd == "token" {
			lexer()
		} else if cmd == "parser" {
			parser()
		}
		fmt.Println("\nThanks to use LAAIT")
	}
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
	fmt.Println(" ____                                     ")
	fmt.Println("|  _ \\    __ _   _ __   ___    ___   _ __")
	fmt.Println("| |_) |  / _` | | '__| / __|  / _ \\ | '__|")
	fmt.Println("|  __/  | (_| | | |    \\__ \\ |  __/ | |")
	fmt.Println("|_|      \\__,_| |_|    |___/  \\___| |_|")
	laait_parser.Parser_Start(os.Stdin, os.Stdout)
}

func lexer() {
	fmt.Println(" _____           _")
	fmt.Println("|_   _|   ___   | | __   ___   _ __")
	fmt.Println("  | |    / _ \\  | |/ /  / _ \\ | '_ \\ ")
	fmt.Println("  | |   | (_) | |   <  |  __/ | | | |")
	fmt.Println("  |_|    \\___/  |_|\\_\\  \\___| |_| |_|")
	repl.Start(os.Stdin, os.Stdout)
}

func evalutor() {
	fmt.Println(" _          _         _      ___   _____")
	fmt.Println("| |        / \\       / \\    |_ _| |_   _|")
	fmt.Println("| |       / _ \\     / _ \\    | |    | |")
	fmt.Println("| |___   / ___ \\   / ___ \\   | |    | |")
	fmt.Println("|_____| /_/   \\_\\ /_/   \\_\\ |___|   |_|")
	evaluator.Start(os.Stdin, os.Stdout)
}
