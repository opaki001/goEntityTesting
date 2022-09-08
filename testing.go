package main

import (
    "fmt"

    "github.com/opaki001/goEntityCache"
)

func main() {
    // Get a greeting message and print it.
    message := cacheMain.Hello("Gladys")
    fmt.Println(message)
}