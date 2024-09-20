package main

import "fmt"

func main() {
    
    isActive := true

    
    var status int
    if isActive {
        status = 1
    } else {
        status = 0
    }

 
    fmt.Println("Status :", status)

    status = 0
    if isActive {
        status = 1
    }

    fmt.Println("Status après conversion explicite :", status)
}
