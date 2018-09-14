package block

import (
        //"encoding/hex"
        "fmt"
	"lurcury/core/transaction"
	"lurcury/db"
	"lurcury/types"
        "testing"
)

func TestBlock(t *testing.T){
        core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../dbdata")
	bc := ExpBlock()
	fmt.Println(bc)
	hb:= BlockEncode(bc)
	sb:= BlockSign("ab70ef5f36dbfd9e403ed4ffd5b1c51dc7ce761ee21c8dc72570c6d73bb9412b0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba" ,hb)
	fmt.Println("hash:",bc.Hash)
	fmt.Println("sb:",sb)
	fmt.Println("hb:",hb)

	//流程
	//CreateNewBlock
	//=>BlockEncode
	//=>BlockSign(加簽章)
	//=>

        bb := transaction.ExpTransaction()
        core_arg.PendingTransaction = append(core_arg.PendingTransaction, bb)
	cb := CreateNewBlock(hb)

	//cb.Transaction = core_arg.PendingTransaction
	
	re := CreateBlockPOA(core_arg, cb, "ab70ef5f36dbfd9e403ed4ffd5b1c51dc7ce761ee21c8dc72570c6d73bb9412b0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba")
	fmt.Println(re)
	

	//fmt.Println("cb:",cb)

	hb2 := BlockEncode(cb)
	//fmt.Println(hb2)
	sb2 := BlockSign("ab70ef5f36dbfd9e403ed4ffd5b1c51dc7ce761ee21c8dc72570c6d73bb9412b0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba" ,hb2)
	fmt.Println(sb2)
	/*
	fmt.Println(sb2.Transaction[0])
*/
}


