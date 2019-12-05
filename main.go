package main

import (
	"fmt"
)

func main()  {

	var a string = "HelloWord!!!"

	fmt.Println(a)

	animal := Animal{"Tom", "black"}

	dog := Dog{ animal ,19}

	fmt.Print(animal.name)

	fmt.Print(dog.Animal.name)

	// ***
	dog.Animal.eat()
	dog.eat()

	//数组
	arr := []int{1,2,3}
	var i int

	for i = 0; i < 3 ; i++  {
		fmt.Print(arr[i])
	}
}

type Animal struct {
	name string
	color string
}

func (a Animal) eat()  {
	fmt.Printf("%s is eating\n",a.name)
}

type Dog struct {
	Animal
	weight int
}
