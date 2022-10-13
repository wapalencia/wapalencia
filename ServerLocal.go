package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Cancion de amor")

	openSource := "Cancion de amor"
	s := strings.Split(openSource, "")

	for i := 0; i < len(openSource); i++ {
		fmt.Println(s[i])
	}
	fmt.Println("for inverso.")

	num := []int{1, 2, 3, 4, 5}

	fmt.Println(len(num))
	for i := 4; i >= 0; i-- {
		fmt.Println(num[i] * -1)
	}

	for i := -10 * -1; i >= 0; i-- {
		fmt.Println("prueba for inverso. -- ", i*-1)
	}

}
