package main

import "fmt"

func main(){
	l,r,a := 0,0,0
	fmt.Scanf("%d",&l)
	fmt.Scanf("%d",&r)
	fmt.Scanf("%d",&a)
	dist := r -l
	k := dist/a
	//res := k*(dist+1)-((k+1)*k/2)*2
	res := k*(dist + 1)-((k+1)*k/2*a)
	fmt.Println(res)
}