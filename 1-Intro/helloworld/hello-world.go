package main
// This line is called a package declaration and it tells the Go compiler to which package this ‘.go’ file belongs. 
// If this package declaration is ‘package main’, then this program will be compiled into an executable

import "fmt"
// The import keyword allows us to use code from other packages, 
// in this case the Println function from the fmt package. 
// Note how the package name is enclosed with double quotes ".

func main() {
	fmt.Println("Hello World")
}

// The func keyword denotes the start of a function declaration.
// func is followed by the name of the function: main.
// After the name there follows a pair of parentheses () and a set of curly braces {}.
// The function code is written inside the set of curly braces.

// Normally when we write functions, you need to write code to invoke them, 
// otherwise they are unused. However, the main function is different if it resides in the main package.

// When we have a main function in the main package, this is special to Go. 
// When compiled, an executable is produced, and when run, the executable uses the main function as the starting point

// To compile:
// go build hello-world.go
// ./hello-world

// To run without compiling 
// go run hello-world.go