package main

import "fmt"

func ReturnHello() string {
	return "Hello, World!"
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(ReturnHello())
}
