package transaction

import(
	"encoding/hex"
	"fmt"
	eddsa "lurcury/crypto/eddsa"
	"lurcury/crypto"
	"lurcury/db"
	"lurcury/params"
	"lurcury/types"
	"math/big"
	"strings"
	"strconv"
	//"time"
	//"reflect"
)


func NewTransaction(To string, Token string, Balance big.Int , Nonce int, Fee big.Int, Type string, Input string)(types.TransactionJson){
	out := []types.TransactionOut{
		{Balance: Balance, Token: Token, Vout: 0},
	}
	trans := types.TransactionJson{
		Out: out,
		To: To,
		Nonce: Nonce,
		Fee: Fee,
		Type: Type,
		Input: Input,
	}
	return trans
}

func SignTransaction(pri []byte, Transaction types.TransactionJson)(types.TransactionJson){
	Transaction.Tx = string(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	sign := eddsa.EddsaSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
	r := eddsa.EddsaKeyToPublicKey(pri)
	Transaction.Sign = hex.EncodeToString(sign)
	Transaction.PublicKey = hex.EncodeToString(r)
	return Transaction
}

func SignTransaction_hex(prik string, Transaction types.TransactionJson)(types.TransactionJson){
	pri,_ := hex.DecodeString(prik)
	Transaction.Tx = string(crypto.Keccak256([]byte(EncodeForSign(Transaction))))
        sign := eddsa.EddsaSign(pri, crypto.Keccak256([]byte(EncodeForSign(Transaction))))
        r := eddsa.EddsaKeyToPublicKey(pri)
        Transaction.Sign = hex.EncodeToString(sign)
        Transaction.PublicKey = hex.EncodeToString(r)
        return Transaction
}

func EncodeForSign(Transaction types.TransactionJson)(string){
	//fmt.Println("check", Transaction)
        to        := "gx"+StringTransactionEncode(Transaction.To/*[2:]*/, 42)
        nonce_tmp := strconv.FormatInt(int64(Transaction.Nonce), 16)
        nonce     := "hx"+StringTransactionEncode(nonce_tmp, 32)
        fee_tmp   := strconv.FormatInt(int64(Transaction.Fee.Uint64()), 16)
        fee       := "ix"+StringTransactionEncode(fee_tmp, 32)
	//fmt.Println(Transaction.Type)
        typ_tmp:= hex.EncodeToString([]byte(Transaction.Type))
        typ       := "kx"+StringTransactionEncode(string(typ_tmp), 8)
	//fmt.Println("typ:",typ)
        input_tmp_str := hex.EncodeToString([]byte(Transaction.Input))
        input_tmp := strconv.FormatInt(int64(len(input_tmp_str)), 16)
        input     := "lx"+StringTransactionEncode(input_tmp,3) + string(input_tmp_str[:])
        outResult := ""
        token     := ""
        balance   := ""
        vout      := ""
        for i :=0; i< len(Transaction.Out); i++ {
                token_tmp:= hex.EncodeToString([]byte(Transaction.Out[i].Token))
                token     = StringTransactionEncode(string(token_tmp), 8)
                balance_tmp   := strconv.FormatInt(int64(Transaction.Out[i].Balance.Uint64()), 16)
                balance   = StringTransactionEncode(balance_tmp, 32)
                vout_tmp := strconv.Itoa(Transaction.Out[i].Vout)
                vout = StringTransactionEncode(vout_tmp,3)
                outResult = outResult+ "px" + vout + token + balance
        }
	re := to + nonce + fee + typ + outResult + input
	return re
}

func EncodeTransaction(Transaction types.TransactionJson)(string){
	to        := "gx"+StringTransactionEncode(Transaction.To[2:], 42)
	nonce_tmp := strconv.FormatInt(int64(Transaction.Nonce), 16)
	nonce     := "hx"+StringTransactionEncode(nonce_tmp, 32)
	fee_tmp   := strconv.FormatInt(int64(Transaction.Fee.Uint64()), 16)
	fee       := "ix"+StringTransactionEncode(fee_tmp, 32)
	typ_tmp  := hex.EncodeToString([]byte(Transaction.Type))
	typ       := "kx"+StringTransactionEncode(string(typ_tmp), 8)
	input_tmp_str := hex.EncodeToString([]byte(Transaction.Input))
	input_tmp := strconv.FormatInt(int64(len(input_tmp_str)), 16)
	input     := "lx"+StringTransactionEncode(input_tmp,3) + string(input_tmp_str[:])
	outResult := ""
	token     := ""
	balance   := ""
	vout      := ""
	for i :=0; i< len(Transaction.Out); i++ {
		token_tmp:= hex.EncodeToString([]byte(Transaction.Out[i].Token))
		token     = StringTransactionEncode(string(token_tmp), 8)
		balance_tmp   := strconv.FormatInt(int64(Transaction.Out[i].Balance.Uint64()), 16)
		balance   = StringTransactionEncode(balance_tmp, 32)
		vout_tmp := strconv.Itoa(Transaction.Out[i].Vout)
		vout = StringTransactionEncode(vout_tmp,3)
		outResult = outResult+ "px" + vout + token + balance
	}
	sign      := "mx"+StringTransactionEncode(Transaction.Sign, 128)
	publicKey := "nx"+StringTransactionEncode(Transaction.PublicKey, 64)
	tx        := "gx"+StringTransactionEncode(Transaction.Tx, 64)
	//tx        := "gx"+StringTransactionEncode(Transaction.Tx, 64)

	re := to+nonce+fee+typ + outResult + input + sign + publicKey + tx
	return re
}

func DecodeTransaction(transaction string)(types.TransactionJson){
	g := strings.Index(transaction, "gx")
	h := strings.Index(transaction, "hx")
	i := strings.Index(transaction, "ix")
	k := strings.Index(transaction, "kx")
	l := strings.Index(transaction, "lx")
	m := strings.Index(transaction, "mx")
	n := strings.Index(transaction, "nx")
	p := strings.Index(transaction, "px")
	//fmt.Println(p,g,h,i,k,l,m,n)
	To := "gx"+transaction[g+4:g+44]
	Nonce,_ := strconv.Atoi(transaction[h+2:h+34])
	Fee := new(big.Int)
	Fee.SetString(transaction[i+2:i+34],10)
	Type_tmp,_ := hex.DecodeString(transaction[k+2:k+10])
	//fmt.Println("ch:",transaction[k+2:k+10])
	Type := string(Type_tmp)
	input_length, _ := strconv.ParseInt("0x"+transaction[l+2:l+5], 0, 64)
	Input_tmp,_ := hex.DecodeString(transaction[l+5:l+5+int(input_length)])
	Input := string(Input_tmp)
	Sign := transaction[m+2:m+130]
	PublicKey := transaction[n+2:n+66]

	out := transaction[p:p+45]
	Vout,_ := strconv.ParseInt(out[2:5],16, 10)
	Token,_ := hex.DecodeString(out[5:13])
	//fmt.Println(string(Token))
	balance := new(big.Int)
	balance, _ = balance.SetString(out[13:45], 16)
	outJson := []types.TransactionOut{
		{
			Token:string(Token),
			Balance:*balance,
			Vout:int(Vout)},
	}

	//outJson = append(outJson, TransactionOut{Token:string(Token)})
	//fmt.Println(outJson[1].Token)
        trans := types.TransactionJson{
                Out: outJson,
                To: To,
                Nonce: Nonce,
                Fee: *Fee,
                Type: Type,
                Input: Input,
                Sign: Sign,
                PublicKey:PublicKey,
        }
	return trans
	/*
	fmt.Println("to:", transaction[g:g+44])
	fmt.Println("nonce:", transaction[h:h+34])
	fmt.Println("fee:", transaction[i:i+34])
	fmt.Println("type:", transaction[k:k+34])
	input_length, _ := strconv.ParseInt("0x"+transaction[l+2:l+5], 0, 64)
	fmt.Println("input:", transaction[l+5:l+5+int(input_length)])
	fmt.Println("sign:", transaction[m:m+130])
	fmt.Println("publicKey:", transaction[n:n+66])
	fmt.Println("out", transaction[p:p+45])
	*/
}
/*
func DecodeTransaction2(transaction string){
	//fmt.Println(transaction[:42])
	//fmt.Println(transaction[42:74])
	//fmt.Println(transaction[74:106])
	//fmt.Println(transaction[106:138])
	x , _ :=strconv.Atoi(transaction[138:139])
	for i:= 0; i< x; i ++{
		fmt.Println(transaction[139:147])
		fmt.Println(transaction[147:179])
	}
	inputlength_int,_ := strconv.Atoi(transaction[179:182])
	//fmt.Println(transaction[182:182+inputlength_int])
	//fmt.Println(transaction[182+inputlength_int:182+inputlength_int+128])
	//fmt.Println(transaction[182+inputlength_int+128:182+inputlength_int+128+64])
}
*/
func StringTransactionEncode(feeString string,times int)(string){
        feeStringLen := len(feeString)
        for i :=0; i<(times-feeStringLen);i++{
                feeString = "0"+feeString;
        }
        return feeString
}

func IntTransactionEncode(fee int, times int)(string){
        feeString := strconv.Itoa(fee)
        feeStringLen := len(feeString)
        for i :=0; i<(times-feeStringLen);i++{
                feeString = "0"+feeString;
        }
        return feeString
}

func BigIntTransactionEncode(fee big.Int, times int)(string){
        feeString := fee.String()
        feeStringLen := len(feeString)
	for i :=0; i<(times-feeStringLen);i++{
		feeString = "0"+feeString;
	}
	return feeString
}

func VerifyTransactionSign(Transaction types.TransactionJson)( bool){
	pub,_ := hex.DecodeString(Transaction.PublicKey)
	msg := crypto.Keccak256([]byte(EncodeForSign(Transaction)))
	sign,_:= hex.DecodeString(Transaction.Sign)
	re := eddsa.EddsaVerify(pub,
		msg,
		sign,
	)
	return re
}

func VerifyTransactionBalanceAndNonce(core_arg types.CoreStruct ,Transaction types.TransactionJson)(bool, string){
        address := /*"gx"+*/crypto.KeyToAddress_hex(Transaction.PublicKey)
	fmt.Println(Transaction.PublicKey)
	fmt.Println(address)
	fromAccountInfo := db.AccountHexGet(core_arg.Db, address)
	//feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Sue.FeeAddress)
	fmt.Println(fromAccountInfo)
        fmt.Println("testn:", Transaction.Nonce)
        fmt.Println("testn:", fromAccountInfo.Nonce)

	if(Transaction.Nonce != fromAccountInfo.Nonce){
		return false, "nonce error"
	}

        feeAccountInfo := db.AccountHexGet(core_arg.Db, params.Chain().Version.Sue.FeeAddress)

	feeResult, fromAccountInfo, feeAccountInfo := VerifyFee(Transaction, fromAccountInfo, feeAccountInfo)

        if(feeResult != true){
                return false, "fee error"
        }

        if(Transaction.Nonce != fromAccountInfo.Nonce){
                return false, "nonce error"
        }
	
	toAccountInfo := db.AccountHexGet(core_arg.Db, Transaction.To)
	//ToAddress_s :=Transaction.To
	fmt.Println("from:",fromAccountInfo)
	fmt.Println("to:",toAccountInfo)
	//if(fromAccountInfo.Address != toAccountInfo.Address){
	if (fromAccountInfo.Address == toAccountInfo.Address){
		return false, "same address"
	}
	balanceResult, fromAccountInfo, toAccountInfo := VerifyBalance(Transaction, fromAccountInfo, toAccountInfo)
	//}
	//fmt.Println("test1",fromAccountInfo)
	if(balanceResult != true){
		return false, "balance error"
	}

	fromAccountInfo.Nonce = fromAccountInfo.Nonce+1
	//toAccountInfo.Nonce = toAccountInfo.Nonce+1
	//compare balance
	db.AccountHexPut(core_arg.Db, address, fromAccountInfo)
	db.AccountHexPut(core_arg.Db, Transaction.To, toAccountInfo)

        return true, "success"
}

func VerifyFee(transaction types.TransactionJson, fromAccount types.AccountData, feeAccount types.AccountData)(bool, types.AccountData, types.AccountData){
	for u:= 0; u < len(fromAccount.Balance); u++{
		if( fromAccount.Balance[u].Token == params.Chain().Version.Sue.FeeToken){
			if(transaction.Fee.Cmp(&fromAccount.Balance[u].Balance)>=0){
			/*fromAccount.Balance[u].Balance =*/ //fromAccount.Balance[u].Balance.Sub(&fromAccount.Balance[u].Balance, &transaction.Fee)
				return false, fromAccount, feeAccount
			}else{
				fromAccount.Balance[u].Balance.Sub(&fromAccount.Balance[u].Balance, &transaction.Fee)
				feeAccount.Balance[u].Balance.Add(&feeAccount.Balance[u].Balance, &transaction.Fee)
			}
		}
	}
	return true, fromAccount, feeAccount
}

func VerifyBalance(transaction types.TransactionJson, fromAccount types.AccountData, toAccount types.AccountData)(bool, types.AccountData, types.AccountData){
	for i:=0; i< len(transaction.Out); i++ {
		for u:= 0; u < len(fromAccount.Balance); u++{
			if(transaction.Out[i].Token[1:] == fromAccount.Balance[u].Token){
				if(transaction.Out[i].Balance.Cmp(&fromAccount.Balance[u].Balance)>=0){
					return false, fromAccount, toAccount
				}else {
					fromAccount.Balance[u].Balance.Sub(&fromAccount.Balance[u].Balance,&transaction.Out[i].Balance)
					toAccount.Balance[u].Balance.Add(&toAccount.Balance[u].Balance,&transaction.Out[i].Balance)
				}
			}
		}
	}
	return true, fromAccount, toAccount
}

func ExpTransaction()(types.TransactionJson){
        re := NewTransaction(
                //"gx5ee464a101d58877f00957eff452c148e7f75834",
		"264411884d6d2aca8ca2d2a77c9dc95ffdcee529",
		"deh",
                *big.NewInt(1000),
                0,
                *big.NewInt(0),
                "def",
                "none",
        )
        a,_ := hex.DecodeString("ab70ef5f36dbfd9e403ed4ffd5b1c51dc7ce761ee21c8dc72570c6d73bb9412b0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba")
        result := SignTransaction(a,re)
        re2 := EncodeTransaction(result)
        bb := DecodeTransaction(re2)
	return bb
}

/*
func main(){
        b := ExpTransaction()
        fmt.Println(b)
        fmt.Println(VerifyTransactionSign(b))
        core_arg := &types.CoreStruct{}
        core_arg.Db = db.OpenDB("../dbdata")

        //初始化金額
        account_tmp := account.Account_exp()
        db.AccountHexPut(core_arg.Db, account_tmp.Address, account_tmp)
        fmt.Println(db.AccountHexGet(core_arg.Db, account_tmp.Address))
        pp := ExpTransaction()
        fmt.Println(VerifyTransactionSign(pp))
        fmt.Println("sss:", pp.Out[0].Token)
        fmt.Println(crypto.KeyToAddress_hex(pp.PublicKey))
        a1 := db.AccountHexGet(core_arg.Db, account_tmp.Address)
        a2 := db.AccountHexGet(core_arg.Db, account_tmp.Address)
        r1,r2,r3 := VerifyBalance( pp, a1, a2)
        fmt.Println(r1)
        fmt.Println(r2, r2.Balance[0].Token)
        fmt.Println(r3, r3.Balance[0].Token)
        //fmt.Println(r4)
        //fmt.Println(r5)
        if(r2.Balance[0].Token == r3.Balance[0].Token){
                fmt.Println("yyy")
        }

}
*/

