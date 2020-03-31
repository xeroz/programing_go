package main

import "fmt"


func main()  {
	fmt.Println("array")

	var x [5]int
	fmt.Println(x)
	x[0] = 42
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	x1 := []int{4, 5, 6, 7, 8, 9}

	fmt.Println(x1)
	x1 = append(x1, 1, 2, 3)
	fmt.Println(x1)

	x2 := []int{4, 5, 6, 7, 8, 9}

	x1 = append(x1, x2...)
	fmt.Println(x1, "wwwwwwwwwww")

	fmt.Println(x1[1:], "ssssssss")
	fmt.Println(x1[1:3])

	for index, value := range x1{
		fmt.Println(index, value)
	}

	jb := []string{"James", "Bond", "Chocolate"}
	mp := []string{"James", "Bond", "Chocolate"}

	xp := [][]string{jb, mp}

	fmt.Println(xp)

	m := map[string]int{
		"james": 32,
		"james2": 32,
		"james3": 32,
	}

	fmt.Println(m)

}
