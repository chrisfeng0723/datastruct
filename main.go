package main

import (
	"datastruct/LinkList"
)

func main(){
	l := LinkList.LinkList{HeadNode:LinkList.NewNode(nil)}
	//fmt.Println(l)
	l.Add(3)
	//l.Add(4)
	//fmt.Println(l)
	//fmt.Println(l.IsEmpty())
	//fmt.Println(l.Length())
	l.Print()
}
