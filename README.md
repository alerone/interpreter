# Interpreter in Go
Making an interpreter with Go, following the book "Writing an interpreter in Go" by Thorsten Ball

## What is an interpreter?
Programming languages are made from Compilers and interpreters, the interpreter takes source code and evaluates it producing some visible, intermediate result that can be later executed. The compilers produce an output in another language that the underlying system can understand. 

## Structure of the interpreter

The intepreter is a mix of a lexer, a parser and a evaluator.

### Lexer

The lexer takes the source code and tokenizes it, so the language knows if the current token is a '=', a identifier, a digit,... it only "marks" the text for later being parsed.

### Parser

This parser is a tree-walking parser. The parser creates a Abstract Syntax Tree (AST). It's called abstract because some details visible in the source code are omitted in the AST (semicolons, spaces, newlines,...). This parser is a `top down operator precedence` parser, sometimes called "Pratt parser", after its inventor Vaughan Pratt. This parser makes learning about making an interpreter something fun! 🥳🎉

### Evaluator

What makes a programming language come to life from source code is a good Evaluator. When the interpreter walks by the parsing tree it evaluates every node thanks to the evaluator so an integer returns an Object.Integer, an expression, a statement is evaluated and so on. The evaluator makes the source code execute and return objects from the object system of the programming language.

## Functionality of the interpreter

This interpreter is a simple programming language made with Golang. It has the next functionality:
- `Integers` (5, 6, 987)
- `Booleans` (true or false)
- `Strings` ("Hello World")
- `Floats` (5.65, -34.55)
- `Function literals` (fn (x, y) {x + y})
- `Variables` (let x = 5, let add = fn(x, y) {x + y})
- `Arrays` (\[34, 44, 50\] or \[fn(x, y){x/y}, 345, true\]) --> see how the types of the array doesn't really matter.
- `Hash` ({"hello": "world", "foo": "bar"} or {true: false, 1: 35, "goo": true}) --> see how the types in key or in value doesn't really matter.

The language comes with some [`Builtin functions`](./evaluator/builtins.go):
- len(\<string or array\>): returns the length of the object
- first(\<array\>): returns the first element of the array (same with \<array\>\[0\])
- last(\<array\>): returns the last element of the arrray ((same with \<array\>\[len(\<array\>) - 1\])
- rest(\<array\>): returns the elemets a copy of the array without the first element
- push(\<array\>, elem): returns a copy of the array with the element as second argument in the last position of the array
- puts(object): It prints the object to the stdout.

 With those builtin functions you can build something interesting, like a map function (applies a function to all elements of an array)

 ```lisp
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
```

You can notice that this Monkey programming Language doesn't have a loop, all is Recursive!

## How to use the interpreter?

You can download the source code of the Repo and with a console in the root of the project with go installed type:
```cmd
go run ./main.go
```
![repl](https://github.com/user-attachments/assets/c7407d56-722f-4308-82aa-d66fb51821d3)


This is going to start the REPL (Read-Evaluate-Print-Loop) so you can write lines in the REPL and see how they are evaluated with the Monkey programming language!.

