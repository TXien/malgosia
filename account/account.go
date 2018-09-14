package account

import (
	"math/big"
	"encoding/hex"
	//"fmt"
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
                Balance:*big.NewInt(100000),
        }
        c := types.BalanceData{
                Token:"deh",
                Balance:*big.NewInt(100000),
        }
        s := types.AccountData{
               Address:"264411884d6d2aca8ca2d2a77c9dc95ffdcee529",
                Nonce:0,
                Balance:[]types.BalanceData{b,c},
                Transaction:[]types.TransactionJson{},
        }
        //fmt.Println(b,s)
        return s
}

func InitAccount(address string, nonce int)(types.AccountData){
	/*
	b := types.BalanceData{
		//token:"def",
		//balance:0
	}
	*/
	s := types.AccountData{
		Address:address,
		Nonce:nonce,
		Balance:[]types.BalanceData{},
		Transaction:[]types.TransactionJson{},
	}
	//fmt.Println(b,s)
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
