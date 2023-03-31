package main

import (
    "fmt"
    "sync"
)

// ritorna il numero di corrispondenze trovate
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
    var word string
    fmt.Println("inserire una parola:")
    fmt.Scanln(&word)

    var char []byte
    fmt.Println("inserire un carattere:")
    fmt.Scanln(&char)

    res := make(chan int)

    go char_count(word, char[0], res)
    count := <- res

    fmt.Printf("il appare '%c' appare %d volta/e\n", char[0], count)
}

// cose che si possono aggiungere:
// input da utente 
