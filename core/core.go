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
	"lurcury/params"
	"lurcury/types"
	//"math/big"
	"os"
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
        go http.Server(core_arg)
	initAccount(*core_arg, GenesisBlock())//InitBlock())
	pri:="219a634773d787cfbaf1e5c915d56b14be2a3695ed8e46bbeb01573bf211d0ef8773580834eb42a2f2ee856b029a88dfee639e27f08b1e0235f8eb04eecf4089"
	tmpBlock := GenesisBlock()
	for i:=0;i!=-1;i++{
		if(len(core_arg.PendingTransaction)>0){
			tmpBlock = block.CreateBlockPOA(core_arg, tmpBlock, pri)
			time.Sleep(1 * time.Second)
		}else{
			fmt.Println("no transaction")
			time.Sleep(1 * time.Second)
		}
	}
	if(os.Args[1]=="init"){
		b := block.BlockEncode(InitBlock())
		db.BlockHexPut(core_arg.Db, b.Hash, b)
	}
	//go http.Server(core_arg)

	exTrans := transaction.ExpTransaction()
	core_arg.PendingTransaction = append(core_arg.PendingTransaction, exTrans)
	genesis := GenesisBlock()
	account_tmp := account.InitAccount(genesis.Allocate[0].Address, params.Chain().Version.Sue.FeeToken, genesis.Allocate[0].Amount)
	fmt.Println(genesis.Allocate[0].Amount)
	db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)

	accountInfo := db.AccountHexGet(core_arg.Db, account_tmp.Address)
        fmt.Println("hexGet:",accountInfo.Token[0].Balance)
	

	/*
	balanceData := types.BalanceData{}
	accountInfo.Balance = append(accountInfo.Balance, balanceData)
	for index,element := range accountInfo.Balance{
		fmt.Println("index:",index)
		fmt.Println("token:",element)
	}
	fmt.Println(accountInfo)
	*/
	/*
	fmt.Println(core_arg.PendingTransaction)
	for i := 0; i<10000; i++{
		time.Sleep(1 * time.Second)
		fmt.Println(core_arg.PendingTransaction)
	}
	*/
	time.Sleep(30 * time.Second)
	//fmt.Println(core_arg.PendingTransaction)
	
}

