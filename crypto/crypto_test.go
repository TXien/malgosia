package crypto

import (
	"encoding/hex"
	//eddsa "lurcury/crypto/eddsa"
	"fmt"
	"testing"
	"reflect"
)

func TestCrypto(t *testing.T){
        check := func(f string, got, want interface{}) {
                if !reflect.DeepEqual(got, want) {
                        t.Errorf("%s mismatch: got %v, want %v", f, got, want)
                }
        }
	address := hex.EncodeToString(KeyToAddress([]byte("9dc8a221a27d4bf0df46ba54c04e28cca51d13d10ccb1e9cb700bfa7a88a212c")))
	keccak := Keccak256([]byte("9dc8a221a27d4bf0df46ba54c04e28cca51d13d10ccb1e9cb700bfa7a88a212c"))
	x,_ := hex.DecodeString("0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba")
	fmt.Println("key:",hex.EncodeToString(KeyToAddress(x)))

	fmt.Println("hex address:",KeyToAddress_hex("0b1d7080dd923a7dfe42de42ee3e13feebd9c56f4c5cff6862e2d2890b4e1aba"))
	check("keyToAddress()", address, "3c01f961399c0c50a7ad37778eee8a20e1c1bb32")
	check("keccak256()", keccak,"7a5189acc8bb077a3e97489661da6facdbdbf805c8b69c2ece50135c2f0dbf74")
}

