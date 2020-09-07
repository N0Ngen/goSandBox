package main

// Import testing package
import "testing"

// Test function
// "t *testing.T" is the HOCK to testing package
func TestHello(t *testing.T) {

	// Simple test function
	// got := Hello("Alix")
	// want := "Hello, Alix"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	// Cut the repetted parts
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// Subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Alix", "")
		want := "Hello, Alix"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Reo", "Spanish")
		want := "Hola, Reo"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sam", "French")
		want := "Bonjour, Sam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Arabic", func(t *testing.T) {
		got := Hello("Ali", "Arabic")
		want := "مرحبا, Ali"
		assertCorrectMessage(t, got, want)
	})

}
