package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := ReturnHello()
	want := "Hello, World!"

	assertCorrectMessage(t, got, want)
}

func TestGreeting(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		name := "Bill"
		got := Greet(name, "")
		want := fmt.Sprintf("Hello, %s!", name)

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		name := ""
		got := Greet(name, "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say greeting in spanish", func(t *testing.T) {
		name := "Diego"
		got := Greet(name, "Spanish")
		want := fmt.Sprintf("Hola, %s!", name)
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
