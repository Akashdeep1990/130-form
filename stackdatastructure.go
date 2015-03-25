
package main

import "fmt"

type Stack struct {
	maxsize int
	con[10]int
	head int
		}
func pop(a *Stack) {
 if a.head == 0 {
        fmt.Println("STACK EMPTY")
    } else {
	fmt.Println("ELEMENT: ",a.con[a.head])
	a.head--
    }
}

func put(a *Stack,k int){
 if a.head == a.maxsize {
        fmt.Println("STACK FULL")
    } else {
 	a.head++
        a.con[a.head]=k
    }
}
func main(){
	s:= Stack{maxsize:6,head:0}
	put(&s,1)

put(&s,2)
put(&s,3)
put(&s,4)
put(&s,5)
put(&s,6)
i:=0
 for i<=6{
 pop(&s)
i++
} 

}
