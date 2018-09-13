package db

import (
        //"encoding/hex"
        "fmt"
        "reflect"
        "testing"
	"time"
)

func TestAccount(t *testing.T){
        check := func(f string, got, want interface{}) {
                if !reflect.DeepEqual(got, want) {
                        t.Errorf("%s mismatch: got %v, want %v", f, got, want)
                }
        }
        db := OpenDB("../dbdata")
	t1 := time.Now()
	for i :=1; i <=10000; i++{
		db.Put( []byte("keyddd"), []byte("11"),nil)
	}
	fmt.Println("put10000:",time.Now().Sub(t1))
	t2 := time.Now()
	for i2 :=1; i2 <=10000; i2++{
		f,err := db.Get([]byte("keyddd"),nil)
		if(err != nil){
			fmt.Println(f)
		}
	}
	fmt.Println("get10000:",time.Now().Sub(t2))
	f,_ := db.Get([]byte("keydd1"),nil)
	d := string(f)
        fmt.Println(f)
	check("get:","11",d)
}
