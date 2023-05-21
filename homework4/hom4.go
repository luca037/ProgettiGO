package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// durata simulazione (in secondi)
const simulationDuration = 60

// soglie sopra/sotto le quali vendere/acquistare
const sellTresholdEurUsd float32 = 1.20
const buyTresholdGbpUsd float32 = 1.35
const buyTresholdJpyUsd float32 = 0.0085

type marketCurrencies struct {
	eur_usd, gbp_usd, jpy_usd chan float32
}

// genera un cambio di valuta randomico nel range [min, max] e lo invia nel
// canale passato; se il channel non è vuoto, scarta l'ultimo valore inserito
// per mantenere la valuta aggiornata; scrive il valore generato in res
func generateCurrencie(wg *sync.WaitGroup, ch chan float32, min, max float32, res *float32) {
	rnd := randFloat32(min, max)
	if len(ch) != 0 { <-ch }
	if res != nil { *res = rnd }
	ch <- rnd
	wg.Done()
}

// random number generator nel range [min, max]
func randFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// simula l'andamento di un mercato per tot secondi, le valute vengono aggiornate
// ogni secondo
func simulateMarketData(wg *sync.WaitGroup, curr *marketCurrencies, sec int, done *atomic.Bool) {
	var senders sync.WaitGroup
	var e_u, g_u, j_u float32  // valute correnti

    // inizio simulazione
	for i := 0; i < sec; i++ {
		senders.Add(3)
		go generateCurrencie(&senders, curr.eur_usd, 1.0, 1.5, &e_u)
		go generateCurrencie(&senders, curr.gbp_usd, 1.0, 1.5, &g_u)
		go generateCurrencie(&senders, curr.jpy_usd, 0.006, 0.009, &j_u)
		senders.Wait()
		time.Sleep(time.Second)
		log.Printf("\tcambi valute correnti: eu = %v, gu = %v, ju = %v", e_u, g_u, j_u)
	}

	// chiudo la ricezione
	done.Store(true)
	close(curr.eur_usd)
	close(curr.gbp_usd)
	close(curr.jpy_usd)
	wg.Done()
}

// algoritmo che cattura le variazioni di prezzo e decide se vendere o acquistare
func selectPair(wg *sync.WaitGroup, curr *marketCurrencies, done *atomic.Bool) {
	for !done.Load() {
		select {
		case x := <-curr.eur_usd:
			if x > sellTresholdEurUsd {
				time.Sleep(4 * time.Second)
				log.Println("VENDUTI eur_usd dal valore di ", x)
			}
		case x := <-curr.gbp_usd:
			if x < buyTresholdGbpUsd {
				time.Sleep(3 * time.Second)
				log.Println("ACQUISTATI gbp_usd dal valore di", x)
			}
		case x := <-curr.jpy_usd:
			if x < buyTresholdJpyUsd {
				time.Sleep(3 * time.Second)
				log.Println("ACQUISTATI jpy_usd dal valore di", x)
			}
		}
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano()) // per la randomicità

	var done atomic.Bool // per far terminare il thread selectPair
	done.Store(false)

	curr := marketCurrencies{
		eur_usd: make(chan float32, 1),
		gbp_usd: make(chan float32, 1),
		jpy_usd: make(chan float32, 1),
	}

	fmt.Println("#### INIZIO SIMULAZIONE ####")
	var wg sync.WaitGroup
	wg.Add(2)
	go simulateMarketData(&wg, &curr, simulationDuration, &done)
	go selectPair(&wg, &curr, &done)
	wg.Wait()
	fmt.Println("#### FINE SIMULAZIONE ####")
}

// da fixare:
// simulare un possibile guadagno

// fatti:
// come chiudere correttamente la goroutine del select;
// con i channel buffer si ha che i valori non sono mai aggiornati,
// gli acquisti e le vendite non vengono effettuate con i valori più recenti;
// controllare se questa variante funziona;
// aggiungere un metodo per stampare le valute correnti;
// camibiare nome struct e variabile 'valute'
