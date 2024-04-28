## Brainf\*ck Interpreter in Go

Brainfuck is an esoteric programming language created in 1993 by Urban Müller. Notable for its extreme minimalism, the language consists of only eight simple commands, a data pointer and an instruction pointer.

Read more about the language on the [wiki page](https://en.wikipedia.org/wiki/Brainfuck).

This implementation follows this unofficial language specification : https://github.com/sunjay/brainfuck/blob/master/brainfuck.md

### Usage

Run the executable and pass a brainf\*ck file as command line argument to run it.

```sh
bfigo - Brainf*ck Interpreter in Go

usage:
    bfigo [filepath]
```

### Building from source

The project uses go's cli as it's build system, just clone the repo and run

```sh
go run .
```
