package Others

import (
	"fmt"
	"strconv"
)

//func Atoi(str string) (int, error)

func AtoiGo() {

	string_value := "245"
	fmt.Printf("Öncesi: Type: %T - Value: %v", string_value, string_value)
	int_value, err := strconv.Atoi(string_value)
	if err == nil {
		fmt.Printf("\nSonrası: Type: %T  - Value: %v", int_value, int_value)
	}
}

func TypeConversion() {
	integer_to_convert := 245
	fmt.Printf("\nType: %T - Value: %v", integer_to_convert, integer_to_convert)
	float_value := float64(integer_to_convert)
	fmt.Printf("\nType: %T - Value: %v", float_value, float_value)
	uint_value := uint(float_value)
	fmt.Printf("\nType: %T - Value: %v", uint_value, uint_value)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BitCount() {
	_ = bitCount(28)
	//fmt.Println(x)
}

func bitCount(number int) int {
	if number == 1 {
		fmt.Print("1")
	} else if number%2 == 0 {
		fmt.Print("0")
		bitCount(number / 2)
	} else {
		fmt.Print("1")
		bitCount((number - 1) / 2)
	}
	return 0
}
