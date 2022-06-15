package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people in english", func(t *testing.T) {
		got := Hello("Jacob", "english")
		want := "Hello, Jacob!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in french", func(t *testing.T) {
		got := Hello("Camille", "french")
		want := "Bonjour, Camille!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("default response from Hello() func", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
}
