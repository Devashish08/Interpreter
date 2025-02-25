# Go Interpreter

A powerful interpreter implemented in Go, based on Thorsten Ball's "Writing an Interpreter in Go" with extended features and improvements.

## Features

- Lexical Analysis and Tokenization
- Recursive Descent Parser
- Abstract Syntax Tree (AST) Implementation
- REPL (Read-Eval-Print Loop) Interface
- Support for:
  - Integer and Boolean data types
  - String data types
  - Array data structures
  - Hash data structures
  - First-class functions
  - Built-in functions
  - Prefix and Infix operators

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Devashish08/InterPreter-Compiler.git
cd InterPreter-Compiler
```

2. Build the interpreter:
```bash
go build
```

### Usage

#### REPL Mode
```bash
./interpreter repl
```

#### Run a file
```bash
./interpreter run examples/fibonacci.monkey
```

## Project Structure

```
.
├── ast/          # Abstract Syntax Tree implementation
├── evaluator/    # Expression evaluation logic
├── lexer/       # Lexical analysis
├── parser/      # Parsing logic
├── object/      # Object system implementation
├── repl/        # REPL implementation
├── token/       # Token definitions
└── examples/    # Example programs
```

## Examples

Here's a simple program that demonstrates the language features:

```monkey
let fibonacci = fn(x) {
  if (x < 2) {
    return x;
  }
  return fibonacci(x - 1) + fibonacci(x - 2);
};

fibonacci(10);
```

## Development

### Running Tests

```bash
make test
```

### Running with Coverage

```bash
make coverage
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Based on Thorsten Ball's "Writing an Interpreter in Go"
- Extended with additional features and improvements
