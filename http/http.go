package http

import (
	//"lurcury/core/block"
	//"lurcury/core"
	"encoding/json"  
	"fmt"
	"io/ioutil"
	"net/http"
	"lurcury/http/route"
	"lurcury/types"
	"time"

)

func readFile(filename string) (map[string]string, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
        return nil, err
    }
    var j = map[string]string{}
    if err := json.Unmarshal(bytes, &j); err != nil {
        fmt.Println("Unmarshal: ", err.Error())
        return nil, err
    }
    return j, nil
}

func httpSet (coreStruct *types.CoreStruct){
	config,err :=readFile("config.json")
	if(err!=nil){
		fmt.Println("error:", err)
	}
	route.Router(coreStruct)
	fmt.Println("connect port"+config["port"])
	err2:= http.ListenAndServe(config["port"], nil)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

}

func httpSet2 (coreStruct types.CoreStruct){
	fmt.Println("connect port"+":14456") 
	route.Test(coreStruct)
        err2:= http.ListenAndServe(":14456", nil)
        if err2 != nil {
                fmt.Println("error:", err2)
        }

}


func Server(coreStruct *types.CoreStruct) { 
	httpSet(coreStruct)
	//httpSet2(coreStruct)
	time.Sleep(100 * time.Second)
}
