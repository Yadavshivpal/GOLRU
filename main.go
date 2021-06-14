package main

import "fmt"
import l "github.com/Yadavshivpal/GOLRU/cache"

func main(){

  obj := l.Constructor(2)
  obj.Put(2,20)
  obj.Put(3,30)
  obj.Put(1,10)
  out := obj.Get(3)
  fmt.Println(out)
  out1 := obj.Get(1)
  fmt.Println(out1)

}
