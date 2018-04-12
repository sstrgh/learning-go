package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello world when name Chris is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Chris", ""), "Hello, Chris")
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, Hello("", ""), "Hello, World")
	})

	t.Run("reply in spanish when passed spanish", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Esmarelda", "Spanish"), "Hola!, Esmarelda")
	})

	t.Run("reply in french when passed french", func(t *testing.T) {
		assertCorrectMessage(t, Hello("Esmarelda", "French"), "Bonjour, Esmarelda")
	})
}
