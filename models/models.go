package models

import "time"

// Wallet stores the information for a single wallet
type Wallet struct {
	UserID       string  `json:"user_id"`
	IsIdentified bool    `json:"is_identified"`
	Balance      float64 `json:"balance"`
	LastModified time.Time
}

// WalletTransaction stores information about a single transaction
type WalletTransaction struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	Amount          float64   `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
	TransactionType string    `json:"transaction_type"`
}

//NewWallet creates a new wallet
func NewWallet(userID string, isIdentified bool, balance float64) Wallet {
	return Wallet{
		UserID:       userID,
		IsIdentified: isIdentified,
		Balance:      balance,
		LastModified: time.Now(),
	}
}

//NewWalletTransaction creates a new wallet transaction
func NewWalletTransaction(id, userID string, amount float64, transactionType string) WalletTransaction {
	return WalletTransaction{
		ID:              id,
		UserID:          userID,
		Amount:          amount,
		CreatedAt:       time.Now(),
		TransactionType: transactionType,
	}
}
