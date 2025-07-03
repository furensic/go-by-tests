package pointers_errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("got nil, want error")
		}

		if !errors.Is(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	assertErrorNil := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("got %q want nil", got)
		}
	}

	withdrawTests := []struct {
		name    string
		balance Bitcoin
		amount  Bitcoin
		want    Bitcoin
	}{
		{name: "deposit 5", balance: Bitcoin(10), amount: Bitcoin(5), want: Bitcoin(5)},
		{name: "deposit 0", balance: Bitcoin(0), amount: Bitcoin(0), want: Bitcoin(0)},
	}

	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	for _, testCase := range withdrawTests {
		t.Run(testCase.name, func(t *testing.T) {
			wallet := Wallet{testCase.balance}
			err := wallet.Withdraw(testCase.amount)

			assertErrorNil(t, err)
			assertBalance(t, wallet, testCase.want)
		})
	}

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrorInsufficientFunds)
	})
}
