/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
package leetcodego

type LRUCache struct {
	size, capacity int
	cache          map[int]*LRUNode
	head, tail     *LRUNode // 头尾节点，dummy node
}

type LRUNode struct {
	key, value int
	pre, next  *LRUNode
}

func initLRUNode(key, value int) *LRUNode {
	return &LRUNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		head:     initLRUNode(0, 0),
		tail:     initLRUNode(0, 0),
		cache:    map[int]*LRUNode{},
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1 // 如果不存在，直接返回-1
	}
	// 如果存在，返回对应的value，并且将节点放在最前面
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.removeNode(node) // 先将node从原来的位置移出
	this.addToHead(node) // 然后将node移动到链表最前面
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
	node.next = nil
	node.pre = nil
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		// 如果key不存在，构造新的节点放入链表和缓存cache，并判断链表长度是否超过capacity
		newNode := initLRUNode(key, value)
		this.addToHead(newNode)
		this.cache[key] = newNode
		this.size++
		if this.size > this.capacity {
			// 如果链表长度达到上限，移出链表末尾的元素，并从cache中清除
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
		return
	}
	this.moveToHead(node)
	node.value = value
}

func (this *LRUCache) removeTail() *LRUNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
