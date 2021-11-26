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

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	count := 1
	for {
		fmt.Fprintf(out, ">>>"+strconv.Itoa(count)+" ")
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

		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
		}

		fmt.Println(comp.Bytecode())
		machine := vm.New(comp.Bytecode())
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s \n", err)
			continue
		}

		stackTop := machine.StackTop()
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
	Start(os.Stdin, os.Stdout)
}
