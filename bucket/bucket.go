package bucket

type Node struct {
    Key int
    Next *Node
    Prev *Node
}

type List struct {
    Head *Node
    Tail *Node
}

type Entry struct {
    Value int
    N *Node
}

func Movenode(li List, n *Node) {
    Head := li.Head

    n.Next.Prev = n.Prev
    n.Prev.Next = n.Next

    n.Next = Head.Next
    Head.Next.Prev = n

    n.Prev = Head
    Head.Next = n
}

func Insertnode(li List, n *Node) {
    Head := li.Head

    Head.Next.Prev = n
    n.Next = Head.Next
    n.Prev = Head
    Head.Next = n
}

func Removenode(li List) {
    Tail := li.Tail
    Tail.Prev.Prev.Next = Tail
    Tail.Prev = Tail.Prev.Prev
}
