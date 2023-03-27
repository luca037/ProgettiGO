package main

import (
    "fmt"
    "time"
    "sync"
    "math/rand"
)

func simulateMarketData(wg *sync.WaitGroup, eur_usd, gbp_usd, jpy_usd chan<- float32, sec int) {
    for i := 0; i < sec; i++ {
        eu := randFloat32(1.0, 1.5)
        eur_usd <- eu
        fmt.Println("eu ->", eu)
        gu := randFloat32(1.0, 1.5)
        gbp_usd <- gu
        fmt.Println("gu ->", gu)
        ju := randFloat32(0.006, 0.009)
        jpy_usd <- ju
        fmt.Println("ju ->", ju)
        time.Sleep(time.Second)
    }
    close(eur_usd)
    close(gbp_usd)
    close(jpy_usd)
    wg.Done()
}

func selectPair(eur_usd, gbp_usd, jpy_usd <-chan float32) {
    for {
        select {
        case x := <- eur_usd:
            if x > 1.20 {
                time.Sleep(4 * time.Second)
                fmt.Println("venduti eur_usd ->", x)
            }
        case x := <- gbp_usd:
            if x < 1.35 {
                time.Sleep(3 * time.Second)
                fmt.Println("acquistati gbp_usd ->", x)
            }
        case x := <- jpy_usd:
            if x < 0.0085 {
                time.Sleep(3 * time.Second)
                fmt.Println("acquistati jpy_usd ->", x)
            }
        }
    }
}

func randFloat32(min, max float32) float32 {
    return min + rand.Float32() * (max - min)
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var wg sync.WaitGroup

    eur_usd := make(chan float32)
    gbp_usd := make(chan float32)
    jpy_usd := make(chan float32)

    wg.Add(1)
    go simulateMarketData(&wg, eur_usd, gbp_usd, jpy_usd, 10)
    go selectPair(eur_usd, gbp_usd, jpy_usd)
    wg.Wait()
}

// da fixare: con i channel buffer si ha che i valori non sono mai aggiornati,
// gli acquisti e le vendite non vengono effettuate con i valori più recenti
