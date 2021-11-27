package main

import (
	"bufio"
	"fmt"
	"io"
	"laait/compiler"
	"laait/lexer"
	"laait/parser"
	"laait/vm"
	"os"
	"strconv"
)

func Start(in io.Reader, out io.Writer, show_bytecode bool) {
	scanner := bufio.NewScanner(in)
	count := 1
	if show_bytecode {
		fmt.Println(" === Bytecode Verbose Mode === ")
	}
	for {
		fmt.Fprintf(out, ">>> "+"["+strconv.Itoa(count)+"] ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		comp := compiler.New()
		err := comp.Compile(program)
		//
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
		}
		//
		if show_bytecode {
			io.WriteString(out, " ====== Bytecode ===== \n")
			fmt.Println(comp.Bytecode())
			io.WriteString(out, " ====== * ===== \n")
		}
		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s \n", err)
			continue
		}
		//
		stackTop := machine.LastPoppedStackElem()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
		count++
	}

}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func main() {
	show_bytecode := false
	if len(os.Args) > 1 {
		if os.Args[1] == "-b" {
			show_bytecode = true
		}
	}
	Start(os.Stdin, os.Stdout, show_bytecode)
}
