# Concurrent Transaction Executor

A Go implementation of a deterministic transaction processor for blockchain-like systems, handling concurrent transaction execution while maintaining account consistency.

## Features

- Processes blocks of transactions in order
- Maintains account state consistency
- Handles transaction failures gracefully
- Provides deterministic results
- Example transfer transaction implementation

## Prerequisites

- Go 1.24+ (https://golang.org/dl/)
- Git

## Installation

##### 1. Clone the repository:
```bash
git clone https://github.com/henrymatteng/transaction-executor.git
cd transaction-executor
```
##### 2. Run the project
```go
go run ./main.go
```
##### 3. Run the tests
```go
go test -v ./executor/
```
