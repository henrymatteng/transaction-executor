package block

import (
	"github.com/henrymatteng/transaction_executor/account"
)

type Transaction interface {
	Updates(state account.AccountState) ([]account.AccountUpdate, error)
}

type Block struct {
	Transactions []Transaction
}
