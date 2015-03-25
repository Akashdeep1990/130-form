
package main

import "fmt"

type Stack struct {
	maxsize int
	con[]int
	head int
		}
// by value method deleating from a slice
func pop(a *Stack) {
 if a.head == 0 {
        fmt.Println("STACK EMPTY")
    } else {a.con = a.con[:len(a.con)-1]
	fmt.Println("by reference function call",a.con)
	a.head--
    }
}
// by reference method
func pops(a Stack) {
 if a.head == 0 {
        fmt.Println("STACK EMPTY")
    } else {
	fmt.Println("by value function call",a.con)
	a.head--
   }
}

func put(a *Stack,k int){
 if a.head == a.maxsize {
        fmt.Println("STACK FULL")
    } else {
 	a.head++
        a.con=append(a.con,k)
fmt.Println("adding to slice ",a.con)
    }
}
//appding to slice part
var a = []int{3,4}
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
put(&s,6)
pops(s)
pops(s)

}
