package main

import "fmt"

var x bool

func main()  {

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Printf("primera iteracion %d segunda iteracion %d\n", i, j)
		}
	}

}

