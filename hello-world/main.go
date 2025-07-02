package main

import "fmt"

func ReturnHello() string {
	return "Hello, World!"
}

const greeterString = "Hello, "
const greeterStringSpanish = "Hola, "

const spanish = "Spanish"

func Greet(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return fmt.Sprintf("%s%s!", greeterStringSpanish, name)
	default:
		return fmt.Sprintf("%s%s!", greeterString, name)
	}
}

func main() {
	fmt.Println(ReturnHello())
}
