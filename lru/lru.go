package lru

import l "github.com/Yadavshivpal/GOLRU/bucket"
import "fmt"

type LRUCache struct {
    capacity int
    currSize int
    Mapping map[int]*l.Entry
    Li l.List
}

func Newcache(capacity int) LRUCache {
    head, tail := &l.Node{Prev: nil}, &l.Node{Next: nil}
    head.Next, tail.Prev = tail, head

    return LRUCache{
        capacity: capacity,
        currSize: 0,
        Mapping: make(map[int]*l.Entry),
        Li: l.List{
            Head: head,
            Tail: tail,
        },
    }
}

func (this LRUCache) Printall(){
  for k, v := range this.Mapping {
      fmt.Println("Key is: ",k, "value is: ",v.Value)
  }
}

func (this *LRUCache) Get(key int) int {
    if e, ok := this.Mapping[key]; ok {
        l.Movenode(this.Li, e.N)
        return e.Value
    }

    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    if e, ok := this.Mapping[key]; ok {
        l.Movenode(this.Li, e.N)
        this.Mapping[key].Value = value
    } else {
        if (this.currSize == this.capacity) {
            delete(this.Mapping, this.Li.Tail.Prev.Key)
            l.Removenode(this.Li)
            this.currSize--
        }

        n := &l.Node{Key: key}
        l.Insertnode(this.Li, n)

        this.Mapping[key] = &l.Entry{
            Value: value,
            N: n,
        }

        this.currSize++
    }
}
