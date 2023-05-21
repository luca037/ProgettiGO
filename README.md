# Progetti in Go
Progetti (homework) assegnati da scrivere in Go.

## Homework 1 - conta caratteri in una stringa
Il programma chiede in input all'utente una stringa e un carattere, in output verrà restituito il numero di volte in cui appare tale carattere nella parola.

Per ogni carattere della stringa viene lanciata una goroutine. Ogni goroutine effettua un semplice confronto tra due caratteri e inserisce in un channel il risultato.
Il numero di volte in cui appare il carattere specificato all'interno della parola fornita corrsiponde alla lunghezza del channel.

## Homework 2 - car rental
Il programma genera una lista di clienti e associa casualmente ad ognuno di essi un tipo di veicolo tra tre tipologie: SUV, Berlina e Station Wagon. Successivamente il programma conta quanti veicoli sono stati noleggiati per ciascuna tipologia e stampa un resoconto finale.

Al suo avvio, il programma genera una lista di 10 clienti. Successivamente, il programma avvia una goroutine che gestisce il noleggio dei veicoli e una che che conteggia i veicoli noleggiati per tipologia. La goroutine che gestisce il noleggio lancia, per ciascun cliente, una goroutine che associa casualmente un tipo di veicolo tra tre tipologie. Quando un veicolo viene noleggiato, il tipo di veicolo viene inviato ad un channel. La goroutine che conta i veicoli noleggiati estrae le tipologie di veicoli dal channel ed effettua il conteggio finché il channel non viene chiuso. Infine viene stampato un resoconto.

## Homework 3 - produzione di torte
Il programma simula la preparazione di 5 torte in tre fasi:
1. cottura: ogni torta richiede un secondo;
2. guarnizione: ogni torta cotta viene guarnita in 2 secondi;
3. decorazione: ogni torta guarnita viene decorata in 4 secondi.

Il programma utilizza tre goroutine per gestire le tre fasi. La prima goroutine si occupa di cucinare le torte, la seconda di guarnirle e la terza di decorarle.

Per gestire la comunicazione tra le goroutine, il programma utilizza due channel. Il primo channel "cooked" viene utilizzato per passare le torte cotte alla seconda goroutine che le guarnirà. Il secondo channel "garnished" viene utilizzato per passare le torte guarnite alla terza goroutine che le decorerà.

## Homework 4 - simulazione mercato valutario
Questo programma simula l'andamento di un mercato valutario, generando delle variazioni di prezzo per le coppie di valute EUR/USD, GBP/USD e JPY/USD in maniera casuale e inviandole attraverso appositi canali.

Una goroutine si occupa di generare le variazioni di prezzo per le coppie di valute, un'altra goroutine si occupa di catturare tali variazioni e di acquistare/vendere le valute in base ad alcuni criteri specificati.
