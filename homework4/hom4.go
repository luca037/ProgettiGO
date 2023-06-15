package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Durata simulazione (in secondi).
const simulationDuration = 60

// Soglie sopra/sotto le quali vendere/acquistare.
const sellTresholdEurUsd float32 = 1.20
const buyTresholdGbpUsd float32 = 1.35
const buyTresholdJpyUsd float32 = 0.0085

// Valute di un mercato azionario.
type MarketCurrencies struct {
	EurUsd, GbpUsd, JpyUsd chan float32
}

// Genera un cambio di valuta randomico nel range [min, max] e lo invia nel
// canale passato; se il channel non è vuoto viene scartato il valore presente
// in modo da mantenere la valuta aggiornata.
// wg viene utilizzato per sincronizzare la go routine.
// ch è il cananle in cui viene inviato il valore generato.
// In res viene salvato il valore che è stato generato.
func generateCurrencie(wg *sync.WaitGroup, ch chan float32, min, max float32, res *float32) {
	defer wg.Done()

	rnd := randFloat32(min, max)

	if len(ch) != 0 {
		<-ch
	}

	if res != nil {
		*res = rnd
	}

	ch <- rnd
}

// Genera un numero random nel range [min, max].
func randFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// Simula l'andamento di un mercato per tot secondi, le valute vengono aggiornate
// con un valore randomico ad ogni secondo.
// wg viene utilizzato per sincronizzare la go routine.
// curr contiene le valute del mercato azionario.
// sec indica la durata della simulazione in secondi.
// done viene utilizzata per comunicare che la simulazione è terminata.
func SimulateMarketData(wg *sync.WaitGroup, curr *MarketCurrencies, sec int, done *atomic.Bool) {
	defer wg.Done()
	var senders sync.WaitGroup
	var e_u, g_u, j_u float32 // valute correnti (nomi abbreviati)

	// inizio simulazione
	for i := 0; i < sec; i++ {
		senders.Add(3)
		go generateCurrencie(&senders, curr.EurUsd, 1.0, 1.5, &e_u)
		go generateCurrencie(&senders, curr.GbpUsd, 1.0, 1.5, &g_u)
		go generateCurrencie(&senders, curr.JpyUsd, 0.006, 0.009, &j_u)
		senders.Wait()
		time.Sleep(time.Second)
		log.Printf("\tcambi valute correnti: EUR/USD = %v, GBP/USD = %v, JPY/USD = %v", e_u, g_u, j_u)
	}

	// chiudo la ricezione
	done.Store(true)
	close(curr.EurUsd)
	close(curr.GbpUsd)
	close(curr.JpyUsd)
}

// Cattura le variazioni di mercato delle valute.
// wg utilizzato per sincronizzare la go routine.
// curr contiene le valute del mercato soggette a variazioni.
// done viene utilizzata per far terminare la ricezione.
func SelectPair(wg *sync.WaitGroup, curr *MarketCurrencies, done *atomic.Bool) {
	defer wg.Done()

	for !done.Load() {
		select {
		case x := <-curr.EurUsd:
			if x > sellTresholdEurUsd {
				time.Sleep(4 * time.Second)
				log.Println("VENDUTI EUR/USD dal valore di ", x)
			}
		case x := <-curr.GbpUsd:
			if x < buyTresholdGbpUsd {
				time.Sleep(3 * time.Second)
				log.Println("ACQUISTATI GBP/USD dal valore di", x)
			}
		case x := <-curr.JpyUsd:
			if x < buyTresholdJpyUsd {
				time.Sleep(3 * time.Second)
				log.Println("ACQUISTATI JPY/USD dal valore di", x)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // per la randomicità

	var done atomic.Bool // per far terminare il thread selectPair
	done.Store(false)

	curr := MarketCurrencies{
		EurUsd: make(chan float32, 1),
		GbpUsd: make(chan float32, 1),
		JpyUsd: make(chan float32, 1),
	}

	fmt.Println("#### INIZIO SIMULAZIONE ####")
	var wg sync.WaitGroup
	wg.Add(2)
	go SimulateMarketData(&wg, &curr, simulationDuration, &done)
	go SelectPair(&wg, &curr, &done)
	wg.Wait()
	fmt.Println("#### FINE SIMULAZIONE ####")
}
