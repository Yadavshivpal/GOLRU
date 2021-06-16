package main

import "fmt"
import l "github.com/Yadavshivpal/GOLRU/lru"



func main(){
  fmt.Println("enter 1 for run code or enter 2 for exit ")
  var a int
  fmt.Scanln(&a)
  if a==1{
    fmt.Println("Enter capacity")
    var cap int
    fmt.Scanln(&cap)
    obj := l.Newcache(cap)
    fmt.Println("Enter keys and value: ")
    for i:=0;i<cap;i++{
      var k int
      var v int
      fmt.Scanln(&k)
      fmt.Scanln(&v)
      obj.Put(k, v)
    }
    var count = true
    for count{
      fmt.Println("Enter 1 for GET key ")
      fmt.Println("Enter 2 for PUT key")
      fmt.Println("For exit Enter 3")
      var i int
      fmt.Scanln(&i)
      switch i {
      case 1:
        fmt.Println("Enter the KEY to get value ")
        var input int
        fmt.Scanln(&input)
        fmt.Println(obj.Get(input))
      case 2:
        fmt.Println("Enter the key and value insert")
        var k int
        var v int
        fmt.Scanln(&k)
        fmt.Scanln(&v)
        obj.Put(k, v)
      case 3:
        count = false
        fmt.Println("Thnank you")
      }

    }
  }else
  {
    fmt.Println("Thnank you")
  }


}
