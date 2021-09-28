package main

import "fmt"

func count(n int) int  {
	if n<1 {
		return 0
	}
	count := 0
	base := 1
	round := n
	for round>0 {
		weight := round%10
		round/=10
		count+=round*base
		if weight==1 {
			count+=(n%base)+1
		}else if weight > 1 {
			count+=base
		}
		base*=10
	}
	return count
}

func main()  {
	t1:=count(534)
	t2:=count(530)
	t3:=count(504)
	t4:=count(514)
	t5:=count(10)
	fmt.Println(t1,t2,t3,t4,t5)
}
