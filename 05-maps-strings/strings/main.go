package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "Hello, Go!"

    fmt.Println(strings.ToUpper(text))
    fmt.Println(strings.Contains(text, "Go"))
}
