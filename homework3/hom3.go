package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// Struttura che rappresenta una torta.
type Cake struct {
	Name        string // nome torta
	IsCooked    bool   // indica se è stata cucinata
	IsGarnished bool   // indica se è stata farcita
	IsDecorated bool   // indica se è stata decorata
}

// Mansione che svolge il cuoco: cucina le torte passate.
// wg utilizzato per sincronizzare la go routine.
// cooked è il canale in cui vengono inserite le torte cucinate.
// cakes sono le torte da cucinare.
func Cook(wg *sync.WaitGroup, cooked chan<- *Cake, cakes []Cake) {
	defer wg.Done()
	for i := range cakes {
		time.Sleep(time.Second)
		cakes[i].IsCooked = true
		cooked <- &cakes[i]
		log.Printf("Cuoco: %s è stata cucinata\n", cakes[i].Name)
	}
	close(cooked)
}

// Mansione che svolge il farcitore.
// wg utilizzato per sincronizzare la go routine.
// cooked è il canale da cui preleva le torte cucinate.
// garnished è il canale in cui vengono inserite le torte decorate.
func Garnish(wg *sync.WaitGroup, cooked <-chan *Cake, garnished chan<- *Cake) {
	defer wg.Done()
	for cake := range cooked {
		time.Sleep(4 * time.Second)
		cake.IsGarnished = true
		garnished <- cake
		log.Printf("Guarnitore: %s è stata guarnita\n", cake.Name)
	}
	close(garnished)
}

// Mansione che svolge il decoratore.
// wg utilizzato per sincronizzare la go routine.
// garnished è il cananle da cui preleva le torte da decorare.
func Decorate(wg *sync.WaitGroup, garnished <-chan *Cake) {
	defer wg.Done()
	for cake := range garnished {
		time.Sleep(8 * time.Second)
		cake.IsDecorated = true
		log.Printf("Decoratore: %s è stata decorata\n", cake.Name)
	}
}

func main() {
	// alloco le 5 torte che devo cucinare.
	cakes := make([]Cake, 5)
	for i, name := range []string{"Tiramisù", "Sacher", "Cheesecake", "Crostata", "Meringata"} {
		cakes[i].Name = name
	}

	cooked := make(chan *Cake, 2)    // torte cucinate
	garnished := make(chan *Cake, 2) // torte guarnite

	fmt.Println("#### INIZIO PRODUZIONE ####")
	var wg sync.WaitGroup
	wg.Add(3)
	go Cook(&wg, cooked, cakes)
	go Garnish(&wg, cooked, garnished)
	go Decorate(&wg, garnished)
	wg.Wait()
	fmt.Println("#### FINE PRODUZIONE ####")
}
