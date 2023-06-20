/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
package leetcodego

import "container/list"

type entry struct {
    key, value, fre int
}

type LFUCache struct {
    size, cap, minFre int
    list              map[int]*list.List
    cache             map[int]*list.Element
}

func Constructor1(capacity int) LFUCache {
    l := LFUCache{
        cap:   capacity,
        cache: make(map[int]*list.Element),
        list:  make(map[int]*list.List),
    }
    return l
}

func (this *LFUCache) Get(key int) int {
    if this.cap == 0 {
        return -1
    }
    if e, ok := this.cache[key]; ok {
        Value := this.list[e.Value.(*entry).fre].Remove(e)
        // delete
        if this.minFre == Value.(*entry).fre && this.list[Value.(*entry).fre].Len() == 0 {
            this.minFre++
			delete(this.list, Value.(*entry).fre)
        }
        Value.(*entry).fre++
        if _, ok := this.list[Value.(*entry).fre]; !ok {
            this.list[Value.(*entry).fre] = list.New()
        }
        // add
        e = this.list[Value.(*entry).fre].PushFront(Value)
        this.cache[Value.(*entry).key] = e
        return e.Value.(*entry).value
    }
    return -1
}


func (this *LFUCache) Put(key int, value int) {
    if this.cap == 0 {
        return
    }
    if e, ok := this.cache[key]; ok {
        Value := this.list[e.Value.(*entry).fre].Remove(e)
        // delete
        if this.minFre == Value.(*entry).fre && this.list[Value.(*entry).fre].Len() == 0 {
			delete(this.list, Value.(*entry).fre)
            this.minFre++
        }
        Value.(*entry).fre++
        Value.(*entry).value = value
        if _, ok := this.list[Value.(*entry).fre]; !ok {
            this.list[Value.(*entry).fre] = list.New()
        }
        // add
        e = this.list[Value.(*entry).fre].PushFront(Value)
        this.cache[Value.(*entry).key] = e
    } else {
        // delete if fully
        if this.size == this.cap {
            e := this.list[this.minFre].Back()
            Value := this.list[this.minFre].Remove(e)
            delete(this.cache, Value.(*entry).key)
            this.size--
        }

        this.minFre = 1
        newEntry := &entry{key: key, value: value, fre: 1}
        if _, ok := this.list[this.minFre]; !ok {
            this.list[this.minFre] = list.New()
        }
        e := this.list[this.minFre].PushFront(newEntry)
        this.cache[newEntry.key] = e
        this.size++
    }
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
