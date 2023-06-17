package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Conta quante volte un carattere è presente in una parola.
// word è la parola da analizzare.
// c è il carattere utilizzato per il conteggio.
// res è il canale in cui viene inviato il risultato del conteggio.
func CharCount(word string, c byte, res chan<- int) {
	var wg sync.WaitGroup
	count := make(chan byte, len(word))

	// controllo caratteri
	for i := range word {
		wg.Add(1)
		go func(x byte) {
			defer wg.Done()
			if x == c {
				count <- 1
			}
		}(word[i])
	}

	wg.Wait()
	close(count)

	res <- len(count)
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("inserire una parola:")
	word, _ := in.ReadString('\n')

	fmt.Println("inserire un carattere:")
	char, _ := in.ReadByte()

	res := make(chan int) // canale in cui viene passato il conteggio
	go CharCount(word, char, res)
	count := <-res

	fmt.Printf("il carattere '%c', appare %d volta/e\n", char, count)
}
