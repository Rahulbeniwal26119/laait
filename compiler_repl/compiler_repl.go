package main

import (
	"bufio"
	"fmt"
	"io"
	"laait/compiler"
	"laait/lexer"
	"laait/object"
	"laait/parser"
	"laait/vm"
	"os"
	"strconv"
)

// main code initiatiator
func Start(in io.Reader, out io.Writer, show_bytecode bool) {
	scanner := bufio.NewScanner(in)
	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalSize)
	symbolTable := compiler.NewSymbolTable()

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

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(program)
		//
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
			continue
		}
		//
		if show_bytecode {
			io.WriteString(out, " ====== Bytecode ===== \n")
			fmt.Println(comp.Bytecode())
			io.WriteString(out, " ====== * ===== \n")
		}
		code := comp.Bytecode()
		constants = code.Constants
		machine := vm.NewWithGlobalStore(code, globals)
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
