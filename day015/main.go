package main

import "fmt"

func main()  {
	/*var A,B,C,D,E int
	fmt.Scanf("%d",&A)
	fmt.Scanf("%d",&B)
	fmt.Scanf("%d",&C)
	fmt.Scanf("%d",&D)
	fmt.Scanf("%d",&E)
	res := Temperature(A,B,C,D,E)
	fmt.Println(res)*/
	a := [][]int{

		{1, 1, 1} ,   /*  第一行索引为 0 */

		{1, 2, 2} ,   /*  第二行索引为 1 */

		{1, 3, 2},   /* 第三行索引为 2 */

		{2,1},

		{1,4,4},

		{2,2},
	}

	k :=3

	fmt.Println(LRU(a,k))
}
