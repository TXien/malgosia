package main

import "time"
import "fmt"
type buck struct{v int}

var quack = func(d *buck) {
     // do something
}

// the function we are testing:
func duck(d *buck) {
    quack(d)
    d.v = 1
}

func main(){
	b := &buck{}
	go duck(b)
	time.Sleep(1 * time.Second)
	fmt.Println(b.v)
}
