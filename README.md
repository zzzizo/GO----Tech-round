# Live Coding Challenge: Multi-Signature Wallet Service

## Overview

You will implement a secure service for managing multi-signature cryptocurrency wallets. This service should handle transaction submission, signature collection, and transaction execution with proper security and concurrency considerations.

## Problem Requirements

### Core Functionality

Your service must support:

1. **Transaction Submission**: Accept new multi-signature transaction requests
2. **Signature Collection**: Allow authorized users to add signatures to pending transactions
3. **Transaction Execution**: Execute transactions when sufficient signatures are collected
4. **Audit Logging**: Log all operations for compliance and security
5. **Concurrent Safety**: Handle multiple simultaneous operations safely

### Technical Requirements

- Use the provided interfaces and implementations
- Implement proper error handling
- Ensure thread-safe operations
- Use Go context for operation lifecycle management
- Follow Go best practices and conventions

## Getting Started

### Provided Code

The following interfaces and helper implementations are available for your use:

```go
package main

import (
    "context"
    "crypto/ecdsa"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "sync"
    "time"
)

// Data Structures
type Transaction struct {
    ID          string    `json:"id"`
    From        string    `json:"from"`
    To          string    `json:"to"`
    Amount      string    `json:"amount"` // Using string for precision
    RequiredSigs int      `json:"required_sigs"`
    Signatures   []string `json:"signatures"`
    Status      string    `json:"status"` // "pending", "ready", "executed", "failed"
    CreatedAt   time.Time `json:"created_at"`
}

type AuditLog struct {
    Timestamp time.Time `json:"timestamp"`
    Action    string    `json:"action"`
    UserID    string    `json:"user_id"`
    TxID     string    `json:"tx_id"`
    Details   string    `json:"details"`
}

// Service Interfaces
type SignatureValidator interface {
    ValidateSignature(pubKey, message, signature string) error
}

type AuditLogger interface {
    LogAction(ctx context.Context, action, userID, txID, details string) error
}

type TransactionStore interface {
    StoreTransaction(ctx context.Context, tx *Transaction) error
    GetTransaction(ctx context.Context, txID string) (*Transaction, error)
    UpdateTransactionStatus(ctx context.Context, txID, status string) error
}

// Helper Implementations (Ready to Use)
type ECDSAValidator struct{}
func (e *ECDSAValidator) ValidateSignature(pubKey, message, signature string) error {
    // Simplified validation for interview
    if len(signature) < 64 {
        return errors.New("invalid signature length")
    }
    return nil
}

type ConsoleAuditor struct{}
func (c *ConsoleAuditor) LogAction(ctx context.Context, action, userID, txID, details string) error {
    log.Printf("AUDIT: %s | User: %s | TX: %s | %s", action, userID, txID, details)
    return nil
}

type InMemoryStore struct {
    mu           sync.RWMutex
    transactions map[string]*Transaction
}

func NewInMemoryStore() *InMemoryStore {
    return &InMemoryStore{
        transactions: make(map[string]*Transaction),
    }
}

func (s *InMemoryStore) StoreTransaction(ctx context.Context, tx *Transaction) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.transactions[tx.ID] = tx
    return nil
}

func (s *InMemoryStore) GetTransaction(ctx context.Context, txID string) (*Transaction, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    tx, exists := s.transactions[txID]
    if !exists {
        return nil, errors.New("transaction not found")
    }
    // Return a copy to avoid race conditions
    txCopy := *tx
    return &txCopy, nil
}

func (s *InMemoryStore) UpdateTransactionStatus(ctx context.Context, txID, status string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    if tx, exists := s.transactions[txID]; exists {
        tx.Status = status
        return nil
    }
    return errors.New("transaction not found")
}

// YOUR IMPLEMENTATION STARTS FROM HERE
// WHAT YOU HAVE TO DO: Implement a MultiSigWalletService that can:
// 1. Accept new multi-signature transaction requests
// 2. Allow users to add signatures to pending transactions  
// 3. Execute transactions when they have sufficient signatures
// 4. Handle concurrent operations safely
// 5. Log all operations for audit purposes
//
// Design the service structure and methods as you see fit.

func main() {
    
    fmt.Println("Multi-signature wallet service implementation")
}
```

### Implementation Guidelines

1. **Design First**: Take a moment to think about the service structure and required methods
2. **Use Provided Components**: Leverage the interfaces and implementations provided
3. **Focus on Core Logic**: Don't spend time on boilerplate - focus on the essential functionality
4. **Ask Questions**: Clarify requirements if anything is unclear
5. **Explain Your Approach**: Communicate your thinking as you implement

## What You Can Use

- Standard Go library packages
- The provided interfaces and implementations
- Any reasonable design patterns
- Documentation access (Go docs, etc.)

## Expected Behavior

### Transaction Lifecycle
1. **Pending**: New transactions start in pending status
2. **Ready**: Transactions with sufficient signatures become ready for execution
3. **Executed**: Successfully processed transactions
4. **Failed**: Transactions that couldn't be executed

### Security Considerations
- Validate all signatures before accepting them
- Ensure atomic operations where necessary
- Prevent race conditions in concurrent access
- Log all significant operations

## Notes

- This is a collaborative exercise - feel free to discuss your approach
- Perfect completion isn't expected - focus on demonstrating your thought process
- Code doesn't need to be production-ready, but should show good Go practices

---

**Good luck!** 🚀
