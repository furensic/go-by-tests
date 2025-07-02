package main

import "fmt"

func ReturnHello() string {
	return "Hello, World!"
}

const (
	greeterStringEnglish = "Hello, "
	greeterStringSpanish = "Hola, "
	greeterStringFrench  = "Bonjour, "
	spanish              = "Spanish"
	french               = "French"
)

func greetingPrefix(language string) (prefix string) {
	prefix = greeterStringEnglish

	switch language {
	case spanish:
		prefix = greeterStringSpanish
	case french:
		prefix = greeterStringFrench
	default:
		prefix = greeterStringEnglish
	}

	return prefix
}

func Greet(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s%s!", greetingPrefix(language), name)
}

func main() {
	fmt.Println(ReturnHello())
}
