package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/ProninIgorr/epayments-restapi/auth"
	"github.com/ProninIgorr/epayments-restapi/handlers"

	"github.com/gorilla/mux"
)

func main() {

	//Generate X-Digest
	h := hmac.New(sha1.New, []byte(auth.SecretKey))
	digest := hex.EncodeToString(h.Sum(nil))
	fmt.Println("X-Digest:", digest)

	//Start server
	router := mux.NewRouter()
	router.HandleFunc("/wallet/exists", handlers.CheckWalletExists).Methods("GET")
	router.HandleFunc("/wallet/replenish", handlers.ReplenishWallet).Methods("POST")
	router.HandleFunc("/wallet/transactions", handlers.GetTransactions).Methods("GET")
	router.HandleFunc("/wallet/balance", handlers.GetBalance).Methods("GET")
	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//Для проверки работы программы можно использовать curl
//curl -X POST -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c257dac" http://localhost:8000/wallet/exists
//curl -X GET -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c257dac" http://localhost:8000/wallet/balance
//curl -X POST -H "X-UserId: user1" -H "X-Digest: f92a65136e9c7b7a31590f87b850c96f7c25ac" -d '{"amount": 100}' http://localhost:8000/wallet/replenish
