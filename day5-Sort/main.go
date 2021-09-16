package main

import (
	"day5-Sort/sort"
	"fmt"
	"math/rand"
	"time"
)

const NUM = 10000000

func main()  {

	/*nums :=make([]int,NUM)
	nums1 :=make([]int,NUM)
	nums2 :=make([]int,NUM)*/
	nums3 :=make([]int,NUM)
	nums4 := make([]int, NUM)

	// 随机种子
	rand.Seed(time.Now().Unix())
	// 从 [0, 100) 范围的伪随机数为待排序nums初始化

	for i := 0; i < NUM; i++ {
		/*nums[i] = rand.Intn(100)
		nums1[i] = rand.Intn(100)
		nums2[i] = rand.Intn(100)*/
		nums3[i] = rand.Intn(100)
		nums4[i] = rand.Intn(100)
	}

	/*fmt.Println("随机数列1",nums)
	fmt.Println("随机数列2",nums1)
	fmt.Println("随机数列3",nums2)
	fmt.Println("随机数列4",nums3)
	fmt.Println("随机数列5",nums4)*/

	/*fmt.Println("======BubbleSort=====")
	start := time.Now()
	//some func or operation,This is BubbleSport operation
	nums = sort.BubbleSort(nums)
	cost := time.Since(start)
	fmt.Printf("useTime=[%v]\n",cost)
	fmt.Println()

	fmt.Println("======InsertionSort=====")
	start = time.Now()
	//some func or operation,This is InsertionSort operation
	nums1 = sort.InsertionSort(nums1)
	cost = time.Since(start)
	fmt.Printf("useTime=[%s]\n",cost)
	fmt.Println()

	fmt.Println("======SelectionSort=====")
	start = time.Now()
	//some func or operation,This is SelectionSort operation
	nums2 = sort.SelectionSort(nums2)
	cost = time.Since(start)
	fmt.Printf("useTime=[%s]\n",cost)
	fmt.Println()
*/
	fmt.Println("======MergeSort=====")
	start := time.Now()
	//some func or operation,This is MergeSort operation
	nums3 = sort.MergeSort(nums3)
	cost := time.Since(start)
	fmt.Printf("useTime=[%s]\n",cost)
	fmt.Println()

	fmt.Println("======QuickSort=====")
	start = time.Now()
	//some func or operation,This is QuickSort operation
	nums4 = sort.QuickSort(nums4,0,len(nums4)-1)
	cost = time.Since(start)
	fmt.Printf("useTime=[%s]\n",cost)
	fmt.Println()

}
