package executor

import (
	"github.com/henrymatteng/transaction_executor/account"
	"github.com/henrymatteng/transaction_executor/block"
)

const NumWorkers = 4

var genesisAccounts = []account.AccountValue{
	{Name: "A", Balance: 20},
	{Name: "B", Balance: 30},
	{Name: "C", Balance: 40},
}

func ExecuteBlock(blk block.Block) ([]account.AccountValue, error) {
	state := make(map[string]uint)
	for _, acc := range genesisAccounts {
		state[acc.Name] = acc.Balance
	}

	modifiedAccounts := make([][]string, len(blk.Transactions))

	var anyError error
	for i, txn := range blk.Transactions {
		snapshot := make(map[string]uint)
		for k, v := range state {
			snapshot[k] = v
		}

		updates, err := txn.Updates(&accountStateView{state: snapshot})
		if err != nil {
			anyError = err
			continue
		}

		var modified []string
		for _, update := range updates {
			modified = append(modified, update.Name)

			current := state[update.Name]
			newBalance := int(current) + update.BalanceChange
			if newBalance < 0 {
				newBalance = 0
			}
			state[update.Name] = uint(newBalance)
		}
		modifiedAccounts[i] = modified
	}

	var result []account.AccountValue
	for name, balance := range state {
		result = append(result, account.AccountValue{
			Name:    name,
			Balance: balance,
		})
	}
	return result, anyError
}

type accountStateView struct {
	state map[string]uint
}

func (v *accountStateView) GetAccount(name string) account.AccountValue {
	balance, exists := v.state[name]
	if !exists {
		return account.AccountValue{Name: name, Balance: 0}
	}
	return account.AccountValue{Name: name, Balance: balance}
}
