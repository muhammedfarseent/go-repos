package main

import "fmt"

func main (){

	n := 10  
    first, second := 0, 1 

    fmt.Println("Fibonacci Series:")

    for i := 0; i < n; i++ {
        fmt.Print(first, " ")

        next := first + second
        first = second
        second = next
}
}
