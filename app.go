package main

import "fmt"

// Constant
const pi = 3.14

// Declaration list
const (
	Pi = 3.14
	E  = 2.718
)

// main
func main() {
	fmt.Println("Hello app", pi)
}

// function
func greet(name string) string {
	return "hello" + name + "!"
}
