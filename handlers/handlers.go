package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ProninIgorr/epayments-restapi/models"
)

const (
	MaxUnidentifiedBalance = 10000
	MaxIdentifiedBalance   = 100000
)

// create a map of wallets and a slice of transactions
var (
	Wallets = map[string]models.Wallet{
		"user1": models.NewWallet("user1", true, 100),
		"user2": models.NewWallet("user2", false, 50),
	}

	Transactions = []models.WalletTransaction{
		{ID: "t1", UserID: "user1", Amount: 100, CreatedAt: time.Now(), TransactionType: "replenishment"},
		{ID: "t2", UserID: "user2", Amount: 50, CreatedAt: time.Now(), TransactionType: "replenishment"},
	}
)

// CheckWalletExists checks if the wallet exists
func CheckWalletExists(w http.ResponseWriter, r *http.Request) {

	userID := r.Header.Get("X-UserId")
	if userID == "" {
		http.Error(w, "X-UserId header is required", http.StatusBadRequest)
		return
	}

	if _, ok := Wallets[userID]; ok {
		w.Write([]byte("Wallet exists"))
		return
	}
	http.Error(w, "Wallet does not exist", http.StatusBadRequest)
	w.Write([]byte("Wallet does not exist"))

}

// ReplenishWallet replenishes the wallet
func ReplenishWallet(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-UserId")
	if userID == "" {
		http.Error(w, "X-UserId header is required", http.StatusBadRequest)
		return
	}

	wallet, ok := Wallets[userID]
	if !ok {
		http.Error(w, "Wallet does not exist", http.StatusBadRequest)
		return
	}

	var transaction models.WalletTransaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	maxBalance := MaxUnidentifiedBalance
	if wallet.IsIdentified {
		maxBalance = MaxIdentifiedBalance
	}

	if int(wallet.Balance+transaction.Amount) > maxBalance {
		http.Error(w, "Exceeded maximum balance", http.StatusBadRequest)
		return
	}

	wallet.Balance += transaction.Amount
	wallet.LastModified = time.Now()
	Wallets[userID] = wallet

	transaction.CreatedAt = time.Now()
	Transactions = append(Transactions, transaction)

	w.Write([]byte("Success"))
}

// GetTransactions gets the transactions for the current month
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-UserId")
	if userID == "" {
		http.Error(w, "X-UserId header is required", http.StatusBadRequest)
		return
	}

	_, ok := Wallets[userID]
	if !ok {
		http.Error(w, "Wallet does not exist", http.StatusBadRequest)
		return
	}

	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	monthEnd := monthStart.AddDate(0, 1, 0)

	var currentMonthTransactions []models.WalletTransaction
	for _, transaction := range Transactions {
		if transaction.UserID == userID && transaction.CreatedAt.After(monthStart) && transaction.CreatedAt.Before(monthEnd) {
			currentMonthTransactions = append(currentMonthTransactions, transaction)
		}
	}

	response, _ := json.Marshal(currentMonthTransactions)
	w.Write(response)
}

// GetBalance gets the balance of the wallet
func GetBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-UserId")
	if userID == "" {
		http.Error(w, "X-UserId header is required", http.StatusBadRequest)
		return
	}

	wallet, ok := Wallets[userID]
	if !ok {
		http.Error(w, "Wallet does not exist", http.StatusBadRequest)
		return
	}

	response, _ := json.Marshal(map[string]float64{"balance": wallet.Balance})
	w.Write(response)
}
