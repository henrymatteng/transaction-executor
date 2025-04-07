package transaction

import (
	"fmt"

	"github.com/henrymatteng/transaction_executor/account"
)

type Transfer struct {
	From  string
	To    string
	Value int
}

func (t Transfer) Updates(state account.AccountState) ([]account.AccountUpdate, error) {
	if t.Value <= 0 {
		return nil, fmt.Errorf("transfer value must be positive")
	}

	fromAcc := state.GetAccount(t.From)
	if uint(t.Value) > fromAcc.Balance {
		return nil, fmt.Errorf("insufficient funds in %s", t.From)
	}

	return []account.AccountUpdate{
		{Name: t.From, BalanceChange: -t.Value},
		{Name: t.To, BalanceChange: t.Value},
	}, nil
}
