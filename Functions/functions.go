package Functions

import "fmt"

func FunctionSignature(a string, b int) {}

func AnonimFonksiyon() {
	a := func() {
		fmt.Println("Anonim fonksiyon tanımlandı ve çağırıldı.")
	}
	a()                          // Anonim fonksiyonu çağırıyorum
	fmt.Printf("Tipi: %T \n", a) // Tipi ne ola ki?

	///////////////////////////////////////////////////////////////////////////////////////////////////////////

	func() {
		fmt.Println("Anonim fonksiyon tanımlandığı gibi çağrıldı.")
	}()

	///////////////////////////////////////////////////////////////////////////////////////////////////////////

	func(n string) {
		fmt.Println("Parametreli anonim fonksiyona yollanan değer: ", n)
	}("Gophers")
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

type add func(a int, b int) int

func DegiskenOlarakFonksion() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Toplam: ", s)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func fonsiyonaParametre(a func(a, b int) int) {
	fmt.Println(a(60, 7)) // Parametre olarak gelen fonksiyona parametre yolluyorum.
}

func ParametreOlarakFonksiyon() {
	f := func(a, b int) int {
		return a + b
	}

	fonsiyonaParametre(f) // Değişken olarak yaratılan fonksiyon, başka bir fonksiyona parametre olarak gönderiliyor.
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func fonksiyondanDonus() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func DonusDegeriOlarakFonksiyon() {
	s := fonksiyondanDonus()
	fmt.Println(s(60, 7))
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func passByValue(s string) {
	fmt.Println("incoming", s)
	s = "change"
	fmt.Println("changed", s)
}

func passByRef(s *string) {
	fmt.Println("incoming", *s)
	*s = "change"
	fmt.Println("changed", *s)
}

func PassByValueAndParameter() {
	m := "hello"
	passByValue(m)
	fmt.Println("m", m)

	passByRef(&m)
	fmt.Println("m", m)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func greet(names ...string) {
	for _, name := range names {
		fmt.Println("hello " + name)
	}
}

func CokluParametre() {
	greet("ali", "veli")
	greet("can", "mert", "ozan", "hasan")
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func cokluDon(n string) (string, error, int) {
	return n, nil, 0
}

func CokluDonus() {
	m, _, x := cokluDon("hello") // Error'u kullanmayacağız o yüzden "_"
	fmt.Println("m", m)
	fmt.Println("x", x)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func nakedDonus(name string) (result string) {
	result = "Hello " + name
	return
	// Herhangi bir return değeri belirtmedik ama Go neyi döneceğimizi anladı.
	// Aslında return değerini fonksiyonu tanımlarken dönüş değişkeni olarak belirtiyoruz.
}

func NakedReturn() {
	fmt.Println(nakedDonus("Ozan")) //
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

//func NakedShadowing() (id int, err error) {
//	id = 10
//
//	if id == 10 {
//		err := fmt.Errorf("Invalid Id\n") // Compiler-Time Error: inner declaration of var err error
//		fmt.Println("err: ", err)
//		return // Compile-Time Error: result parameter err not in scope at return
//	}
//
//	return
//}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func Recursive() {
	fmt.Println(fact(3)) // 6
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

// Closures

func fib() func() int { // fib() fonksiyonu, integer dönüş tipine sahip bir fonksiyon döndüren bir fonksiyon.
	a, b := 0, 1

	// bu fonksiyon a ve b yi kullanabilir
	return func() int {
		a, b = b, a+b
		return b
	}
}

func ClosesOverFonksiyon() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

func deferredFunc(n int) (x int) {
	x = n
	defer func() {
		x = 10
		fmt.Println("deferredFunc defer ediliyor.")
	}()
	fmt.Println("deferredFunc işlemleri bitti. x: ", x)
	return x
}

func DeferFonksiyon() {
	defer func() {
		fmt.Println("DeferFonksiyon defer ediliyor.")
	}()

	var ret = deferredFunc(100) // Defer fonksiyonumuz dönüş değerini ezdi. Dönüşü 100 bekliyorduk.
	fmt.Println("DeferFonksiyon işlemleri bitti. x: ", ret)
}
