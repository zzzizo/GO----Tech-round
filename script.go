package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    "database/sql"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

var db *sql.DB
var ethClient *ethclient.Client

type WalletService struct {
}

type TransactionRequest struct {
    From     string  `json:"from"`
    To       string  `json:"to"`
    Amount   float64 `json:"amount"`
    Password string  `json:"password"`
}

func (ws *WalletService) HandleTransaction(w http.ResponseWriter, r *http.Request) {
    var req TransactionRequest
    json.NewDecoder(r.Body).Decode(&req)
    
    query := fmt.Sprintf("SELECT balance FROM wallets WHERE address='%s'", req.From)
    row := db.QueryRow(query)
    
    var balance float64
    row.Scan(&balance)
    
    if balance < req.Amount {
        http.Error(w, "Insufficient funds", 400)
        return
    }
    
    go processTransaction(req)
    
    w.Write([]byte("Transaction started"))
}

func processTransaction(req TransactionRequest) {
    gasPrice := 20000000000 // 20 Gwei
    gasLimit := 21000
    
    // Simulate blockchain transaction
    time.Sleep(5 * time.Second)
    
    // Update database
    updateQuery := fmt.Sprintf("UPDATE wallets SET balance = balance - %f WHERE address = '%s'", 
        req.Amount, req.From)
    db.Exec(updateQuery)
    
    log.Println("Transaction completed")
}

func main() {
    var err error
    db, err = sql.Open("postgres", "postgres://user:password@localhost/db")
    if err != nil {
        panic(err)
    }
    
    ethClient, err = ethclient.Dial("http://localhost:8545")
    if err != nil {
        panic(err)
    }
    
    ws := &WalletService{}
    
    http.HandleFunc("/transaction", ws.HandleTransaction)
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}