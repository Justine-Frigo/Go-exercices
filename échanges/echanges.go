package main

import "fmt"

func main() {
	x := 5
	y := 10

	fmt.Println("Avant échange :")
    fmt.Println("x =", x)
    fmt.Println("y =", y)

	x, y = y, x

    fmt.Println("Après échange :")
    fmt.Println("x =", x)
    fmt.Println("y =", y)
}