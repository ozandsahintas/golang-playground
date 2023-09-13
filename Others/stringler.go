package Others

import (
	"bytes"
	"fmt"
	"strings"
)

func Strings() {

	fmt.Println("\n", "~~~ REPEAT ~~~")
	q := strings.Repeat("Test ", 5)
	fmt.Println(q)

	fmt.Println("\n", "~~~ UPPER & LOWER ~~~")
	q = strings.ToUpper("TeSt YaZısı test - 2")
	fmt.Println(q)
	q = strings.ToLower("TeSt YaZısı test - 2")
	fmt.Println(q)

	fmt.Println("\n", "~~~ TRIM ~~~")
	q = strings.Trim("__MESAJ__", "_")
	fmt.Println(q)

	//q = strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")
	//q = strings.TrimPrefix("HEY_MESAJ__", "HEY_")
	q = strings.TrimLeft("__MESAJ__", "_")
	fmt.Println(q)

	//q = strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡")
	q = strings.TrimRight("__MESAJ__", "_")
	fmt.Println(q)

	f1 := "          Test deneme test"
	fmt.Println(f1)
	fmt.Println(strings.TrimSpace(f1))

	fmt.Println("\n", "~~~ COMPARE ~~~")

	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	fmt.Println("\n", "~~~ REPLACER ~~~")
	s1 := "This is _A MESSAGE_ such @HUGE MESSAGE@ <===>"
	fmt.Println(s1)
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;", "_", "-", "@", "*")
	fmt.Println(r.Replace(s1))

	fmt.Println("\n", "~~~ BUILDER ~~~")
	builder := strings.Builder{}

	builder.WriteString("Hello and ")
	builder.Write([]byte{72, 101, 108, 108, 111})

	fmt.Println(builder.String(), " -> Length: ", builder.Len())

	fmt.Println("\n", "~~~ REPLACE ALL & CONTAINS ~~~")
	s2 := "Askerler"
	if strings.Contains(s2, "ler") {
		fmt.Println(strings.ReplaceAll(s2, "ler", ""))
	}

	fmt.Println("\n", "~~~ SPLIT ~~~")
	s3 := "bir numara, iki numara, üç numara"
	fmt.Println(ReplaceSeperator(",", "|-|", s3))
	fmt.Println(strings.ReplaceAll(s3, ",", "|-|"))

	fmt.Println("\n", "~~~ BYTE ARRAYS ~~~")

	fmt.Println([]byte("Ozan"), " - ", []byte("ozan"))
	fmt.Println("Ozan > ozan: ", "Ozan" > "ozan", "\n")

	fmt.Println([]byte("a"), " - ", []byte("b"))
	fmt.Println("a > b:", "a" > "b", "\n")

	fmt.Println([]byte("Ozan"), " - ", []byte("Okan"))
	fmt.Println("Ozan > Okan: ", "Ozan" > "Okan", "\n")

	d1 := []byte("Selam")
	d1[0] = d1[0] + 8
	fmt.Println(string(d1))

	text := "yalnızca baş harfleri büyük olacak yazı."
	fmt.Println("Before: ", text)
	fmt.Println("After: ", ReplaceAllCapital(text))

	bytesOfArray := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33}
	str1 := bytes.NewBuffer(bytesOfArray).String()
	fmt.Println("Converting into String =", str1)
	str2 := fmt.Sprintf("%s", bytesOfArray)
	fmt.Println("str2 =", str2)
}

func ReplaceSeperator(o string, n string, t string) string {
	c := strings.Split(t, o)

	var b strings.Builder
	b.WriteString(strings.Join(c, n))

	return b.String()
}

func ReplaceAllCapital(s string) string {

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
