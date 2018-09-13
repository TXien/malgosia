package route

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	//"lurcury/core/transaction"
	"lurcury/types"
)

type TestBodys_Struct struct {
        Id   int64  `json:"id"`
        Name string `json:"name"`
}

func Router(coreStruct *types.CoreStruct){
        Broadcast := func(res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
		b, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		var msg types.TransactionJson
		json.Unmarshal(b, &msg)//coreStruct.PendingTransaction)
		coreStruct.PendingTransaction = append(coreStruct.PendingTransaction,msg)
		//fmt.Println(coreStruct.PendingTransaction)
		res.Header().Set("content-type", "application/json")
		res.Write([]byte("suc"))
        }

        http.HandleFunc("/broadcast", Broadcast)

        PendingTransaction := func(res http.ResponseWriter, req *http.Request){
                res.Header().Add("Access-Control-Allow-Origin","*")
                //b, _ := ioutil.ReadAll(req.Body)
                //defer req.Body.Close()
                //var msg types.TransactionJson
                //json.Unmarshal(b, &msg)
                res.Header().Set("content-type", "application/json")
                res.Write([]byte(""))
        }

	http.HandleFunc("/broadcast", Broadcast)
}
