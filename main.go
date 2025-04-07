package main

import (
	"fmt"

	"github.com/henrymatteng/transaction_executor/block"
	"github.com/henrymatteng/transaction_executor/executor"
	"github.com/henrymatteng/transaction_executor/transaction"
)

func main() {
	block1 := block.Block{
		Transactions: []block.Transaction{
			transaction.Transfer{From: "A", To: "B", Value: 5},
			transaction.Transfer{From: "B", To: "C", Value: 10},
			transaction.Transfer{From: "B", To: "C", Value: 30}, // should fail
		},
	}

	result1, err := executor.ExecuteBlock(block1)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Example 1 Results:")
	for _, acc := range result1 {
		fmt.Printf("%s: %d\n", acc.Name, acc.Balance)
	}

	block2 := block.Block{
		Transactions: []block.Transaction{
			transaction.Transfer{From: "A", To: "B", Value: 5},
			transaction.Transfer{From: "C", To: "D", Value: 10},
		},
	}

	result2, _ := executor.ExecuteBlock(block2)
	fmt.Println("\nExample 2 Results:")
	for _, acc := range result2 {
		fmt.Printf("%s: %d\n", acc.Name, acc.Balance)
	}
}
