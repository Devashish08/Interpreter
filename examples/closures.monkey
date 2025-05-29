let newAdder = fn(x) {
    fn(y) { x + y };
};

let addTwo = newAdder(2);
puts(addTwo(5));
