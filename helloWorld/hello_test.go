package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello English", func(t *testing.T) {
		got := Hello("Mesh", "")
		want := "Hello, Mesh"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello Spanish", func(t *testing.T) {
		got := Hello("Mesh", "Spanish")
		want := "Hola, Mesh"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello Swahili", func(t *testing.T) {
		got := Hello("Mesh", "Swahili")
		want := "Jambo, Mesh"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say \"Hello, world\" when it is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}