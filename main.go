package main
import "github.com/syndtr/goleveldb/leveldb"
import "fmt"
import "strconv"
import "time"
func main() {
	db, _ := leveldb.OpenFile("path/to/db", nil)
	now := time.Now()
	fmt.Println(now)
	for i:=0; i <= 1000000; i++{
		//fmt.Printf(strconv.Itoa(i)+"\n")
		_ = db.Put([]byte("key"+string(i)), []byte("445566"), nil)
	}
        now = time.Now()
        fmt.Println(now)
	for i:=0; i <= 1000000 ; i++{
		data, _ := db.Get([]byte("key"+string(i)), nil)
		if(string(data)=="0"){
			fmt.Printf(string(data))
			fmt.Printf(strconv.Itoa(i)+"\n")
		}
	}
        now = time.Now()
	fmt.Println(111)
}
