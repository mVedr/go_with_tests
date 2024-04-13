package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris") //imported from "main/hello.go"
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
