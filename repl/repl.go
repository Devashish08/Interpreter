// Package repl provides the interactive shell/REPL functionality

package repl

import (
	"bufio"
	"fmt"
	"io"
	"lang/evaluator"
	"lang/lexer"
	"lang/parser"
)

// PROMPT defines the interactive prompt symbol
const PROMPT = ">> "

// Start initializes and runs the REPL with the following workflow:
// 1. Shows prompt (>>)
// 2. Reads user input line
// 3. Creates lexer for the input
// 4. Tokenizes and prints each token
// 5. Repeats until EOF/exit
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
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
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}

}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

// Key components:
// 1. Uses bufio.Scanner for reading input
// 2. Creates new lexer instance for each line
// 3. Tokenizes input using lexer.NextToken()
// 4. Prints each token until EOF
// 5. Loops continuously until scanner fails/EOF

// Usage:
// repl.Start(os.Stdin, os.Stdout)  // For interactive shell
