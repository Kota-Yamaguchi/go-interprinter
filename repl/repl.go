package repl

import (
	"bufio"
	"fmt"

	"io"
	"monkey/monkey/evaluator"
	"monkey/monkey/lexer"
	"monkey/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
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
		eval := evaluator.Eval(program)
		if eval != nil {
			io.WriteString(out, eval.Inspect())
		}

		io.WriteString(out, "\n")

	}

}
func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
