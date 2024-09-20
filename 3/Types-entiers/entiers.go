package main

import "fmt"

func main() {
    var a int8 = 10
    var b int16 = 1000
    var c int32 = 100000
    var d int64 = 10000000000

    fmt.Println("int8 :", a)
    fmt.Println("int16 :", b)
    fmt.Println("int32 :", c)
    fmt.Println("int64 :", d)

    sum := int32(b) + c

    fmt.Println("Somme de int16 et int32 (avec conversion) :", sum)
}
