package main

import "testing"

func TestHello(t *testing.T) {
	got := ReturnHello()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
