package main


import "C"
import (
	//"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	secp256k1 "github.com/bitherhq/go-bither/crypto/secp256k1"
	secp256k1go "github.com/haltingstate/secp256k1-go"
	"time"
	//"unsafe"
	//"io"
	//"testing"
)

func generateKeyPair() (pubkey, privkey []byte) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), key.X, key.Y)

	privkey = make([]byte, 32)
	blob := key.D.Bytes()
	copy(privkey[32-len(blob):], blob)
	return pubkey, privkey
}

func generateKeyPair_string() (string,string) {
        key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
        if err != nil {
                panic(err)
        }
	pubkey := elliptic.Marshal(secp256k1.S256(), key.X, key.Y)
	privkey := make([]byte, 32)
        blob := key.D.Bytes()
        copy(privkey[32-len(blob):], blob)
        return hex.EncodeToString(pubkey), hex.EncodeToString(privkey)
}

func sign(msg []byte,seckey[]byte) ([]byte){
	sig, err := secp256k1.Sign(msg, seckey)
	if err != nil {
		fmt.Println("signature error: %s", err)
	}
	return sig
}

func sign_string(message string , key string) (string){
	msg := []byte(message)
	seckey, _ := hex.DecodeString(key)
	sig, err := secp256k1.Sign(msg, seckey)
        if err != nil {
                fmt.Println("signature error: %s", err)
        }
        return hex.EncodeToString(sig)
}

func signatureVerify_string(msg string ,sig []byte,pub []byte) int{
	//pubkey,_:= hex.DecodeString(pub)
	return secp256k1go.VerifySignature([]byte(msg), []byte(sig),pub)
}

func main(){ 
	fmt.Println("123")

	x , y := generateKeyPair_string()
	fmt.Println("pub:",x,"priv:",y)
	_, b := generateKeyPair()
	now1 := time.Now()
	ss := sign([]byte("11111111111111111111111111111111"), b)
	
	now2 := time.Now()
	fmt.Println("sign_time",now2.Sub(now1))
	fmt.Println("sign:",sign_string("11111111111111111111111111111111", y))
	//now3 := time.Now()
	f,_ := hex.DecodeString(`03fe43d0c2c3daab30f9472beb5b767be020b81c7cc940ed7a7e910f0c1d9feef1`)
	now3 := time.Now()
	for i:=0; i <= 100000; i++{
		fmt.Println(signatureVerify_string("11111111111111111111111111111111",ss ,f ))
	}
	fmt.Println("sign_string_time",time.Now().Sub(now3))

}




