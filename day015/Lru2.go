package main

/*
* Lru second design
*@param [][]int 整性二维数组
*@param k int the length of the cache
*@ return []int 整型一维数组
 */

type LRUMap struct {
	head *LRUNode
	tail *LRUNode
	m map[int]*LRUNode
	k int
}

type LRUNode struct {
	Key int
	Val int
	Next *LRUNode
	Pre *LRUNode
}

func getLRUMap(k int) *LRUMap  {
	return &LRUMap{
		m: make(map[int]*LRUNode,k),
		k:k,
	}
}

func LRU( operators [][]int ,  k int ) []int {
	res := make([]int, 0, k)
	// write code here
	lru := getLRUMap(k)
	for _, operator := range operators {
		if getOpt(operator) == 1 {
			lru.addLRU(operator[1], operator[2])
		} else if getOpt(operator) == 2 {
			a := lru.getLRU(operator[1])
			res =append(res, a)
		}
	}
	return res
}

func (p *LRUMap) addLRU(key, val int) {
	if cur, ok := p.m[key]; ok {
		cur.Val = val
		if cur == p.head {
			return
		} else {
			p.getSwitchLRU(cur)
		}
	} else {
		cur = &LRUNode{
			Val: val,
			Key: key,
		}
		if p.head == nil {
			p.head = cur
		} else {
			cur.Next = p.head
			p.head.Pre = cur
			p.head = cur
		}
		if p.tail == nil {
			p.tail = cur
		}
		p.m[cur.Key] = cur
	}
	if len(p.m) > p.k {
		tailKey := p.tail.Key
		delete(p.m, tailKey)

		pre := p.tail.Pre
		pre.Next = nil
		p.tail = pre
	}
}

func (p *LRUMap) getLRU(key int) int {
	res := -1
	if cur, ok := p.m[key]; ok {
		res = cur.Val
		if  cur != p.head {
			p.getSwitchLRU(cur)
		}
	}
	return res
}

func (p *LRUMap) getSwitchLRU(cur *LRUNode) {
	next := cur.Next
	pre := cur.Pre
	cur.Next = nil
	cur.Pre = nil

	if next != nil {
		next.Pre = pre
	}
	if pre != nil {
		pre.Next = next
	}
	if cur == p.tail {
		p.tail = pre
	}

	cur.Next = p.head
	p.head.Pre = cur
	p.head = cur
}

func getOpt(s []int) int {
	if len(s) > 0 {
		return s[0]
	}
	return 0
}