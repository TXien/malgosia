package types

import (
	"math/big"
	"github.com/syndtr/goleveldb/leveldb"
	//"lurcury/core/block"
	//"lurcury/core/transaction"
	"lurcury/params"
)
/*
type CoreStruct struct {
        Test string
        PendingTransaction []transaction.TransactionJson
        Db *leveldb.DB
        Config *params.ChainConfigStructure
        PendingBlock []block.BlockJson
}
*/



type BalanceData struct{
        Token string
        Balance big.Int
}

type AccountData struct{
        Address string
        Nonce int
        Balance []BalanceData
        Transaction []TransactionJson
}

type CoreStruct struct {
	Test string
	PendingTransaction []TransactionJson
        Db *leveldb.DB
        Config *params.ChainConfigStructure
	PendingBlock []BlockJson
}
