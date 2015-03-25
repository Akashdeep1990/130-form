//3
package main

import "fmt"



func sum(i int,j int)int{
	return i+j
}
func sub(i int,j int)int{
	return i-j
}
func mul(i int,j int)int{
	return i*j
}
func div(i int,j int)int{
	return i/j
}

func action(a string,b int, c int)int{
	switch a{
	case "SUM":return sum(b,c)
	case "SUB":return sub(b,c)
	case "MUL":return mul(b,c)
	case "DIV":return div(b,c)

}
return 0
}

func increment(j int)func()int{
	i:=j
return func() (ret int) {
		ret = i
		if i%2==0{
		i+=1}else{i+=2}
		d:=false
		for d==false{
j:=i-1
			for j>1	{
			if i%j ==0{
				d=true
				break}
			j-=2
				}
		if d==true{
			i+=2
				}else{
		break	
		
}}
		return
	
}}

func main(){
fmt.Println("Applying Arithmetic functions on 3 ,2")
b:="SUM"
	a:=action(b,3,2)
	fmt.Println("SUM: ",a)
b="SUB"
	a=action(b,3,2)
	fmt.Println("SUB: ",a)
b="MUL"
	a=action(b,3,2)
	fmt.Println("MUL: ",a)
b="DIV"
	a=action(b,3,2)
	fmt.Println("DIV: ",a)
	isnext:=increment(2)
fmt.Println("SEQUENCIENCIALLY PRINTING NEXT PRIME BY FUNCTION CALL")
fmt.Println(isnext())
fmt.Println(isnext())
fmt.Println(isnext())
fmt.Println(isnext())
fmt.Println(isnext())
fmt.Println(isnext())

}
