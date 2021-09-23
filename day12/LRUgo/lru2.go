package main


type LruCache struct {
	//现有元素数量
	Size int
	//lru容量
	Capacity int
	//map存储链表节点，查询速度O(1)
	HashMap map[Value]*Node
	//哨兵头尾，方便节点插入初始位置
	Dummy,Tail *Node
}

// InitNode 初始化新节点
func InitNode(key,value Value)*Node  {
	return &Node{
		Key: key,
		Value: value,
	}
}

// InitLruCache 初始化Lru
func InitLruCache(capacity int) LruCache  {
	l:=LruCache{
		Capacity: capacity,
		HashMap: make(map[Value]*Node,capacity),
		Dummy: InitNode(0,0),
		Tail: InitNode(0,0),
	}
	//头尾连接
	l.Dummy.Next=l.Tail
	l.Tail.Pre = l.Dummy
	return l
}

//两个基础函数，增加到头，移除节点，方便使用

// AddToHead 增加到头部
func (m*LruCache)AddToHead(n*Node)  {
	//双向链表，两边都增加
	n.Next = m.Dummy.Next
	n.Pre = m.Dummy
	m.Dummy.Next.Pre = n
	m.Dummy.Next = n
}

// Remove 移除节点
func (m*LruCache)Remove(n*Node)  {
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
}

// MoveHead 节点移动，删除节点，移动到头部
func (m*LruCache)MoveHead(n*Node)  {
	m.MoveHead(n)
	m.AddToHead(n)
}

// DelTail 超过容量，删除最后一个
//其实是删除tail的pre,返回node，方便移除map中的数据
func (m*LruCache)DelTail()*Node  {
	n:=m.Tail.Pre
	m.Remove(n)
	return n
}

//主要的get和put函数

// Get key获取value 没有返回-1
func (m*LruCache)Get(key Value) Value {
	if _,ok :=m.HashMap[key];!ok{
		return -1
	}
	//使用到，移到头，返回value
	n:=m.HashMap[key]
	m.MoveHead(n)
	return n.Value
}

// Put put  存在直接改值，不存在，加入判断是否超出长度,超出删尾部
func (m*LruCache)Put(key,value Value) {
	if _,ok:=m.HashMap[key];ok{
		n:=m.HashMap[key]
		n.Value = value
		m.MoveHead(n)
	}else{
		if m.Size+1>m.Capacity{
			node:=m.DelTail()
			//上面是移除链表，记得移除map中的数据
			delete(m.HashMap,node.Key)
			m.Size--
		}
		n:=InitNode(key,value)
		m.HashMap[key] = n
		//map添加元素，但是链表未增加节点，所以调用AddToHead，添加节点到链表即可，不要使用MoveHead
		m.AddToHead(n)
		m.Size++
	}
}