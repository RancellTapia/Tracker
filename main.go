package main

import (
	"log"
	"time"
	keepers "tracker/keepers"
	"tracker/types"
)

func main() {
	tk := keepers.NewTransactionKeeper()

	transaction := tk.GetAllTransaction()


	printTransactions(transaction)

	tk.AddTransaction(&types.Transaction{
		Id: "jfhfhkjfkjhfjh",
		Date: time.Now(),
		Coin: "Atom",
		Value: 342,
	})

	log.Println("---------------------------------")

	printTransactions(transaction)
}

func printTransactions (t map[string]*types.Transaction){
	for id, v:= range t {
		log.Printf("%s -> Coin: %s Value: %f", id, v.Coin, v.Value)
	}
}