package main

import "fmt"

func main() {
   
    var greeting string = "Bonjour"
    var language string = "Go!"

    fmt.Println(greeting)
    fmt.Println(language)

    escapedString := "Voici une chaîne avec un retour chariot,\nun tabulateur,\tet des guillemets doubles : \"GoLang\"."

    fmt.Println(escapedString)
}
