package account

import (

	//"math/big"
	"encoding/hex"
	//"fmt"
        "lurcury/db"
	eddsa "lurcury/crypto/eddsa"
	crypto "lurcury/crypto"
	"lurcury/types"
	//"lurcury/core/transaction"
)
/*
type BalanceData struct{
	token string
	balance uint
}

type AccountData struct{
	address string
	nonce uint
	balance []BalanceData
	transaction []string
}
*/
func GenesisAccount(core_arg types.CoreStruct,address string,balance string)(bool){
        s := types.AccountData{
                Nonce:0,
		Balance:balance,
        }
        db.AccountHexPut(core_arg.Db, address, s)
	return true
}

func NewAccount()(string, string, string){
	pri ,pub := eddsa.EddsaGenerateKey()
	addr := crypto.KeyToAddress(pri)
	return hex.EncodeToString(pri), hex.EncodeToString(pub), (hex.EncodeToString(addr))
}

func Account_exp()(types.AccountData){
//privateKey: 01985a94b7077b77a28eec1993bdc232c6ebe323703ff59368fd1062a2599e1d1c6cd46e521a769bdc11c67df770a2605759d54b59d79c16d4c972dc88edb36d 
//publicKey: 1c6cd46e521a769bdc11c67df770a2605759d54b59d79c16d4c972dc88edb36d 
//address: gx21ece2f0b9d99cde7254a0309d8065bf1ad070d6
        b := types.BalanceData{
                Token:"def",
                Balance:"100000",//big.NewInt(100000),
        }
        c := types.BalanceData{
                Token:"deh",
                Balance:"100000",//big.NewInt(100000),
        }
        s := types.AccountData{
               Address:"264411884d6d2aca8ca2d2a77c9dc95ffdcee529",
                Nonce:0,
		Balance:"1000000000000000000",
                Token:[]types.BalanceData{b,c},
                Transaction:[]types.TransactionJson{},
        }
        //fmt.Println(b,s)
        return s
}

func InitAccount(address string, tokenName string, amount string/**big.Int*/)(types.AccountData){
	
	b := types.BalanceData{
		Token: tokenName,
		Balance: amount,
	}
	
	s := types.AccountData{
		Address:address,
		Nonce:0,
		Token:[]types.BalanceData{b},
		Transaction:[]types.TransactionJson{},
	}
	return s
}

/*
func main(){
	b := []BalanceData{{token:"abc", balance: 1}}
	b = append(b,BalanceData{token:"abc", balance: 1})
	g := AccountData{address:"123",nonce:1,balance:b}
	fmt.Println(g)
}
*/
