package Mixed

import (
	"fmt"
	"strings"
)

func TestMixed() {
	fmt.Println(UpperAllFirstLetter("merhaba, bu bir test yazısıdır."))
}

// Makes uppercase all first letters of a string.
func UpperAllFirstLetter(s string) string {

	var builder strings.Builder
	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		letter := []byte(words[i])
		firstLetter := string(letter[0])

		words[i] = strings.Replace(words[i], firstLetter, strings.ToUpper(firstLetter), 1)

		builder.WriteString(words[i] + " ")
	}

	return builder.String()
}
