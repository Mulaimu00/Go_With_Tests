package counter2

import "testing"

func TestIncrement(t *testing.T) {
	t.Run("Incrementing to one", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		got := val.value
		want := 1

		checkValue(t, got, want)
	})
	t.Run("Incrementing to 3", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		val.Increment()
		val.Increment()
		got := val.value
		want := 3

		checkValue(t, got, want)
	})
}

func TestDecrement(t *testing.T) {
	t.Run("Decrementing Number above 0", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		err := val.Decrement()
		got := val.value

		checkNoError(t, err)
		checkValue(t, got, 0)
	})
	t.Run("Decreenting zero", func(t *testing.T) {
		val := Counter{}
		err := val.Decrement()
		checkError(t, err, ErrcounterAtZero)
	})
}

func TestReturnValue(t *testing.T) {
	t.Run("Returning a value", func(t *testing.T) {
		val := Counter{}
		got := val.Value()
		want := 0

		checkValue(t, got, want)
	})
}

func checkValue(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d, got%d", want, got)
	}
}

func checkNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got error did not want one")
	}
}

func checkError(t testing.TB, err, want error)  {
	t.Helper()
	if err == nil {
		t.Fatal("wanted error did not get one")
	}
	if err != want {
		t.Errorf("got %q, wanted %q", err, want)
	}
}
