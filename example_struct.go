package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func main()  {

	p1 := person{
		first: "James",
		last: "bond",
	}

	sa1 := secretAgent{
		person: p1,
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)

	fmt.Println(sa1.person.first)

}

