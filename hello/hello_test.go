package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Nicolas", "")
		want := "Hello, Nicolas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Nicolas", "Portuguese")
		want := "Olá, Nicolas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Nicolas", "French")
		want := "Bonjour, Nicolas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func(t *testing.T) {
		got := Hello("Nicolas", "Russian")
		want := "Privet, Nicolas"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
