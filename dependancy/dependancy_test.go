package dependancy

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("normal greet func with name", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Mesh")

		got := buffer.String()
		want := "Hello, Mesh"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}