package main

import (
    "fmt"
    "sync"
    "time"
)

// cucina n torte
func cook(wg *sync.WaitGroup, cooked chan<- byte, n int) {
    var b byte = 'A'
    for i := 0; i < n; i++ {
        time.Sleep(time.Second)
        fmt.Printf("%c -> cucinata\n", b)
        cooked <- b
        b++
    }
    close(cooked)
    wg.Done()
}

// decora le torte passate da cook
func garnish(wg *sync.WaitGroup, cooked <-chan byte, toDecorate chan<- byte) {
    for c := range(cooked) {
        time.Sleep(2 * time.Second)
        fmt.Printf("%c -> guarnita \n", c)
        toDecorate <- c
    }
    close(toDecorate)
    wg.Done()
}

// decora le torte passate da garnish
func decorate(wg *sync.WaitGroup, toDecorate <-chan byte) {
    for c := range(toDecorate) {
        time.Sleep(4 * time.Second)
        fmt.Printf("%c -> decorata\n", c)
    }
    wg.Done()
}

func main() {
    cooked := make(chan byte, 2) // torte cucinate
    toDecorate := make(chan byte, 2) // torte da decorare

    var wg sync.WaitGroup
    wg.Add(3)
    go cook(&wg, cooked, 5)
    go garnish(&wg, cooked, toDecorate)
    go decorate(&wg, toDecorate)
    wg.Wait()
}

/*
cose che si possono aggiungere:
    struct cake con 3 campi booleani che identificano le 3 fasi e nome della torta
    slice di cake stile homework 2 nel main
    stampa non in sequenza ma stile tabella
*/
