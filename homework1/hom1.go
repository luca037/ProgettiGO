package main

import (
	"fmt"
	"sync"
)

// Conta quante volte un carattere è presente in una parola.
// word è la parola da analizzare.
// c è il carattere utilizzato per il conteggio.
// res è il canale in cui viene inviato il risultato.
func CharCount(word string, c byte, res chan<- int) {
	var wg sync.WaitGroup
	ch := make(chan byte, len(word))

	// controllo caratteri in parallelo
	for i := range word {
		wg.Add(1)
		go func(x byte) {
			defer wg.Done()
			if x == c {
				ch <- 1
			}
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
	go CharCount(word, char[0], res)
	count := <-res

	fmt.Printf("il carattere '%c' appare %d volta/e\n", char[0], count)
}
