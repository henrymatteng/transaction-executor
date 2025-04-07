package executor

import (
	"testing"

	"github.com/henrymatteng/transaction_executor/block"
	"github.com/henrymatteng/transaction_executor/transaction"
)

func TestExecuteBlock(t *testing.T) {
	tests := []struct {
		name         string
		transactions []block.Transaction
		want         map[string]uint
		wantErr      bool
	}{
		{
			name: "Single successful transfer",
			transactions: []block.Transaction{
				transaction.Transfer{From: "A", To: "B", Value: 5},
			},
			want: map[string]uint{
				"A": 15,
				"B": 35,
				"C": 40,
			},
		},
		{
			name: "Failed transfer due to insufficient funds",
			transactions: []block.Transaction{
				transaction.Transfer{From: "A", To: "B", Value: 50},
			},
			want: map[string]uint{
				"A": 20,
				"B": 30,
				"C": 40,
			},
			wantErr: true,
		},
		{
			name: "Concurrent non-conflicting transfers",
			transactions: []block.Transaction{
				transaction.Transfer{From: "A", To: "B", Value: 5},
				transaction.Transfer{From: "C", To: "D", Value: 10},
			},
			want: map[string]uint{
				"A": 15,
				"B": 35,
				"C": 30,
				"D": 10,
			},
		},
		{
			name: "Conflicting transfers processed serially",
			transactions: []block.Transaction{
				transaction.Transfer{From: "A", To: "B", Value: 5},
				transaction.Transfer{From: "B", To: "C", Value: 10},
			},
			want: map[string]uint{
				"A": 15,
				"B": 25,
				"C": 50,
			},
		},
		{
			name: "Mixed success and failure",
			transactions: []block.Transaction{
				transaction.Transfer{From: "A", To: "B", Value: 5},  // Success
				transaction.Transfer{From: "B", To: "C", Value: 40}, // Fail
				transaction.Transfer{From: "A", To: "C", Value: 5},  // Success
			},
			want: map[string]uint{
				"A": 10,
				"B": 35,
				"C": 45,
			},
			wantErr: true,
		},
		{
			name: "New account creation",
			transactions: []block.Transaction{
				transaction.Transfer{From: "A", To: "X", Value: 5},
			},
			want: map[string]uint{
				"A": 15,
				"B": 30,
				"C": 40,
				"X": 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blk := block.Block{Transactions: tt.transactions}
			got, err := ExecuteBlock(blk)

			if (err != nil) != tt.wantErr {
				t.Errorf("ExecuteBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Convert result to map for easier comparison
			gotMap := make(map[string]uint)
			for _, acc := range got {
				gotMap[acc.Name] = acc.Balance
			}

			// Check expected accounts
			for name, wantBalance := range tt.want {
				if gotBalance, exists := gotMap[name]; !exists || gotBalance != wantBalance {
					t.Errorf("Account %s balance = %d, want %d", name, gotBalance, wantBalance)
				}
			}

			// Verify no extra accounts were created
			for name := range gotMap {
				if _, exists := tt.want[name]; !exists && name != "D" && name != "X" {
					t.Errorf("Unexpected account created: %s", name)
				}
			}
		})
	}
}

func TestEmptyBlock(t *testing.T) {
	blk := block.Block{Transactions: []block.Transaction{}}
	got, err := ExecuteBlock(blk)
	if err != nil {
		t.Errorf("ExecuteBlock() unexpected error: %v", err)
	}

	// Should return genesis state
	expected := map[string]uint{"A": 20, "B": 30, "C": 40}
	for _, acc := range got {
		if expected[acc.Name] != acc.Balance {
			t.Errorf("Account %s balance = %d, want %d", acc.Name, acc.Balance, expected[acc.Name])
		}
	}
}
