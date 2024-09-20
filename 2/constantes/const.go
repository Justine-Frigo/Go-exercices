package main

import "fmt"

func main() {
   
    const Pi float64 = 3.14159


    var radius float64 = 4.5

   
    circonférence := 2 * Pi * radius

   
    fmt.Printf("La circonférence du cercle est de : %.2f\n", circonférence)
}
