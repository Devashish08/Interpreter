// 🎭 Monkey Language Showcase
// This file demonstrates the key features of the Monkey programming language

puts("=== Welcome to Monkey Language Showcase ===")
puts("")

// 1. Basic Data Types and Operations
puts("1. Basic Types and Arithmetic:")
puts("5 + 3 * 2 =", 5 + 3 * 2)
puts("10 > 5 =", 10 > 5)
puts("\"Hello\" + \" \" + \"World\" =", "Hello" + " " + "World")
puts("")

// 2. Variables and Scope
puts("2. Variables:")
let name = "Monkey"
let version = 1.0
puts("Language:", name, "Version:", version)
puts("")

// 3. Arrays and Built-in Functions
puts("3. Arrays and Built-ins:")
let numbers = [1, 2, 3, 4, 5]
puts("Array:", numbers)
puts("Length:", len(numbers))
puts("First:", first(numbers))
puts("Last:", last(numbers))
puts("Rest:", rest(numbers))
puts("")

// 4. Hash Maps
puts("4. Hash Maps:")
let person = {
    "name": "Alice",
    "age": 30,
    "city": "Wonderland"
}
puts("Person:", person)
puts("Name:", person["name"])
puts("Age:", person["age"])
puts("")

// 5. Functions
puts("5. Functions:")
let multiply = fn(x, y) {
    return x * y;
}
puts("multiply(6, 7) =", multiply(6, 7))

// 6. Higher-Order Functions
puts("6. Higher-Order Functions:")
let applyTwice = fn(func, value) {
    return func(func(value));
}

let addOne = fn(x) { x + 1 }
let result = applyTwice(addOne, 5)
puts("Apply addOne twice to 5:", result)
puts("")

// 7. Closures (Advanced Feature)
puts("7. Closures:")
let makeCounter = fn(start) {
    let count = start;
    fn() {
        count = count + 1;
        count
    }
}

let counter = makeCounter(10)
puts("Counter starting at 10:")
puts("Call 1:", counter())
puts("Call 2:", counter())
puts("Call 3:", counter())
puts("")

// 8. Recursive Functions
puts("8. Recursion - Factorial:")
let factorial = fn(n) {
    if (n <= 1) {
        return 1;
    } else {
        return n * factorial(n - 1);
    }
}

puts("5! =", factorial(5))
puts("7! =", factorial(7))
puts("")

// 9. Array Processing
puts("9. Array Processing:")
let reduce = fn(arr, initial, func) {
    let iter = fn(arr, result) {
        if (len(arr) == 0) {
            result
        } else {
            iter(rest(arr), func(result, first(arr)));
        }
    };
    iter(arr, initial);
}

let sum = reduce([1, 2, 3, 4, 5], 0, fn(acc, x) { acc + x })
puts("Sum of [1,2,3,4,5]:", sum)

let product = reduce([1, 2, 3, 4], 1, fn(acc, x) { acc * x })
puts("Product of [1,2,3,4]:", product)
puts("")

// 10. Conditional Logic
puts("10. Conditional Logic:")
let checkNumber = fn(n) {
    if (n > 0) {
        return "positive";
    } else {
        if (n < 0) {
            return "negative";
        } else {
            return "zero";
        }
    }
}

puts("checkNumber(42):", checkNumber(42))
puts("checkNumber(-10):", checkNumber(-10))
puts("checkNumber(0):", checkNumber(0))
puts("")

puts("=== Showcase Complete! ===")
puts("This demonstrates:")
puts("✓ Variables and data types")
puts("✓ Functions and closures") 
puts("✓ Arrays and hash maps")
puts("✓ Higher-order functions")
puts("✓ Recursion")
puts("✓ Built-in functions")
puts("✓ Conditional logic") 