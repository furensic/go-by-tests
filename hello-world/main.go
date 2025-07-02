package main

import "fmt"

func ReturnHello() string {
	return "Hello, World!"
}

const greeterString = "Hello, "

func Greet(name string) string {
	if name == "" || len(name) == 0 {
		return "Hello, World!"
	}

	return fmt.Sprintf("%s%s!", greeterString, name)
}

func main() {
	fmt.Println(ReturnHello())
}
