package lru

import l "LRU/dll"

type LRUCache struct {
    capacity int
    currSize int
    mapping map[int]*l.Entry
    Li l.List
}


func Constructor(capacity int) LRUCache {
    head, tail := &l.Node{Prev: nil}, &l.Node{Next: nil}
    head.Next, tail.Prev = tail, head

    return LRUCache{
        capacity: capacity,
        currSize: 0,
        mapping: make(map[int]*l.Entry),
        Li: l.List{
            Head: head,
            Tail: tail,
        },
    }
}

func (this *LRUCache) Get(key int) int {
    if e, ok := this.mapping[key]; ok {
        l.Movenode(this.Li, e.N)
        return e.Value
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if e, ok := this.mapping[key]; ok {
        l.Movenode(this.Li, e.N)
        this.mapping[key].Value = value
    } else {
        if (this.currSize == this.capacity) {
            delete(this.mapping, this.Li.Tail.Prev.Key)
            l.Removenode(this.Li)
            this.currSize--
        }

        n := &l.Node{Key: key}
        l.Insertnode(this.Li, n)

        this.mapping[key] = &l.Entry{
            Value: value,
            N: n,
        }

        this.currSize++
    }
}
