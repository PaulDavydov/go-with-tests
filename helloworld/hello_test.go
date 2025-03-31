package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello in english to people when language is blank", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say Hello World when empty name and language is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in spanish when spanish is specified", func(t *testing.T) {
		got := Hello("eddie", "Spanish")
		want := "Hola eddie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
