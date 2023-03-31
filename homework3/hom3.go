package main

import (
    "sync"
    "time"
    "log"
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
        log.Printf("%s -> cucinata\n", cakes[i].name)
        cooked <- &cakes[i]
    }
    close(cooked)
    wg.Done()
}

// decora le torte passate da cook
func garnish(wg *sync.WaitGroup, cooked <-chan *cake, garnished chan<- *cake) {
    for c := range(cooked) {
        time.Sleep(2 * time.Second)
        c.isGarnished = true
        log.Printf("%s -> guarnita \n", c.name)
        garnished <- c
    }
    close(garnished)
    wg.Done()
}

// decora le torte passate da garnish
func decorate(wg *sync.WaitGroup, toDecorate <-chan *cake) {
    for c := range(toDecorate) {
        time.Sleep(4 * time.Second)
        c.isDecorated = true
        log.Printf("%s -> decorata\n", c.name)
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
    garnished := make(chan *cake, 2) // torte da decorare

    var wg sync.WaitGroup
    wg.Add(3)
    go cook(&wg, cooked, cakes)
    go garnish(&wg, cooked, garnished)
    go decorate(&wg, garnished)
    wg.Wait()
}

/* cose che si possono aggiungere:
 * output */
