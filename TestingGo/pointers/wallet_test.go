package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalanace := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalanace(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		_ = wallet.Withdraw(Bitcoin(10))

		assertBalanace(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(100))

		// balance remains the same when withdrawing more than
		// funds present
		assertBalanace(t, wallet, Bitcoin(20))

		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	})
}
