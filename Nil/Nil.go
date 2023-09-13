package Nil

import (
	"fmt"
	"strings"
)

var s *string
var ss = new(string)

func NewKeyword() {
	fmt.Printf("type: %T value: %[1]v\n", s)                  // type: *string value: <nil>
	fmt.Printf("type: %T value: %[1]p pvalue: %v\n", ss, *ss) // type: *string value: 0x14000104210 pvalue:  - yani ""
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type stack struct {
	data []string // "zero value"
}

// Push appends new string to data
func (s *stack) Push(v string) {
	s.data = append(s.data, v)
}

// Pop removes last item from data
func (s *stack) Pop() string {
	if l := len(s.data); l > 0 {
		p := s.data[l-1]
		s.data = s.data[:l-1]
		return p
	}
	return "empty stack"
}

func (s stack) String() string {
	if s.data == nil {
		return "<niiiiiiiiiil>"
	}
	return strings.Join(s.data, ",")
}

func StackImplementation() {

	users := stack{}
	fmt.Println("users", users)
	fmt.Printf("%T %[1]v %[1]p\n", users.data)
	fmt.Println(users.data == nil)

	users.Push("vigo")
	users.Push("turbo")
	fmt.Println("users", users)
	var x = users.Pop()
	fmt.Println("users", users)
	fmt.Println("users: ", x)
	fmt.Println(users.data == nil)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type list struct {
	val int
}

func (l *list) sum() int {
	if l == nil {
		return 0
	}
	return l.val
}

func DefaultValue() {
	ff := list{}
	fmt.Println(ff.sum()) // 0
	fmt.Printf("type: %T value: %[1]p pvalue: %v\n", ff.val)

	ff.val = 10
	fmt.Println(ff.sum()) // 10
	fmt.Printf("type: %T value: %[1]p pvalue: %v\n", ff.val)
}
