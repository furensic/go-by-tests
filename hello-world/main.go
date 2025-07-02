package main

import "fmt"

func ReturnHello() string {
	return "Hello, World!"
}

const greeterStringEnglish = "Hello, "
const greeterStringSpanish = "Hola, "
const greeterStringFrench = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func Greet(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greeterStringEnglish

	switch language {
	case spanish:
		prefix = greeterStringSpanish
	case french:
		prefix = greeterStringFrench
	default:
		prefix = greeterStringEnglish
	}

	return fmt.Sprintf("%s%s!", prefix, name)
}

func main() {
	fmt.Println(ReturnHello())
}
