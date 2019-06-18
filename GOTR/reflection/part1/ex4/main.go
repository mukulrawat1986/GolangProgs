package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	Println(true)
	Println(2010)
	Println(9.14)
	Println(7 + 7i)
	Println("Hello, World!")
	Println(ID("fkdjfd"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Person{name: "Anna"})
	Println(&Person{name: "Anna"})
	Println(make(chan int))
	Println(nil)
	var i interface{}

	Println(i)
}

// Println is my simple print function
// Example of type switching
func Println(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Println(x.(bool))

	case int:
		fmt.Println(x.(int))

	case float64:
		fmt.Println(x.(float64))

	case complex128:
		fmt.Println(x.(complex128))

	case string:
		fmt.Println(x.(string))

	case Person:
		fmt.Println(x.(Person))

	case *Person:
		fmt.Println(x.(*Person))

	case chan int:
		fmt.Println(x.(chan int))

	default:
		fmt.Printf("Type is %T, value is %v\n", x, x)
	}
}
