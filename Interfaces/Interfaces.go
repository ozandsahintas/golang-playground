package Interfaces

import (
	"fmt"
	"io"
	"strings"
)

func emptyInterface(t interface{}) string {
	return fmt.Sprintf("%v", t)
}
func EmptyInterface() {
	var v interface{}

	fmt.Printf("type: %T\n", v)
	fmt.Printf("value: %v\n", v)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func InterfaceReceiver() {
	fmt.Println(emptyInterface("hello"))
	fmt.Println(emptyInterface(1))
	fmt.Println(emptyInterface(3.14))

	u := struct {
		name string
	}{
		"vigo",
	}
	fmt.Println(emptyInterface(u))

	fmt.Println(emptyInterface([]string{"hello"}))
	fmt.Println(emptyInterface(nil))
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type customInt int

func printByType(t interface{}) {
	// .(type) sadece switch içinde çalışır
	switch j := t.(type) {
	case nil:
		fmt.Println(j, " this is nil")
	case int:
		fmt.Println(j, " this is int")
	case customInt:
		fmt.Println(j, " this is customInt")
	case io.Reader:
		fmt.Println(j, " this is io.Reader")
	case string:
		fmt.Println(j, " this is string")
	case bool, rune:
		fmt.Println(j, " this is bool or rune")
	default:
		fmt.Printf("%v no idea: %[1]T\n", j)
	}
}
func InterfaceTypeDetection() {
	printByType(nil)
	printByType(1)
	printByType(3.14)
	printByType("hello")
	printByType(true)
	printByType('a')

	var a customInt
	a = 5
	printByType(a)
	printByType(strings.NewReader("hello"))
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type myInt int

func InterfaceTypeChange() {
	var i interface{}

	fmt.Println("i", i) // i <nil>

	i = 1
	fmt.Println("int", i, i.(int)) // int 1 1

	i = "hello"
	fmt.Println("string", i, i.(string)) // string hello hello

	i = 3.14
	fmt.Println("float", i, i.(float64)) // float 3.14 3.14

	i = []string{"hello"}
	fmt.Println("[]string", i, i.([]string)) // []string [hello] [hello]

	i = map[string]string{
		"hello": "ozan",
	}
	fmt.Println("[]map[string]string", i, i.(map[string]string))
	// []map[string]string map[hello:world] map[hello:world]

	i = myInt(5)
	fmt.Println("myInt", i, i.(myInt)) // myInt 5 5
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type Number int

func (n Number) Positive() bool {
	return n > 0
}

func main() {

	x := Number(5)
	fmt.Println(x.Positive()) // true
}
