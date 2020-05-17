package main

import (
	"fmt"
)

func main() {
	//Hello World :  🙏 - K & R
	fmt.Println("Hello, World! 🙏 - K & R")

	name := "Praseed" // var name string = "sadas"
	year := 2020      //Corona Year-- var year int = 2020
	//:= does variable declaration and assignment, golang supports localvariable typeinference
	//Formating and Printing
	fmt.Printf("Hello %s, It's %d\n", name, year)

	name = "Sarath"
	//Formatting to String
	message := fmt.Sprintf("Hello %s, It's %d", name, year)
	fmt.Println(message)

	//Data types
	var b bool = true
	var ui8 uint8 = 255                    //max value
	var ui16 uint16 = 65535                //max value
	var ui32 uint32 = 4294967295           // max value
	var ui64 uint64 = 18446744073709551615 // max value
	var i8 int8 = int8(ui8 / 2)
	var i16 int16 = int16(ui16 / 2)
	var i32 int32 = int32(ui32 / 2)
	var i64 int64 = int64(ui64 / 2)
	var f32 float32 = 1.21231213123123112312312312
	var f64 float64 = 1.21231213123123112312312312
	var c64 complex64 = complex(f32, f32)
	var c128 complex128 = complex(f64, f64)
	var sPtr *string = &name
	var i16Ptr *int16 = &i16
	fmt.Println(b, ui8, ui16, ui32, ui64, i8, i16, i32, i64, f32, f64, c64, c128, real(c64), imag((c128)), sPtr, i16Ptr, *sPtr, *i16Ptr)

	//Control flow
	if year := 2016; year > 2008 {
		fmt.Println("Gt than 2008")
	} else if year > 2000 {
		fmt.Println("Gt than 2000 but less that or equal to 2008")
	} else {
		fmt.Println("Less than 2000")
	}

	switch name := "X"; name {
	case "X":
		fmt.Println("Its' X")
	case "Z", "z":
		fmt.Println("Got Z")
		fallthrough
	case "Y":
		fmt.Println("Got Y")
		fmt.Println("Its' not X")
	default:
		fmt.Println("Its' Unknown")
	}

	//Loop 	- Go has only for loops (Which comes in many forms)
	//Print 1 to 10
	for i := 0; i < 5; i++ {
		fmt.Println(fmt.Sprintf("%d: Hello beautiful simple for loop", i))
	}

	//Declaring and initializing an array
	nums := []uint32{1, 2, 3, 4, 5}
	//Range loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(fmt.Sprintf("%d: Hello cool loop", nums[i]))
	}
	//Range loop
	for _, num := range nums { // _ ignores the value
		//Alternative: => for i, _ := range nums { nums[i]}
		fmt.Println(fmt.Sprintf("%d: Hello range loop", num))
	}

	//map aka associated array
	associate := map[string]string{
		"hello": "world",
		"foo":   "bar",
		"lorem": "ipsum",
	}

	//range loop over map keys
	for key := range associate {
		fmt.Println(key, ": ", associate[key])
	}

	//range loop over map key value
	for key, value := range associate {
		fmt.Println(key, ": ", value)
	}

	//Add values to an integer list
	intList := make([]uint32, 0)
	fmt.Println(fmt.Sprintf("Cap: %d, Len: %d", cap(intList), len(intList)))
	for i := 0; i < 1000; i++ {
		intList = append(intList, uint32(i))
	}
	fmt.Println(fmt.Sprintf("After appending Cap: %d, Len: %d", cap(intList), len(intList)))

	//Add values to a map
	coutMapOfOddAndEven := make(map[string]uint64)
	fmt.Println(fmt.Sprintf("Len: %d", len(coutMapOfOddAndEven)))
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			coutMapOfOddAndEven["even"]++
		} else {
			coutMapOfOddAndEven["odd"]++
		}
	}
	fmt.Println(fmt.Sprintf("Even numbers: %d", coutMapOfOddAndEven["even"]))
	fmt.Println(fmt.Sprintf("Odd numbers: %d", coutMapOfOddAndEven["odd"]))
	delete(coutMapOfOddAndEven, "odd")
	fmt.Println(fmt.Sprintf("Odd numbers: %d", coutMapOfOddAndEven["odd"]))

	//Find a pythogorean triplet, given the hypotenuse find the other two sides
	k := 5
	for i := 1; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if k*k == i*i+j*j {
				fmt.Println(i, j, k)
			}
		}
	}

	//functioncall
	fmt.Println(*sayHello("Kochi"))
	fmt.Println(func(x, y int) int {
		return x + y
	}(10, 20))
	//Lambda expression - anonynous func
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(10, 20))
	//closure
	z := 100
	add = func(x, y int) int {
		return x + y + z
	}
	fmt.Println(add(10, 20))
	z = 200
	fmt.Println(add(10, 20))

	//Careful with captured variable and mutation
	for i := 0; i < 3; i++ {
		if i == 0 {
			captured := i
			add = func(x, y int) int {
				return x + y + captured
			}
		}
	}
	fmt.Println(add(10, 20))

	//structure

	movie := Movie{Name: "Aham", Rating: 4} //Obeys value semantics comparison with nil fails
	fmt.Println(fmt.Sprintf("%+v", movie))
	movie.RateMovie()
	fmt.Println(fmt.Sprintf("%+v", movie))

	moviePtr := new(Movie)
	moviePtr.Name = "Sandesham"
	moviePtr.Rating = 4
	fmt.Println(fmt.Sprintf("%+v", moviePtr))
	moviePtr.RateMovie()
	fmt.Println(fmt.Sprintf("%+v", moviePtr))

	movieRef := &Movie{Name: "Guru", Rating: 3}
	fmt.Println(fmt.Sprintf("%+v", movieRef))

	var movieRef2 *Movie
	fmt.Println(fmt.Sprintf("Is nil = %t", movieRef2 == nil))

	//Type embedding
	pet := Pet{Name: "Red Chameleons", Species: "Chameleons"}
	fmt.Printf("%+v\n", pet)

	var dog = Dog{Pet{Name: "Husky", Species: "Dog"}, true}
	fmt.Printf("%+v\n", dog)
	fmt.Printf("%s %s %s %s %t\n", dog.Pet.Name, dog.Pet.Species, dog.Name, dog.Species, dog.ChasesCats)
	cat := Cat{Pet{Name: "Persian Cat", Species: "Cat"}, true}
	fmt.Printf("%+v\n", cat)

	//calling methods
	pet.Run()
	dog.Run() //Methods defined on the embedding type can be inherited
	cat.Run()

	//shallowcopy
	x := X{64, 'A'}
	y := Y{&x, 'B'}
	tmp := y
	tmp.Category = 'A'
	fmt.Printf("X: %+v\n", x)
	fmt.Printf("Y: %+v\n", y)
	fmt.Printf("Tmp: %+v\n", tmp)

	//Public and Private accesors in golang
	//Only names starting with capital letters will be available to files outside the defining one
	//Hello Script is supposed to be a single file, try it yourself
}

func sayHello(name string) *string {
	if len(name) == 0 {
		return nil
	}
	value := fmt.Sprintf("Hello %s", name)
	return &value
}

type Movie struct {
	Name   string
	Rating int
}

//Static duck typing

type IMovie interface {
	RateMovie()
}

// func (m Movie) RateMovie() {
// 	//m is passed by value no effect on modification
// 	m.Rating = (len(m.Name) % 10) + 1 //IMDBs rating algorithm. True story!
// }

func (m *Movie) RateMovie() {
	m.Rating = (len(m.Name) % 10) + 1 //IMDBs rating algorithm. True story!
}

//Type embedding
type Pet struct {
	Name    string
	Species string
}

type Dog struct {
	Pet
	ChasesCats bool
}

type Cat struct {
	Pet
	HatesDogs bool
}

type Run interface {
	Run()
}

func (p *Pet) Run() {
	fmt.Println(fmt.Sprintf("%s is running", p.Name))
}

type X struct {
	Number   int32
	Category rune
}

type Y struct {
	*X
	AdditionalCategory rune
}

//public private
type Public struct {
	Pub, priv int
}

type private struct {
	Pub, priv int
}
