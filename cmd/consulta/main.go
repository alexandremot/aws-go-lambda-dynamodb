package main

import (
    "fmt"
    "internal/consulta"
)

func main() {
    message := consulta.Hello("World")
    fmt.Println(message)
}
