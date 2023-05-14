package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Jose")
		want := "Hello, Jose"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
