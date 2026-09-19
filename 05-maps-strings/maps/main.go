package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
    }

    fmt.Println(scores)
    fmt.Println("Alice:", scores["Alice"])
}
