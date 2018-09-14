package transaction

import (
        //"encoding/hex"
        "fmt"
	"lurcury/account"
	"lurcury/crypto"
	"lurcury/db"
        "lurcury/types"
        //"math/big"
        "reflect"
        "testing"
)

func TestTransaction(t *testing.T){
        check := func(f string, got, want interface{}) {
                if !reflect.DeepEqual(got, want) {
                        t.Errorf("%s mismatch: got %v, want %v", f, got, want)
                }
        }
/*
	b := ExpTransaction()
        //fmt.Println(b)
	fmt.Println("sign verify test:",VerifyTransactionSign(b))
*/
	core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../dbdata")

	//初始化金額
	
        account_tmp := account.Account_exp()
	fmt.Println(account_tmp)
	db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)
/*
	account_tmp2 := account.Account_exp()
	account_tmp2.Address = "5ee464a101d58877f00957eff452c148e7f75833"
	//fmt.Println(account_tmp.Nonce)
	db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)
        fmt.Println("put1",db.AccountHexGet(core_arg.Db, "5ee464a101d58877f00957eff452c148e7f75833"))
        //account_tmp2.Address = "eee"
	fmt.Println(account_tmp2.Address[2:])
	db.AccountHexPut(core_arg.Db, account_tmp2.Address, account_tmp2)
	fmt.Println("put2",db.AccountHexGet(core_arg.Db, account_tmp2.Address))
	fmt.Println("put2",db.AccountHexGet(core_arg.Db, "5ee464a101d58877f00957eff452c148e7f75833"))

	fmt.Println("account put and get test:",db.AccountHexGet(core_arg.Db, account_tmp.Address))
*/
	pp := ExpTransaction()
/*
	fmt.Println("sign verify test:",VerifyTransactionSign(pp))
	fmt.Println("test token amount:", pp.Out[0].Token)
	fmt.Println("from address test:",crypto.KeyToAddress_hex(pp.PublicKey))
	fmt.Println("Nonce:",account_tmp.Nonce)

	fmt.Println("sign verify test:",VerifyTransactionSign(pp))
	fmt.Println("pp:",pp)
*/
	m1, m2 := VerifyTransactionBalanceAndNonce(*core_arg, pp)
	fmt.Println("verify balance and nonce:",m1)
	fmt.Println("balanceresult:",m2)


	a3 := db.AccountHexGet(core_arg.Db, account_tmp.Address)
	fmt.Println("test for verify balance and nonce result:",a3)
	fmt.Println(pp.PublicKey)
	a4 := db.AccountHexGet(core_arg.Db, crypto.KeyToAddress_hex(pp.PublicKey))
	fmt.Println(a4)
	
	check("go","123","123")

}
