package models

type Block struct {
	Transactions []Transaction
}

type Transaction interface {
	Updates(AccountState) ([]AccountUpdate, error)
}

type AccountUpdate struct {
	Name          string
	BalanceChange int
}

type AccountValue struct {
	Name    string
	Balance uint
}

type AccountState interface {
	GetAccount(name string) AccountValue
}
