package keepers

import (
	"time"
	types "tracker/types"
)


type TransactionKeeper struct {
	store map[string]*types.Transaction
}

func NewTransactionKeeper() *TransactionKeeper {
	transactionKeeper := &TransactionKeeper{}
	transactionKeeper.store = make(map[string]*types.Transaction)

	transactionKeeper.store["jdskhgkhgdhgfhgjhgsd"] = &types.Transaction{
		Id: "jdskhgkhgdhgfhgjhgsd",
		Date: time.Now(),
		Coin: "osmosis",
		Value: 1345,
	}
	 
	return transactionKeeper
}

func (tk *TransactionKeeper) AddTransaction(transaction *types.Transaction) {
	tk.store[transaction.Id] = transaction
}

func (tk *TransactionKeeper) GetAllTransaction() map[string]*types.Transaction {
	return tk.store
}

func (tk *TransactionKeeper) GetTransaction(id string) *types.Transaction{
	return tk.store[id]
}

func (tk *TransactionKeeper) DeleteTransaction(id string) {
	
}

func (tk *TransactionKeeper) UpdateTransaction(id string) {
	
}