package main

import (
    "fmt"
    "sync"
)

func char_count(word string, c byte, res chan<- int) {
    var wg sync.WaitGroup
    ch := make(chan byte, len(word))

    for i := range(word) {
        wg.Add(1)
        go func(x byte) {
            if x == c {
                ch <- 1
            }
            wg.Done()
        }(word[i])
    }

    wg.Wait()
    close(ch)
    res <- len(ch)
}

func main() {
    var word string = "aaaaaaaaaaaaabbbbbbbbcccccddddccccccfff"
    var char byte = 'c'

    res := make(chan int)

    go char_count(word, char, res)
    count := <- res

    fmt.Println("word =", word)
    fmt.Printf("character '%c' appears %d times\n", char, count)
}

// cose che si possono aggiungere:
// input da utente 
