package main

import "fmt"

func main() {
	var h, m, s, total int
	fmt.Println("donner le nombre dheures : ")
	fmt.Scanln(&h)
	fmt.Println("donner le nombre de minutes : ")
	fmt.Scanln(&m)
	fmt.Println("donner le nombre de second : ")
	fmt.Scanln(&s)
	total = h*3600 + m*60 + s

	Println("le total est : ", total)
}
