package main

import (
	"fmt"
	"github.com/Devashish08/InterPreter-Compiler/evaluator"
	"github.com/Devashish08/InterPreter-Compiler/lexer"
	"github.com/Devashish08/InterPreter-Compiler/object"
	"github.com/Devashish08/InterPreter-Compiler/parser"
	"github.com/Devashish08/InterPreter-Compiler/repl"
	"io/ioutil"
	"os"
	"os/user"
)

func main() {
	if len(os.Args) < 2 {
		startRepl()
		return
	}

	command := os.Args[1]
	switch command {
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a file to execute")
			printHelp()
			os.Exit(1)
		}
		runFile(os.Args[2])
	case "repl":
		startRepl()
	case "help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func startRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runFile(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	env := object.NewEnvironment()
	l := lexer.New(string(input))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		os.Exit(1)
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  interpreter [command] [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  run <file>     Execute a source file")
	fmt.Println("  repl           Start the REPL (interactive mode)")
	fmt.Println("  help           Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  interpreter run examples/fibonacci.monkey")
	fmt.Println("  interpreter repl")
}

func printParserErrors(errors []string) {
	fmt.Printf("Parser errors:\n")
	for _, msg := range errors {
		fmt.Printf("\t%s\n", msg)
	}
}
