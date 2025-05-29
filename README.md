# Monkey Language Interpreter

A fully-featured interpreter for the Monkey programming language, built in Go following Thorsten Ball's "Writing an Interpreter in Go" book.

## 🐒 About Monkey

Monkey is a programming language designed for learning interpreter and compiler construction. It supports:

- **Variables** and **integers**
- **Arithmetic expressions**
- **Built-in functions**
- **First-class functions** and **closures**
- **Strings** and **string manipulation**
- **Arrays** and **hash maps**
- **Control flow** (if/else statements)

## ✨ Features

- **Interactive REPL** (Read-Eval-Print Loop)
- **File execution** mode
- **Comprehensive error reporting**
- **Full test coverage** with Go testing
- **Clean, modular architecture**

## 🚀 Quick Start

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Devashish08/InterPreter-Compiler.git
cd InterPreter-Compiler
```

2. Build the interpreter:
```bash
make build
```

### Usage

#### Interactive Mode (REPL)
```bash
./interpreter repl
```

#### Execute a Monkey file
```bash
./interpreter run examples/fibonacci.monkey
```

#### Show help
```bash
./interpreter help
```

## 📝 Language Examples

### Variables and Functions
```monkey
let name = "Monkey";
let age = 1;
let inspirations = ["Scheme", "Lisp", "JavaScript", "Clojure"];

let greetPeople = fn(name, age) {
    return "Hello " + name + ", you are " + age + " years old!";
};

greetPeople("Anna", 24);
```

### Fibonacci Sequence
```monkey
let fibonacci = fn(x) {
    if (x < 2) {
        return x;
    }
    return fibonacci(x - 1) + fibonacci(x - 2);
};

fibonacci(10); // Returns: 55
```

### Higher-Order Functions
```monkey
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            accumulated
        } else {
            iter(rest(arr), push(accumulated, f(first(arr))));
        }
    };
    iter(arr, []);
};

let double = fn(x) { x * 2 };
map([1, 2, 3, 4], double); // Returns: [2, 4, 6, 8]
```

## 🏗️ Architecture

The interpreter is built with a clean, modular architecture:

```
├── lexer/          # Tokenizes source code
├── parser/         # Builds Abstract Syntax Tree (AST)
├── ast/            # AST node definitions
├── object/         # Object system and environment
├── evaluator/      # Tree-walking evaluator
├── repl/           # Interactive REPL implementation
├── token/          # Token definitions
└── examples/       # Sample Monkey programs
```

### Core Components

1. **Lexer** (`lexer/`): Converts source code into tokens
2. **Parser** (`parser/`): Builds an Abstract Syntax Tree (AST) from tokens
3. **AST** (`ast/`): Defines the structure of the syntax tree
4. **Evaluator** (`evaluator/`): Tree-walking interpreter that executes the AST
5. **Object System** (`object/`): Represents values and manages scope/environment
6. **REPL** (`repl/`): Interactive shell for testing and exploration

## 🧪 Development

### Running Tests
```bash
make test
```

### Generate Coverage Report
```bash
make coverage
```

### Clean Build Artifacts
```bash
make clean
```

## 📚 Built-in Functions

Monkey comes with several built-in functions:

- `len(object)` - Returns length of strings, arrays, or hash maps
- `first(array)` - Returns first element of an array
- `last(array)` - Returns last element of an array
- `rest(array)` - Returns new array with all elements except the first
- `push(array, element)` - Returns new array with element appended
- `puts(object)` - Prints object to stdout

## 🎯 Language Features

### Data Types
- **Integers**: `42`, `-5`
- **Booleans**: `true`, `false`
- **Strings**: `"Hello World"`
- **Arrays**: `[1, 2, 3]`
- **Hash Maps**: `{"key": "value", "number": 42}`
- **Functions**: `fn(x, y) { x + y }`

### Operators
- **Arithmetic**: `+`, `-`, `*`, `/`
- **Comparison**: `==`, `!=`, `<`, `>`
- **Logical**: `!` (bang operator)

### Control Flow
```monkey
if (condition) {
    // do something
} else {
    // do something else
}
```

## 📖 Learning Resources

This interpreter was built following:
- **"Writing an Interpreter in Go"** by Thorsten Ball
- [Official Book Website](https://interpreterbook.com/)

## 🤝 Contributing

This is a learning project, but contributions are welcome! Please feel free to:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

## 🙏 Acknowledgments

- **Thorsten Ball** for the excellent "Writing an Interpreter in Go" book
- The **Go community** for creating such a great language for systems programming
- **Robert Nystrom** for "Crafting Interpreters" which inspired the field

---

**Happy coding! 🐒✨** 