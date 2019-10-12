package LinkList

import "fmt"

type Elem interface{}
type Node struct {
	Data interface{}
	Next *Node
}
type LinkList struct {
	HeadNode *Node
}

func NewNode(data Elem) *Node {

	return &Node{Data: data, Next: nil}

}

func (l *LinkList) Add(e Elem) *Node {
	node := NewNode(e)
	node.Next = l.HeadNode
	l.HeadNode = node
	return node
}
func (l *LinkList) Append(e Elem) {
	node := NewNode(e)
	cur := l.HeadNode
	if l.IsEmpty() {
		l.HeadNode = node
	}
	for nil != cur {
		cur = cur.Next
	}
	cur.Next = node

}

func (l *LinkList) Insert(index int, data Elem) {
	node := NewNode(data)
	if index < 0 {
		l.Add(data)
	} else if index > l.Length() {
		l.Append(data)
	} else {
		pre := l.HeadNode
		for i := 0; i < index; i++ {
			pre = pre.Next
		}
		node.Next = pre.Next
		pre.Next = node

	}

}

func(l *LinkList) Remove()

func (l *LinkList) GetElem(i int) interface{} {
	cur := l.HeadNode
	j := 1
	for j < i && cur != nil {
		cur = cur.Next
	}
	return cur.Data
}

func (l *LinkList) Print() {
	//cur := l.HeadNode
	//for nil != cur {
	//
	//	fmt.Printf("\t%v",cur.Data)
	//	cur = cur.Next
	//}
	if !l.IsEmpty() {
		cur := l.HeadNode
		for {
			fmt.Printf("\t%v", cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	}
}

func (l *LinkList) IsEmpty() bool {
	return l.HeadNode.Next == nil
}
func (l *LinkList) Length() int {
	cur := l.HeadNode
	count := 0
	for nil != cur {
		count++
		cur = cur.Next
	}
	return count
}
