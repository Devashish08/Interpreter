#!/bin/bash

# 🎭 Monkey Interpreter Automated Demo Script
# This script runs a complete demonstration of the interpreter

echo "🐒 Welcome to the Monkey Language Interpreter Demo!"
echo "=================================================="
echo ""

# Check if interpreter is built
if [ ! -f "./interpreter" ]; then
    echo "Building the interpreter..."
    make build
    echo ""
fi

echo "📋 Demo Overview:"
echo "1. Show interpreter help and version"
echo "2. Run comprehensive showcase"
echo "3. Run individual examples"
echo "4. Interactive REPL demo"
echo ""

# Function to pause for effect
pause() {
    echo "Press Enter to continue..."
    read -r
}

echo "🚀 1. Interpreter Help and CLI Options"
echo "======================================"
./interpreter help
echo ""
pause

echo "🎯 2. Comprehensive Language Showcase"
echo "====================================="
echo "Running: ./interpreter run examples/showcase.monkey"
echo ""
./interpreter run examples/showcase.monkey
echo ""
pause

echo "📊 3. Individual Feature Examples"
echo "================================="

echo "🔢 Fibonacci Sequence:"
echo "Running: ./interpreter run examples/fibonacci.monkey"
./interpreter run examples/fibonacci.monkey
echo ""

echo "🔗 Closures Example:"
echo "Running: ./interpreter run examples/closures.monkey"
./interpreter run examples/closures.monkey
echo ""

echo "🔄 Map-Reduce Example:"
echo "Running: ./interpreter run examples/map-reduce.monkey"
./interpreter run examples/map-reduce.monkey
echo ""

echo "📝 String Manipulation:"
echo "Running: ./interpreter run examples/string_manipulation.monkey"
./interpreter run examples/string_manipulation.monkey
echo ""
pause

echo "💬 4. Interactive REPL Demo"
echo "==========================="
echo "Starting REPL with some example commands..."
echo ""
echo "Try typing these commands:"
echo "  let greet = fn(name) { \"Hello, \" + name + \"!\" }"
echo "  greet(\"World\")"
echo "  [1, 2, 3, 4, 5]"
echo "  len(\"Monkey\")"
echo "  exit"
echo ""
echo "Starting REPL..."
./interpreter repl

echo ""
echo "🎉 Demo Complete!"
echo "================="
echo ""
echo "What you just saw:"
echo "✓ Complete programming language interpreter"
echo "✓ Built following 'Writing an Interpreter in Go'"
echo "✓ Supports functions, closures, arrays, hash maps"
echo "✓ Interactive REPL and file execution"
echo "✓ Comprehensive error handling"
echo ""
echo "Architecture:"
echo "📝 Lexer → 🌳 Parser → 🧠 Evaluator"
echo ""
echo "Thank you for watching the demo! 🐒✨" 