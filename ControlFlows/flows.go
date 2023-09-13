package ControlFlows

import "fmt"

func IfStatement() {
	var v = 70
	if v < 200 {
		fmt.Printf("v < 200\n")
	} else if v > 200 && v < 1000 {
		fmt.Printf("200 < v < 1000\n")
	} else {
		// Bu else'i alt satıra yazmaya çalıştığımızda hata veriyor!
		fmt.Printf("v > 1000\n")
	}

	// Bu şekilde de yazabiliyoruz.
	if n := v; n%2 == 0 {
		fmt.Printf("%d bir çift sayıdır.\n", n)
	} else {
		fmt.Printf("%d bir tek sayıdır.\n", n)
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Parantezli

	//if (k := v; n%2 == 0) { // Kısa gösterimde (short statement): Syntax Error
	//	...
	//}
	//
	//if(v < 200){ // Bu şekilde sıkıntı yok.
	//	...
	//}

}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func SwitchStatement() {

	var dayOfWeek = 5
	switch dayOfWeek {
	//Şu şekilde de yazılabilirdi (Short statement)-->    switch dayOfWeek := 6; dayOfWeek {
	case 1:
		fmt.Println("Pazartesi")
	case 2:
		fmt.Println("Salı")
	case 3:
		fmt.Println("Çarşamba")
	case 4:
		fmt.Println("Perşembe")
	case 5:
		fmt.Println("Cuma")
		fmt.Println("TGIF!")
	case 6:
		fmt.Println("Cumartesi")
	case 7:
		fmt.Println("Pazar")
	default:
		fmt.Println("Geçersiz!")
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////

	switch dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Hafta içi")
	case 6, 7:
		fmt.Println("Hafta sonu")
	default:
		fmt.Println("Böyle bi gün yok!")
	}

	///////////////////////////////////////////////////////////////////////////////////////////////////////////

	// İçi boş switch ile sanki bir if-else yapısı gibi de yazılabiliyor.

	switch {
	case dayOfWeek == 1, dayOfWeek == 2, dayOfWeek == 3, dayOfWeek == 4, dayOfWeek == 5:
		fmt.Println("Hafta içi")
	case dayOfWeek == 6, dayOfWeek == 7:
		fmt.Println("Hafta sonu")
	default:
		fmt.Println("Böyle bi gün yok!")
	}
}

func ForLoop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	i := 2
	for ; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
}

func ForLoopAsWhileLoop() {
	i := 2
	for i <= 20 {
		fmt.Printf("%d ", i)
		i *= 2
	}
}

func InfiniteLoop() {
	for {
		fmt.Println("Sonsuz döngü!")
	}
}

func BreakAndContinueInAForLoop() {
	for num := 1; num <= 100; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Printf("3 ve 5'in EKOK değeri: %d\n", num)
			break
		}
	}

	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			continue
		}
		fmt.Printf("%d ", num)
	}
}

func RangeInForLoop() {
	users := []string{"ozan", "deniz", "şahintaş"}

	usersMap := map[string]string{
		"ozan":     "user",
		"deniz":    "admin",
		"şahintaş": "user",
	}

	for i := range users {
		fmt.Println("index", i)
	}

	for i, user := range users {
		fmt.Println(i, user)
	}

	for _, user := range users {
		fmt.Println(user)
	}

	for user := range usersMap {
		fmt.Println(user, usersMap[user])
	}

	for user, level := range usersMap {
		fmt.Println(user, level)
	}
}

func LabelInForLoop() {
outer:
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			if j == 2 {
				break outer // exit!
			}
		}
	}
}
