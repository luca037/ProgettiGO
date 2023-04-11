package main

import (
    "fmt"
    "log"
    "time"
    "sync"
    "sync/atomic"
    "math/rand"
)

// genera un cambio di valuta randomico nel range [min, max] e lo invia nel canale
// se il canale non è vuoto, viene aggiornato il valore contenuto
func generateValute(wg *sync.WaitGroup, ch chan float32, min, max float32, res *float32) {
    if len(ch) != 0 { // se il canale non è vuoto
        <- ch // scarto il valore presente (per mantenere i valori sempre aggiornati)
    } 
    rnd := randFloat32(min, max)
    *res = rnd // per poter stampare il valore delle valute correnti
    ch <- rnd 
    wg.Done()
}

// random number generator nel range [min, max]
func randFloat32(min, max float32) float32 {
    return min + rand.Float32() * (max - min)
}

// simula l'andamento di un mercato per tot secondi;
// le valute vengono aggiornate ogni secondo
func simulateMarketData(wg *sync.WaitGroup, eur_usd, gbp_usd, jpy_usd chan float32, sec int, done *atomic.Bool) {
    var senders sync.WaitGroup
    var eu, gu, ju float32
    for i := 0; i < sec; i++ { // inizio simulazione
        senders.Add(3)
        go generateValute(&senders, eur_usd, 1.0, 1.5, &eu)
        go generateValute(&senders, gbp_usd, 1.0, 1.5, &gu)
        go generateValute(&senders, jpy_usd, 0.006, 0.009, &ju)
        senders.Wait()
        time.Sleep(time.Second)
        log.Printf("\tcambi valute correnti: eu = %v, gu = %v, ju = %v", eu, gu, ju)
    }
    done.Store(true) // faccio terminare selectPair
    close(eur_usd)
    close(gbp_usd)
    close(jpy_usd)
    wg.Done()
}

// algoritmo che cattura le variazioni di prezzo e 
// decide se vendere o acquistare
func selectPair(wg *sync.WaitGroup, eur_usd, gbp_usd, jpy_usd <-chan float32, done *atomic.Bool) {
    for !done.Load() {
        select {
        case x := <- eur_usd:
            if x > 1.20 {
                time.Sleep(4 * time.Second)
                log.Println("VENDUTI eur_usd dal valore di ", x)
            }
        case x := <- gbp_usd:
            if x < 1.35 {
                time.Sleep(3 * time.Second)
                log.Println("ACQUISTATI gbp_usd dal valore di", x)
            }
        case x := <- jpy_usd:
            if x < 0.0085 {
                time.Sleep(3 * time.Second)
                log.Println("ACQUISTATI jpy_usd dal valore di", x)
            }
        }
    }
    wg.Done()
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var done atomic.Bool // utilizzata per interrompere la ricezione di selectPair
    done.Store(false)

    eur_usd := make(chan float32, 1)
    gbp_usd := make(chan float32, 1)
    jpy_usd := make(chan float32, 1)

    fmt.Println("#### INIZIO SIMULAZIONE ####")
    var wg sync.WaitGroup
    wg.Add(2)
    go simulateMarketData(&wg, eur_usd, gbp_usd, jpy_usd, 60, &done)
    go selectPair(&wg, eur_usd, gbp_usd, jpy_usd, &done)
    wg.Wait()
    fmt.Println("#### FINE SIMULAZIONE ####")
}

// da fixare:
// pulire il codice

// fatti:
// come chiudere correttamente la goroutine del select;
// con i channel buffer si ha che i valori non sono mai aggiornati,
// gli acquisti e le vendite non vengono effettuate con i valori più recenti;
// controllare se questa variante funziona;
// aggiungere un metodo per stampare le valute correnti;
