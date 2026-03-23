package pointers

import "testing"

func TestWallet(t *testing.T) {

	checkBalance := func (t testing.TB, wallet Wallet, want Bitcoin)  {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Got %s Want %s", got, want)
		}
	}

	checkError := func (t testing.TB, got, want error)  {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		
	}
	checkNoError := func (t testing.TB, got error)  {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
		
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		checkBalance(t, wallet, Bitcoin(10))
		
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		checkNoError(t, err)
		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficient funds withdrawal", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkError(t, err, ErrInsufficientFunds)
		checkBalance(t, wallet, startingBalance)
	})
}