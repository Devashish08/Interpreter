# 🎭 Monkey Interpreter Demo Script

## 🚀 Quick Demo (2-3 minutes)

### **Setup (30 seconds)**
```bash
# Build the interpreter
make build

# Show help to demonstrate CLI options
./interpreter help
```

### **1. Interactive REPL Demo (60 seconds)**
```bash
./interpreter repl
```

**In REPL, type these in sequence:**
```monkey
// Basic arithmetic
5 + 3 * 2

// Variables
let name = "Monkey"
let age = 1
name + " is " + string(age) + " year old"

// Functions
let greet = fn(x) { "Hello, " + x + "!" }
greet("World")

// Exit REPL
exit
```

### **2. File Execution Demo (60 seconds)**
```bash
# Run the fibonacci example
./interpreter run examples/fibonacci.monkey

# Run closures example
./interpreter run examples/closures.monkey
```

---

## 🎯 Comprehensive Demo (5-10 minutes)

### **1. Introduction (30 seconds)**
"This is a complete interpreter for the Monkey programming language, built in Go following Thorsten Ball's book. It demonstrates fundamental interpreter concepts like lexing, parsing, and tree-walking evaluation."

### **2. Language Features Showcase (3-4 minutes)**

#### **A. Basic Types and Operations**
```monkey
// Numbers and arithmetic
42
3.14159  // if you support floats
5 + 3 * 2
10 / 2 - 1

// Booleans and comparisons
true
false
5 > 3
10 == 10
5 != 10

// Strings
"Hello, World!"
"Hello" + " " + "World"
len("Monkey")
```

#### **B. Data Structures**
```monkey
// Arrays
let arr = [1, 2, 3, 4, 5]
arr[0]
len(arr)
first(arr)
last(arr)
rest(arr)

// Hash maps
let person = {"name": "Alice", "age": 30}
person["name"]
person["age"]
```

#### **C. Functions and Closures**
```monkey
// Simple function
let add = fn(x, y) { x + y }
add(5, 3)

// Higher-order functions
let twice = fn(f, x) { f(f(x)) }
let addOne = fn(x) { x + 1 }
twice(addOne, 5)

// Closures
let counter = fn() {
    let count = 0;
    fn() {
        count = count + 1;
        count
    }
}
let myCounter = counter()
myCounter()
myCounter()
myCounter()
```

### **3. File Examples (2-3 minutes)**
```bash
# Show advanced examples
./interpreter run examples/map-reduce.monkey
./interpreter run examples/fibonacci.monkey
```

---

## 🎪 Demo Tips

### **Presentation Flow:**
1. **Hook** - "I built a programming language interpreter"
2. **Context** - "Following a famous book on interpreter design"  
3. **Live Demo** - Show it working
4. **Architecture** - Brief explanation of components
5. **Q&A** - Answer questions about implementation

### **Common Questions to Prepare For:**
- "How long did this take to build?"
- "What was the hardest part?"
- "How does parsing work?"
- "Could you add feature X?"
- "How fast is it compared to other languages?"

### **Demo Dos and Don'ts:**

**✅ DO:**
- Test your demo beforehand
- Have backup examples ready
- Explain what you're typing as you type
- Keep examples simple and clear
- Show error handling too

**❌ DON'T:**
- Type too fast
- Use overly complex examples first
- Forget to explain the language syntax
- Skip showing the build process
- Assume audience knows interpreter concepts

---

## 🎬 Recording a Demo

If you want to create a GIF or video demo:

```bash
# Use asciinema to record terminal sessions
asciinema rec monkey-demo.cast

# Or use screen recording software for a full demo
```

**GIF Demo Script (30 seconds):**
1. `./interpreter repl` (2s)
2. `let fib = fn(n) { if(n < 2) { n } else { fib(n-1) + fib(n-2) } }` (8s)
3. `fib(10)` (3s)
4. `exit` (2s)
5. `./interpreter run examples/fibonacci.monkey` (5s)
6. Show output (5s)

---

## 🎯 Audience-Specific Demos

### **For Developers:**
- Focus on architecture and implementation details
- Show testing and code organization
- Discuss design decisions and trade-offs

### **For Students:**
- Emphasize learning aspects
- Show how book concepts translate to code
- Demonstrate fundamental CS concepts

### **For Recruiters/Managers:**
- Focus on project scope and completion
- Highlight problem-solving skills
- Show clean code and documentation 