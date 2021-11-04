package main

import "fmt"

func main() {
	var x, va int
	fmt.Printf("Donnez un nombre :")
	fmt.Scanln(&x)
	va = x
	if x < 0 {
		va = -x
	}
	fmt.Println("La valeur absolue est", va)
}
