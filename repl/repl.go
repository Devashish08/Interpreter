// Package repl provides the interactive shell/REPL functionality

package repl

import (
	"bufio"
	"fmt"
	"io"
	"lang/lexer"
	"lang/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
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
