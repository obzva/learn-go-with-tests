package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in Korean", func(t *testing.T) {
		got := Hello("Dongyeong", "Korean")
		want := "안녕, Dongyeong"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in French", func(t *testing.T) {
		got := Hello("Dongyeong", "French")
		want := "Bonjour, Dongyeong"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
