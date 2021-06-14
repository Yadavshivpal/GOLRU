package main

import "fmt"

type node struct {
    key int
    next *node
    prev *node
}

type list struct {
    head *node
    tail *node
}

type entry struct {
    value int
    n *node
}

type LRUCache struct {
    capacity int
    currSize int
    record map[int]*entry
    li list
}

func moveNodeToFront(li list, n *node) {
    head := li.head

    n.next.prev = n.prev
    n.prev.next = n.next

    n.next = head.next
    head.next.prev = n

    n.prev = head
    head.next = n
}

func insertToFront(li list, n *node) {
    head := li.head

    head.next.prev = n
    n.next = head.next
    n.prev = head
    head.next = n
}

func removeLast(li list) {
    tail := li.tail

    tail.prev.prev.next = tail
    tail.prev = tail.prev.prev
}

func Constructor(capacity int) LRUCache {
    head, tail := &node{prev: nil}, &node{next: nil}
    head.next, tail.prev = tail, head

    return LRUCache{
        capacity: capacity,
        currSize: 0,
        record: make(map[int]*entry),
        li: list{
            head: head,
            tail: tail,
        },
    }
}

func (this *LRUCache) Get(key int) int {
    if e, ok := this.record[key]; ok {
        moveNodeToFront(this.li, e.n)
        return e.value
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if e, ok := this.record[key]; ok {
        moveNodeToFront(this.li, e.n)
        this.record[key].value = value
    } else {
        if (this.currSize == this.capacity) {
            delete(this.record, this.li.tail.prev.key)
            removeLast(this.li)
            this.currSize--
        }

        n := &node{key: key}
        insertToFront(this.li, n)

        this.record[key] = &entry{
            value: value,
            n: n,
        }

        this.currSize++
    }
}

func main(){

  obj := Constructor(2)
  obj.Put(2,20)
  obj.Put(3,30)
  obj.Put(1,10)
  param_1 := obj.Get(3)
  fmt.Println(param_1)
  param_2 := obj.Get(1)
  fmt.Println(param_2)

}
