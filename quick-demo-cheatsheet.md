# 🎯 Quick Demo Cheatsheet

## ⚡ 30-Second Demo
```bash
cd Interpreter
./interpreter repl
```

**Type in REPL:**
```monkey
let fib = fn(n) { if(n < 2) { n } else { fib(n-1) + fib(n-2) } }
fib(10)
exit
```

## 🚀 2-Minute Demo Options

### Option A: Automated Demo
```bash
./run-demo.sh
```

### Option B: Manual Showcase
```bash
./interpreter run examples/showcase.monkey
```

### Option C: Interactive Features
```bash
./interpreter repl
```
**Then type:**
```monkey
// Arrays and functions
let numbers = [1, 2, 3, 4, 5]
let double = fn(x) { x * 2 }
let map = fn(arr, f) { /* implementation */ }

// Closures
let counter = fn() { let x = 0; fn() { x = x + 1; x } }
let c = counter()
c()
c()
```

## 🎤 Demo Sound Bites

**Opening:** "I built a programming language interpreter in Go, following the famous 'Writing an Interpreter in Go' book."

**Features:** "It supports functions, closures, arrays, hash maps, and has both a REPL and file execution."

**Technical:** "It uses a tree-walking interpreter with recursive descent parsing."

**Learning:** "This project taught me fundamental CS concepts like lexical analysis, parsing, and evaluation."

## 🔧 Demo Commands Ready to Copy-Paste

### Build and Help
```bash
make build
./interpreter help
```

### Quick Examples
```bash
./interpreter run examples/fibonacci.monkey
./interpreter run examples/closures.monkey
./interpreter run examples/showcase.monkey
```

### REPL One-Liners
```monkey
5 + 3 * 2
"Hello" + " " + "World"
[1, 2, 3][1]
{"name": "Alice"}["name"]
len("Monkey")
```

### Show Error Handling
```monkey
let x = 5
x()  // Show function call error
arr[100]  // Show index error
```

## 📱 Emergency Demo (No Terminal Access)

If you can't run the code, show:
1. **README.md** - Well-documented project
2. **Code structure** - Clean, modular design
3. **Test coverage** - `coverage.out` shows thorough testing
4. **Examples folder** - Working programs in Monkey language

## 🎯 Audience-Specific Hooks

**Developers:** "Want to see how lexers and parsers work? I built one from scratch."

**Students:** "This is what you learn in a compilers course, but actually implemented."

**Recruiters:** "3000+ lines of well-tested Go code implementing fundamental CS concepts."

**Managers:** "Complete project with documentation, testing, and clean architecture." 