package db
import (
	"encoding/hex"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"fmt"
	//"lurcury/core/block"
	//"lurcury/core/transaction"
	"lurcury/types"
)


//type mleveldb leveldb.DB

func OpenDB(path string)(*leveldb.DB){
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func /*(db *mleveldb)*/Get(db *leveldb.DB, key []byte)([]byte){
	data, _ := db.Get(key,nil)
	return data
}


func /*(db *mleveldb)*/Put(db *leveldb.DB, key []byte,data []byte){
	db.Put(key, data, nil)
}

func /*(db *mleveldb)*/Delete(db *leveldb.DB, key []byte){
        db.Delete(key,nil)//, nil)
}

func /*(db *mleveldb)*/AccountHexPut(db *leveldb.DB, keys string,data types.AccountData){
        key,_ := hex.DecodeString(keys)
        data_byte, _ := json.Marshal(data)
        db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/AccountHexGet(db *leveldb.DB, keys string)(types.AccountData){
        key,_ := hex.DecodeString(keys)
        data, _ := db.Get(key, nil)
        inter := types.AccountData{}
        json.Unmarshal(data, &inter)
        return inter
}

func /*(db *mleveldb)*/BlockHexPut(db *leveldb.DB, keys string, data types.BlockJson){
	key,_ := hex.DecodeString(keys)
	data_byte, _ := json.Marshal(data)
        db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/BlockHexGet(db *leveldb.DB, keys string)(types.BlockJson){
	key,_ := hex.DecodeString(keys)
	data, _ := db.Get(key, nil)
        inter := types.BlockJson{}
        json.Unmarshal(data, &inter)
        return inter
}

func /*(db *mleveldb)*/BlockPut(db *leveldb.DB, key []byte,data types.BlockJson){
        data_byte, _ := json.Marshal(data)
        db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/BlockGet(db *leveldb.DB, key []byte)(types.BlockJson){
        data, _ := db.Get(key, nil)
        inter := types.BlockJson{}
        json.Unmarshal(data, &inter)
        return inter
}

func /*(db *mleveldb)*/TransactionPut(db *leveldb.DB, key []byte,data types.TransactionJson){
	data_byte, _ := json.Marshal(data)
	db.Put(key, data_byte, nil)
}

func /*(db *mleveldb)*/TransactionGet(db *leveldb.DB, key []byte)(types.TransactionJson){
	data, _ := db.Get(key, nil)
	inter := types.TransactionJson{}
	json.Unmarshal(data, &inter)
	return inter
}
