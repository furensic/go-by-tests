package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := ReturnHello()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGreeting(t *testing.T) {
	name := "Bill"
	got := Greet(name)
	want := fmt.Sprintf("Hello, %s!", name)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
