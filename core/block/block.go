package block

import (
	"lurcury/crypto/eddsa"
	"lurcury/db"
	"lurcury/params"
	"lurcury/core/transaction"
	"lurcury/types"
	"encoding/hex"
	//"encoding/json"
	"fmt"
	"strconv"
	"time"
	"lurcury/crypto"
)
func CreateBlockPOA(core_arg *types.CoreStruct, parentBlock types.BlockJson, key string)(types.BlockJson){
	newBlock := CreateNewBlock(parentBlock)
	newBlock.Transaction = core_arg.PendingTransaction
	//fmt.Println("tssst:",newBlock.Transaction)
	for i := 0; i < len(core_arg.PendingTransaction); i++{
		result , err:= transaction.VerifyTransactionBalanceAndNonce(*core_arg, core_arg.PendingTransaction[i])
		fmt.Println(err)
		if(result == true){
			newBlock.Transaction = append(newBlock.Transaction, core_arg.PendingTransaction[i])
		}
	}
	//fmt.Println("newBlock:",newBlock)
	//newBlock.Timestamp = 1536924111618191122
	//fmt.Println("testtt:",newBlock.Transaction)
	encodeBlock := BlockEncode(newBlock)
	//fmt.Println("encodeBlock:",encodeBlock)
	signBlock :=  BlockSign(key, encodeBlock)
	//fmt.Println("BlockSign.Hash:",signBlock.Hash)
	db.BlockHexPut(core_arg.Db, signBlock.Hash, signBlock)
        return signBlock
}


func CreateNewBlock(parentBlock types.BlockJson)(types.BlockJson){
        s := types.BlockJson{
                Version: "sue",
                BlockNumber: parentBlock.BlockNumber+1,
                ParentHash: parentBlock.Hash,
                Nonce: 0,
                Timestamp: time.Now().UnixNano(),
                ExtraData: "ka",
                //Hash: hash,
        }
        return s
}

func NowBlock(block types.BlockJson)(types.BlockJson){
        s := types.BlockJson{
                Version: "sue",
                BlockNumber: block.BlockNumber+1,
                ParentHash: block.ParentHash,
                Nonce: 0,
                Timestamp: time.Now().UnixNano(),
                ExtraData: "ka",
                //Hash: hash,
        }
        return s
}

func NewBlock(version string, blockNumber int, parentHash string, nonce int, time int64, extraData string)(types.BlockJson){
	s := types.BlockJson{
		Version: version,
		BlockNumber: blockNumber,
		ParentHash: parentHash,
		Nonce: nonce,
		Timestamp: time,
		ExtraData: extraData,
		//Hash: hash,
	}
	//fmt.Println(s)
	return /*BlockEncode(*/s//)
}

func BlockEncode(block types.BlockJson)(types.BlockJson){
	re := block.Version
	re = re + string(block.BlockNumber)
	re = re + block.ParentHash
	re = re + string(block.Nonce)
	re = re + strconv.FormatInt(block.Timestamp,10)
	re = re + block.ExtraData
	transaction_length := len(block.Transaction)
	//fmt.Println(transaction_length)
	for i:=0; i < transaction_length; i++{
		re = re + block.Transaction[i].Tx
	}

	block.Hash = hex.EncodeToString(crypto.Keccak256([]byte(re)))
	return block
}

func BlockEncode_DB(block types.BlockJson)(types.BlockJson){
        re := block.Version
        re = re + string(block.BlockNumber)
        re = re + block.ParentHash
        re = re + string(block.Nonce)
        re = re + strconv.FormatInt(block.Timestamp,10)
        re = re + block.ExtraData
        transaction_length := len(block.Transaction)
        //fmt.Println(transaction_length)
        for i:=0; i < transaction_length; i++{
                re = re + block.Transaction[i].Tx
        }

        block.Hash = hex.EncodeToString(crypto.Keccak256([]byte(re)))
        return block
}

func BlockSign(priv string, block types.BlockJson)(types.BlockJson){
	pri, _ := hex.DecodeString(priv)
	hash, _ := hex.DecodeString(block.Hash)
	re := eddsa.EddsaSign(pri, hash)
	block.Verifier = append(block.Verifier, types.VerifierJson{Sign:hex.EncodeToString(re),Verifier:hex.EncodeToString(pri[32:]),N:0})
	return block
	//BlockJson.Verifier[0]. := eddsa.EddsaSign(pri, hash)
}

func AppendTransaction(trans types.TransactionJson, block types.BlockJson)(types.BlockJson){
	block.Transaction = append(block.Transaction, trans)
	return block
}
/*
func POA()(bool){
	return true
}
*/
func CheckRecentBlock(core_arg *types.CoreStruct)(string){
	b := params.Chain()
	c := db.HexKeyGet(core_arg.Db, b.Hash)
	return c
}

func ExpBlock()(block types.BlockJson){
        b:=NewBlock("sue",
        0,
        "fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
        0,
        0,
        "mogotisa",
        //"fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
        )
        //fmt.Println(b)
        return b
}

/*
func main(){
	t1 := time.Now()
	b:=NewBlock("sue",
	0,
	"fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
	0,
	t1.UnixNano(),
	"mogotisa",
	//"fea4910f5d3e2d3af187cec5b8d8b1cfe99a9f5545ba50495bd42f4bae234b3a",
	)
	d := BlockEncode(b)
	fmt.Println(d.Hash)
	pri,_ := hex.DecodeString("219a634773d787cfbaf1e5c915d56b14be2a3695ed8e46bbeb01573bf211d0ef8773580834eb42a2f2ee856b029a88dfee639e27f08b1e0235f8eb04eecf4089")
	fmt.Println(BlockSign(pri, d))
	fmt.Println(POA())
	a := transaction.TransactionJson{Tx:"123"}
	dd := AppendTransaction(a,d)
	zz, _ := json.Marshal(dd)
	fmt.Println(zz)
	hh := BlockJson{}
	json.Unmarshal(zz, &hh)
	fmt.Println(hh.Version)
}

*/
