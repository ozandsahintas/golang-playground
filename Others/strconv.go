package Others

import (
	"fmt"
	"strconv"
)

func Strconv() {
	fmt.Println(quote())
}

func quote() string {
	q := strconv.Quote("Hello, 世界")
	return q
}
