
package main

import "fmt"

type Stack struct {
	maxsize int
	con[10]int
	head int
		}
func pop(a *Stack) int{
 i:=0
 if a.head == 0 {
        fmt.Println("STACK EMPTY")
    } else {
	i:=a.con[a.head]
	a.head--
	return i
    }
return i
}

func put(a *Stack,k int){
 if a.head == a.maxsize {
        fmt.Println("STACK FULL")
    } else {
 	a.head++
        a.con[a.head]=k
    }
}

func sumsubmul(a *Stack)(d int,b int,c int){
	if a.head<6{fmt.Println("NEED ATLEAST 6 ELEMENTS in STACK")}else{
	d=pop(a) + pop(a)
	b=pop(a) - pop(a)
	c=pop(a) * pop(a)
}
return
}

func main(){
	s:= Stack{maxsize:6,head:0}
	put(&s,1)

put(&s,2)
put(&s,3)
put(&s,4)
put(&s,5)
put(&s,6)
temp1,temp2,temp3:=sumsubmul(&s)
fmt.Println("SUM: ",temp1)
fmt.Println("SUBRACTION: ",temp2)
fmt.Println("MULTIPLICATION: ",temp3)


}
