package main

import (
    "fmt"
    "sync"
    "math/rand"
    "time"
)

type client struct {
    name string
}

type vehicle struct {
    kind string
}

// nolleggio auto
func rent(wg *sync.WaitGroup, data chan<- int, cl client) {
    var v vehicle
    rnd := rand.Intn(3)
    switch rnd {
    case 0:
        v = vehicle{"SUV"}
    case 1:
        v = vehicle{"Berlina"}
    case 2:
        v = vehicle{"Station Wagon"}
    }
    data <- rnd
    fmt.Println(cl.name, "has rent a car of type", v.kind)
    wg.Done()
}

// conteggio auto nolleggiete per tipo e stampa resoconto
func accountant(wg *sync.WaitGroup, data <-chan int) {
    s, b, sw := 0, 0, 0
    for i := range(data) { // aspetta la chiusura del channel
        if i == 0 {
            s++
        } else if i == 1 {
            b++
        } else {
            sw++
        }
    }
    fmt.Println("numero di SUV =", s)
    fmt.Println("numero di Berline =", b)
    fmt.Println("numero di Station Wagon=", sw)
    wg.Done()
}

func main() {
    rand.Seed(time.Now().UnixNano()) // per la randomicità

    var wgs sync.WaitGroup // for senders
    var wgr sync.WaitGroup // for reciver
    data := make(chan int, 10) // per i nolleggi

    wgr.Add(1)
    go accountant(&wgr, data)

    for i := 0; i < 10; i++ {
        wgs.Add(1)
        c := client{"marco"}
        go rent(&wgs, data, c)
    }
    wgs.Wait()
    close(data)
    wgr.Wait()
}
