package crypto

import (
	//"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func KeyToAddress(key []byte)([]byte){
	return Keccak256(key)[12:]
}

func KeyToAddress_hex(key string)(string){
	c,_ := hex.DecodeString(key)
	re := Keccak256(c)[12:]
	return hex.EncodeToString(re)//Keccak256(key)[12:]
}

func Keccak256(msg []byte)([]byte){
	return crypto.Keccak256(msg)
}


