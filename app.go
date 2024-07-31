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
	name := "App"
	fmt.Println(greet(name))
}

// function
func greet(name string) string {
	return "hello " + name + "!"
}
