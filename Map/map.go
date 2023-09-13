package Map

import "fmt"

func Maps() {
	map_1 := map[int]string{

		1: "Dog",
		2: "Cat",
		3: "Cow",
		4: "Bird",
		5: "Rabbit",
	}

	var map_2 map[int]int
	var map_3 = make(map[int]int)

	fmt.Println("map_1: ", map_1)
	fmt.Println("map_2: ", map_2)
	fmt.Println("map_3: ", map_3)
	fmt.Println("_______________")
	fmt.Println("map_1[4] öncesi: ", map_1[4]) // Kullanırken direkt bu şekil.
	map_1[4] = "Kuş"                           // Var olan kaydı güncelliyorum.
	fmt.Println("map_1[4] sonrası: ", map_1[4])
	fmt.Println("_______________")
	fmt.Println("map_1[99] öncesi: ", map_1[99])  // Böyle bir kayıt yok.
	map_1[99] = "Ördek"                           // Yeni kayıt ekliyorum.
	fmt.Println("map_1[99] sonrası: ", map_1[99]) // Artık böyle bir kayıt var
	fmt.Println("_______________")
	iterateOverMap(map_1)
	fmt.Println("_______________")
	data_99, exists_99 := map_1[99]
	fmt.Println("id: ", data_99)
	fmt.Println("isExists: ", exists_99)
	data_98, exists_98 := map_1[98]
	fmt.Println("id: ", data_98)
	fmt.Println("isExists: ", exists_98)
	fmt.Println("_______________")
	delete(map_1, 99)
	fmt.Println("map_1: ", map_1)
	fmt.Println("_______________")
	var map_4 = map_1 // Map bir REFERANCE TYPE, bu yüzden ikisi de aynı şey.
	fmt.Println("map_4 öncesi: ", map_4)
	delete(map_1, 5)
	fmt.Println("map_4 sonrası: ", map_4)
}

func iterateOverMap(m map[int]string) {
	for id, pet := range m {
		fmt.Println(id, pet)
	}

}
