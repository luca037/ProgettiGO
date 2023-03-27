package main

import (
    "fmt"
    "sync"
    "time"
)

type cake struct {
    name string
    isCooked bool
    isGarnished bool
    isDecorated bool
}

// cucina le torte passate
func cook(wg *sync.WaitGroup, cooked chan<- *cake, cakes []cake) {
    for i := range(cakes) {
        time.Sleep(time.Second)
        cakes[i].isCooked = true
        fmt.Printf("%s -> cucinata\n", cakes[i].name)
        cooked <- &cakes[i]
    }
    close(cooked)
    wg.Done()
}

// decora le torte passate da cook
func garnish(wg *sync.WaitGroup, cooked <-chan *cake, toDecorate chan<- *cake) {
    for c := range(cooked) {
        time.Sleep(2 * time.Second)
        c.isGarnished = true
        fmt.Printf("%s -> guarnita \n", c.name)
        toDecorate <- c
    }
    close(toDecorate)
    wg.Done()
}

// decora le torte passate da garnish
func decorate(wg *sync.WaitGroup, toDecorate <-chan *cake) {
    for c := range(toDecorate) {
        time.Sleep(4 * time.Second)
        c.isDecorated = true
        fmt.Printf("%s -> decorata\n", c.name)
    }
    wg.Done()
}

func main() {
    // alloco 5 torte
    cakes := make([]cake, 5)
    for i := range(cakes) {
        cakes[i].name = string('A' + i)
    }

    cooked := make(chan *cake, 2) // torte cucinate
    toDecorate := make(chan *cake, 2) // torte da decorare

    var wg sync.WaitGroup
    wg.Add(3)
    go cook(&wg, cooked, cakes)
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
