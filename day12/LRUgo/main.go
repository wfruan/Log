package main

import (
	"fmt"
)

var head *Node
var end *Node


func (n *Node)Init(key Value, value Value){
	n.Key = key
	n.Value = value
}

// SetLRUCache Init the Cache length
func SetLRUCache(size int) *LRUCache{
	l := LRUCache{
		Limit:size,
		HashMap: make(map[Value]*Node,size),
	}
	return &l
}

func (l *LRUCache)get(key Value) Value{
	if v,ok:= l.HashMap[key];ok {
		l.refreshNode(v)
		return v.Value
	}else {
		fmt.Printf("Get %v error\n",key)
		return ""
	}
}

func (l *LRUCache)put(key , value Value) {
	if v,ok := l.HashMap[key];!ok{
		if len(l.HashMap) >= l.Limit{
			oldKey := l.removeNode(head)
			delete(l.HashMap, oldKey)
		}
		node := Node{Key:key, Value:value}
		l.addNode(&node)
		l.HashMap[key] = &node
	}else {
		v.Value = value
		l.refreshNode(v)
	}
}

// 更新(1.判断 2.删除 3，插入到头)
func (l *LRUCache) refreshNode(node *Node){
	if node == end {
		return
	}
	l.removeNode(node)
	l.addNode(node)
}

// 删除（两种情况1.恰好在末尾;  2.节点在中间）
func (l *LRUCache) removeNode(node *Node) Value{
	if node == end  {
		end = end.Pre
	}else if node == head{
		head = head.Next
	}else {
		node.Pre.Next = node.Next
		node.Next.Pre = node.Pre
	}
	return node.Key
}


func (l *LRUCache) addNode(node *Node){
	if end != nil{
		end.Next = node
		node.Pre = end
		node.Next = nil
	}
	end = node
	if head == nil {
		head = node
	}
}

func main(){

	/*//具体实现在上面
	lruCache1 := SetLRUCache(4)
	lruCache1.put(1, "用户１1信息")
	lruCache1.put("002", "用户１信息1")
	lruCache1.put("003", "用户１信息")
	lruCache1.put("004", "用户１信息")
	lruCache1.put(5, "用户１信息")
	fmt.Println(lruCache1.get("008"))
	fmt.Println(lruCache1.get("002"))
	//lruCache.get("002")
	lruCache1.put("004", "用户２信息更新")
	lruCache1.put("006", "用户6信息更新")
	fmt.Println(lruCache1.get("001"))
	fmt.Println(lruCache1.get("006"))
	fmt.Print(lruCache1.HashMap)
	fmt.Println()*/


	/*//lru2测试，具体实现见lru2.go,Get Put方法均大写以区别
	lruCache2 := InitLruCache(4)
	lruCache2.Put(1,"rwf1")
	lruCache2.Put(2,"rwf2")
	lruCache2.Put(3,"rwf3")
	lruCache2.Put(4,"rwf4")
	lruCache2.Put(5,"rwf5")
	fmt.Print(lruCache2)*/

	l3 := SetLRUCache(4)
	l3.put(1,"rwf1")
	l3.put(2,"rwf2")
	l3.put(3,"rwf3")
	l3.put(4,"rwf4")
	l3.put(5,"rwf5")
	l3.put(6,"rwf6")
	fmt.Print(l3.HashMap)
	fmt.Println()



	// 了解container/list的用法
	//UseList()

}