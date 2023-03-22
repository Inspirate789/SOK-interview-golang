package main

import "fmt"

type Stack interface {
	Push(int)
	Pop(int) bool
}

type stackRecord struct {
	data, curMax int
}

type StackImpl struct {
	sp    int
	elems []stackRecord
}

func NewStackImpl() *StackImpl {
	return &StackImpl{sp: -1}
}

func (s *StackImpl) appendElem(elem stackRecord) {
	if s.sp == len(s.elems)-1 {
		s.elems = append(s.elems, elem)
	} else {
		s.elems[s.sp+1] = elem
	}
	s.sp++
}

func (s *StackImpl) Push(elem int) {
	if s.sp == -1 {
		s.appendElem(stackRecord{elem, elem})
	} else {
		var newMax int
		if s.elems[s.sp].curMax > elem {
			newMax = s.elems[s.sp].curMax
		} else {
			newMax = elem
		}

		s.appendElem(stackRecord{elem, newMax})
	}
}

func (s *StackImpl) Pop() (int, bool) {
	if s.sp == -1 {
		return 0, false
	}

	elem := s.elems[s.sp].data
	s.sp--

	return elem, true
}

func (s *StackImpl) GetMaxElem() (int, bool) {
	if s.sp == -1 {
		return 0, false
	}

	return s.elems[s.sp].curMax, true
}

func main() {
	s := NewStackImpl()
	s.Push(1)
	s.Push(2)
	s.Push(0)
	s.Push(8)
	s.Push(5)
	max, _ := s.GetMaxElem()
	fmt.Printf("Max: %d\n", max)

	for elem, ok := s.Pop(); ok != false; elem, ok = s.Pop() {
		fmt.Printf("Pop element: %d\n", elem)
		max, exist := s.GetMaxElem()
		if exist == false {
			fmt.Println("Cannot get max data")
			break
		}
		fmt.Printf("Current max element: %d\n", max)
	}
}
