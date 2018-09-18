package main

import (
	//"encoding/hex"
	"fmt"
	//"github.com/syndtr/goleveldb/leveldb"
	"lurcury/account"
	"lurcury/db"
	"lurcury/core/block"
	"lurcury/core/transaction"
	"lurcury/http"
	//"lurcury/params"
	"lurcury/types"
	//"math/big"
	"time"
)
/*
func core_exp()(*types.CoreStruct){
	core_arg := &types.CoreStruct{}
	core_arg.Db = db.OpenDB("../dbdata")

	return core_arg
}
*/
func main(){

	core_arg := &types.CoreStruct{}
	core_arg.Db = db.OpenDB("../dbdata")
	b := block.BlockEncode(InitBlock())
	//fmt.Println("hexPut:",b)
	db.BlockHexPut(core_arg.Db, b.Hash, b)
	//fmt.Println( b.Hash)
	//tt := db.BlockHexGet(core_arg.Db, b.Hash)
	//fmt.Println("hexGet:",tt)
	go http.Server(core_arg)//core_arg)
	//fmt.Println(params.Chain)

	bb := transaction.ExpTransaction()

	core_arg.PendingTransaction = append(core_arg.PendingTransaction, bb)

	//account_tmp := &types.AccountData{}
	account_tmp := account.Account_exp()
	db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)
	
	tt := db.AccountHexGet(core_arg.Db, account_tmp.Address)
        fmt.Println("hexGet:",tt)
	c := types.BalanceData{}
	tt.Balance = append(tt.Balance, c)
	for index,element := range tt.Balance{
		fmt.Println(index)
		fmt.Println(element)
	}
	fmt.Println(tt)
	fmt.Println(core_arg.PendingTransaction)
	for i := 0; i<10000; i++{
		time.Sleep(1 * time.Second)
		fmt.Println(core_arg.PendingTransaction)
	}
	time.Sleep(30 * time.Second)
	fmt.Println(core_arg.PendingTransaction)

}



