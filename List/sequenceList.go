package List

import "fmt"

type Elem int
type SqList struct {
	MaxSize int
	Length  int
	Data    []Elem
}

func New(maxSize int) *SqList {
	return &SqList{MaxSize: maxSize, Data: make([]Elem, maxSize)}
}

func (list *SqList) IsEmpty() bool {

	return 0 == list.Length

}

func (list *SqList) IsFull() bool {
	return list.Length == list.MaxSize
}

func (list *SqList) Insert(i int, elem Elem) bool {
	if i < 1 || i > list.Length {
		fmt.Print("please check i")
		return false
	}
	for k := list.Length; k > i-1; k-- {
		list.Data[k] = list.Data[k-1]
	}
	list.Data[i-1] = elem
	list.Length++
	return true
}

func (list *SqList) Del(i int) bool {
	if i < 1 || i > list.Length {
		fmt.Println("please check i")
		return false
	}
	for k := list.Length; k > i-1; k-- {
		list.Data[k] = list.Data[k+1]
	}
	list.Data[list.Length-1] = 0
	list.Length--
	return true
}

func (list *SqList) GetElem(i int) Elem {
	if i < 1 || i > list.Length {
		fmt.Println("please check i")
		return -1
	}
	return list.Data[i-1]

}

func (list *SqList) Append(elem Elem) bool {
	if list.IsFull() {
		return false
	}
	list.Data[list.Length] = elem
	list.Length++
	return true
}
