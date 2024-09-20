package main

import (
    "fmt"
    "strconv"
)

func main() {

    str := "123"
    number, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Erreur lors de la conversion de string vers int :", err)
    } else {
        fmt.Println("Conversion de string vers int :", number)
    }

   
    num := 456
    strNum := strconv.Itoa(num)
    fmt.Println("Conversion de int vers string :", strNum)
}
