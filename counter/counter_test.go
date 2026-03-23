package counter

import "testing"

func TestIncrement(t *testing.T){
	
	t.Run("Increment", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		want := 1

		checkValue(t, val.value, want)
	})
	t.Run("Inreenting By Three", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		val.Increment()
		val.Increment()
		want := 3

		checkValue(t, val.value, want)
	})
}

func TestDecrement(t *testing.T) {
	t.Run("Decrementing a number above zero", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		err := val.Decrement()

		checkValue(t, val.value, 0)
		checkNoError(t, err)
	})
	t.Run("Decrementing Zero Results to Error", func(t *testing.T) {
		val := Counter{}
		err :=val.Decrement()

		checkError(t, err, ErrCounterAtZero)
	})
}

func TestValue(t *testing.T) {
	t.Run("Fresh conter", func(t *testing.T) {
		val := Counter{}
		result :=val.Value()

		checkValue(t,result, 0)
	})
	t.Run("After Operations", func(t *testing.T) {
		val := Counter{}
		val.Increment()
		val.Increment()
		val.Increment()
		result := val.Value()

		checkValue(t, result, 3)
	})
}

func checkValue (t testing.TB, value, want int )  {
	t.Helper()
	if value!= want{
		t.Errorf("Wanted %d, got %d",want, value)
	}
}

func checkError (t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("expected %q but got nil", want)
	}
	if got != want {
		t.Fatalf("expected %q but got %v", want, got)
	}
}

func checkNoError (t testing.TB, got error){
	t.Helper()
	if got != nil {
		t.Fatalf("did not expect error but got %v", got)
	}
}