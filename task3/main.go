package main

import (
	"fmt"
	"log"
)

type stackRecord struct {
	data byte
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

func (s *StackImpl) Push(elem byte) {
	if s.sp == -1 {
		s.appendElem(stackRecord{elem})
	} else {
		s.appendElem(stackRecord{elem})
	}
}

func (s *StackImpl) Pop() (byte, bool) {
	if s.sp == -1 {
		return 0, false
	}

	elem := s.elems[s.sp].data
	s.sp--

	return elem, true
}

func (s *StackImpl) Top() (byte, bool) {
	if s.sp == -1 {
		return 0, false
	}

	return s.elems[s.sp].data, true
}

func (s *StackImpl) Empty() bool {
	return s.sp == -1
}

func main() {
	s := NewStackImpl()
	var str string
	fmt.Print("Enter the string: ")
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '[' || str[i] == '{' {
			s.Push(str[i])
		} else {
			top, _ := s.Top()
			if str[i] == ')' && !s.Empty() && top == '(' {
				s.Pop()
			} else if str[i] == ']' && !s.Empty() && top == '[' {
				s.Pop()
			} else if str[i] == '}' && !s.Empty() && top == '}' {
				s.Pop()
			} else if (str[i] == ')' || str[i] == ']' || str[i] == '}') && s.Empty() {
				s.Push(byte(i))
				break
			}
		}
	}

	fmt.Println(s.Empty())
}
