package account

type AccountValue struct {
	Name    string
	Balance uint
}

type AccountState interface {
	GetAccount(name string) AccountValue
}

type AccountUpdate struct {
	Name          string
	BalanceChange int
}

type InMemoryState struct {
	accounts map[string]uint
}

func NewInMemoryState(initial []AccountValue) *InMemoryState {
	state := &InMemoryState{
		accounts: make(map[string]uint),
	}
	for _, acc := range initial {
		state.accounts[acc.Name] = acc.Balance
	}
	return state
}

func (s *InMemoryState) GetAccount(name string) AccountValue {
	balance, exists := s.accounts[name]
	if !exists {
		return AccountValue{Name: name, Balance: 0}
	}
	return AccountValue{Name: name, Balance: balance}
}

func (s *InMemoryState) ApplyUpdates(updates []AccountUpdate) {
	for _, update := range updates {
		current := s.GetAccount(update.Name).Balance
		newBalance := int(current) + update.BalanceChange
		if newBalance < 0 {
			newBalance = 0
		}
		s.accounts[update.Name] = uint(newBalance)
	}
}

func (s *InMemoryState) Snapshot() []AccountValue {
	var snapshot []AccountValue
	for name, balance := range s.accounts {
		snapshot = append(snapshot, AccountValue{
			Name:    name,
			Balance: balance,
		})
	}
	return snapshot
}
