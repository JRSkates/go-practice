package main

// To import multiple packages, use a single import with parentheses
// We can also provide an alias to a package by specifying an alias name before the package name.
import (
	"fmt"
	t "time"
)

func main() {
	fmt.Println(t.Now())
}