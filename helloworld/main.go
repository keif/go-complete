package main

// a package is like a project or workspace
// a package can have a lot of files associated with it
// package main would be shared across several files, so they'd all need package main
// why main?
// main is a special package name that tells Go that this is an executable program
// the other kind is a reusable package, which is a package that can be used as a dependency in other programs (reusable packages, code dependencies, helper functions, etc.)

// how do we know when we are making an executable package or a reusable package?
// the name of the package determines if we are making an executable or a reusable package
// package main -> go build -> main.exe

// compiling a non-main package gives nothing

// import gives our package access to some code written in another package
// fmt is part of the standard library - golang.org/pkg for more details on standard libs
import "fmt"

// pacakge main means this is executable
// so we need a func main to be executable
// func is short for function (similar to everything you're thinking of)
// func - declare new function
// name of the function (main)
// list of args () // not in this examples
// function body {}

func main() {
	fmt.Println("Hello, World!")
}
