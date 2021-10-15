package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"laait/environment"
	"laait/evaluator"
	"laait/lexer"
	"laait/parser"
)

func Start(in io.Reader, out io.Writer) {
	const PARSER_PROMPT = ">>> "

	scanner := bufio.NewScanner(in)
	env := environment.NewEnvironment()

	for {
		fmt.Print(PARSER_PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
