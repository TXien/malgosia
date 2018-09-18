package main

import (
	"fmt"
	"time"
)

func DeleteElement(a []string, i int)([]string){
        //i := 2
        a[i] = a[len(a)-1]
        a[len(a)-1] = ""
        a = a[:len(a)-1]
        return a
}

func fast(a []string)([]string){
	//a = []string{"A", "B", "C", "D", "E"}
	i := 2
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	//fmt.Println(a)
	return a
}

func slow(a []string){
	i := 2
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	fmt.Println(a)
}

func main(){
	t1 := time.Now()
	a := []string{"A", "B", "C", "D", "E"}
	//for i :=0; i<1000000; i++{
	/*
	b := fast(a)
	c := fast(b)
	d := fast(c)
	fmt.Println(d)
	*/
		//slow(a)
	//}
	for i:=0; i< len(a); i++ {
		if(a[i]=="A"){
			a = DeleteElement(a,i)
			//fmt.Println(a)
		}
	}
	//a = DeleteElement(a,0)
	fmt.Println(a)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
