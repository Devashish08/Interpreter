
let str = "Hello, World!";
puts("Original string:", str);

let upperStr = upper(str);
puts("Uppercase:", upperStr);
let lowerStr = lower(str);
puts("Lowercase:", lowerStr);


let words = split(str, " ");
puts("Split words:");
puts(words);
let joined = join(words, "-");
puts("Joined with hyphen:", joined);

let greeting = "Hello" + " " + "there!";
puts("Concatenated:", greeting);
