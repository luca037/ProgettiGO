# Progetti in Go
Progetti (homework) assegnati da scrivere in Go.

## Homework 1 - conta caratteri in una stringa
Il programma chiede in input all'utente una stringa e un carattere, in output verrà restituito il numero di volte in cui appare tale carattere nella parola.

Per ogni carattere della stringa viene lanciata una goroutine. Ogni goroutine effettua un semplice confronto tra due caratteri e inserisce in un channel il risultato.
Il numero di volte in cui appare il carattere specificato all'interno della parola fornita corrsiponde alla lunghezza del channel.

Per eseguire il programma:
```bash
go run homework/hom1.go
```

Un possbile output:
```
inserire una parola:
thisisatest
inserire un carattere:
s
il carattere 's' appare 3 volta/e
```

## Homework 2 - car rental
Il programma genera una lista di clienti e associa casualmente ad ognuno di essi un tipo di veicolo tra tre tipologie: SUV, Berlina e Station Wagon. Successivamente il programma conta quanti veicoli sono stati noleggiati per ciascuna tipologia e stampa un resoconto finale.

Al suo avvio, il programma genera una lista di 10 clienti. Successivamente, il programma avvia una goroutine che gestisce il noleggio dei veicoli e una che che conteggia i veicoli noleggiati per tipologia. La goroutine che gestisce il noleggio lancia, per ciascun cliente, una goroutine che associa casualmente un tipo di veicolo tra tre tipologie. Quando un veicolo viene noleggiato, il tipo di veicolo viene inviato ad un channel. La goroutine che conta i veicoli noleggiati estrae le tipologie di veicoli dal channel ed effettua il conteggio finché il channel non viene chiuso. Infine viene stampato un resoconto.

Per eseguire il programma:
```bash
go run homework/hom2.go
```

Un possibile output:
```
name: J, ha noleggiato un veicolo di tipo: SUV
name: A, ha noleggiato un veicolo di tipo: SUV
name: G, ha noleggiato un veicolo di tipo: SUV
name: H, ha noleggiato un veicolo di tipo: Berlina
name: E, ha noleggiato un veicolo di tipo: Berlina
name: F, ha noleggiato un veicolo di tipo: SUV
name: B, ha noleggiato un veicolo di tipo: SUV
name: I, ha noleggiato un veicolo di tipo: SUV
name: C, ha noleggiato un veicolo di tipo: Station Wagon
name: D, ha noleggiato un veicolo di tipo: Berlina

#### RESOCONTO FINALE ####
numero totale di SUV = 6
numero totale di Berline = 3
numero totale di Station Wagon = 1
```

## Homework 3 - produzione di torte
Il programma simula la preparazione di 5 torte in tre fasi:
1. cottura: ogni torta richiede un secondo;
2. guarnizione: ogni torta cotta viene guarnita in 2 secondi;
3. decorazione: ogni torta guarnita viene decorata in 4 secondi.

Il programma utilizza tre goroutine per gestire le tre fasi. La prima goroutine si occupa di cucinare le torte, la seconda di guarnirle e la terza di decorarle.

Per gestire la comunicazione tra le goroutine, il programma utilizza due channel. Il primo channel "cooked" viene utilizzato per passare le torte cotte alla seconda goroutine che le guarnirà. Il secondo channel "garnished" viene utilizzato per passare le torte guarnite alla terza goroutine che le decorerà.

Per eseguire il programma:
```bash
go run homework/hom3.go
```

Un possibile output:
```
#### INIZIO PRODUZIONE ####
2023/05/25 10:40:53 Cuoco: Tiramisù è stata cucinata
2023/05/25 10:40:54 Cuoco: Sacher è stata cucinata
2023/05/25 10:40:55 Cuoco: Cheesecake è stata cucinata
2023/05/25 10:40:55 Guarnitore: Tiramisù è stata guarnita
2023/05/25 10:40:56 Cuoco: Crostata è stata cucinata
2023/05/25 10:40:57 Cuoco: Meringata è stata cucinata
2023/05/25 10:40:57 Guarnitore: Sacher è stata guarnita
2023/05/25 10:40:59 Decoratore: Tiramisù è stata decorata
2023/05/25 10:40:59 Guarnitore: Cheesecake è stata guarnita
2023/05/25 10:41:01 Guarnitore: Crostata è stata guarnita
2023/05/25 10:41:03 Guarnitore: Meringata è stata guarnita
2023/05/25 10:41:03 Decoratore: Sacher è stata decorata
2023/05/25 10:41:07 Decoratore: Cheesecake è stata decorata
2023/05/25 10:41:11 Decoratore: Crostata è stata decorata
2023/05/25 10:41:15 Decoratore: Meringata è stata decorata
#### FINE PRODUZIONE ####
```

## Homework 4 - simulazione mercato valutario
Questo programma simula l'andamento di un mercato valutario, generando delle variazioni di prezzo per le coppie di valute EUR/USD, GBP/USD e JPY/USD in maniera casuale e inviandole attraverso appositi canali.

Una goroutine si occupa di generare le variazioni di prezzo per le coppie di valute, un'altra goroutine si occupa di catturare tali variazioni e di acquistare/vendere le valute in base ad alcuni criteri specificati.

Per eseguire il programma:
```bash
go run homework/hom4.go
```

Un possibile output:
```
#### INIZIO SIMULAZIONE ####
2023/05/25 09:48:14 	cambi valute correnti: eu = 1.1375136, gu = 1.3560054, ju = 0.0061043403
2023/05/25 09:48:15 	cambi valute correnti: eu = 1.2106149, gu = 1.335802, ju = 0.0064044604
2023/05/25 09:48:16 ACQUISTATI jpy_usd dal valore di 0.0061043403
2023/05/25 09:48:16 	cambi valute correnti: eu = 1.3550076, gu = 1.0719882, ju = 0.007801777
2023/05/25 09:48:17 	cambi valute correnti: eu = 1.3850348, gu = 1.4206336, ju = 0.0084509775
2023/05/25 09:48:18 	cambi valute correnti: eu = 1.4633, gu = 1.3954339, ju = 0.0075907717
2023/05/25 09:48:19 ACQUISTATI jpy_usd dal valore di 0.007801777
2023/05/25 09:48:19 	cambi valute correnti: eu = 1.4994874, gu = 1.2894232, ju = 0.007982291
2023/05/25 09:48:20 	cambi valute correnti: eu = 1.1411515, gu = 1.0407734, ju = 0.007857163
2023/05/25 09:48:21 	cambi valute correnti: eu = 1.3348676, gu = 1.2869891, ju = 0.0061860853
2023/05/25 09:48:22 ACQUISTATI gbp_usd dal valore di 1.2894232
2023/05/25 09:48:22 	cambi valute correnti: eu = 1.2166119, gu = 1.4125242, ju = 0.007873127
2023/05/25 09:48:23 	cambi valute correnti: eu = 1.3648001, gu = 1.0271416, ju = 0.008059705
2023/05/25 09:48:24 	cambi valute correnti: eu = 1.1982994, gu = 1.2963296, ju = 0.0068917326
2023/05/25 09:48:25 	cambi valute correnti: eu = 1.0115014, gu = 1.158614, ju = 0.007199998
2023/05/25 09:48:26 VENDUTI eur_usd dal valore di  1.2166119
2023/05/25 09:48:26 	cambi valute correnti: eu = 1.4801215, gu = 1.0715704, ju = 0.00812907
2023/05/25 09:48:27 	cambi valute correnti: eu = 1.4198802, gu = 1.0237627, ju = 0.007595646
2023/05/25 09:48:28 	cambi valute correnti: eu = 1.3826247, gu = 1.2067677, ju = 0.0066298395
2023/05/25 09:48:29 ACQUISTATI jpy_usd dal valore di 0.00812907
2023/05/25 09:48:29 	cambi valute correnti: eu = 1.2944713, gu = 1.3255829, ju = 0.0075440365
2023/05/25 09:48:30 	cambi valute correnti: eu = 1.4062635, gu = 1.3500912, ju = 0.0063363602
2023/05/25 09:48:31 	cambi valute correnti: eu = 1.3048805, gu = 1.3441882, ju = 0.007960531
2023/05/25 09:48:32 	cambi valute correnti: eu = 1.181304, gu = 1.0756094, ju = 0.006289613
2023/05/25 09:48:33 VENDUTI eur_usd dal valore di  1.2944713
2023/05/25 09:48:33 	cambi valute correnti: eu = 1.2855392, gu = 1.1201105, ju = 0.006114744
2023/05/25 09:48:34 	cambi valute correnti: eu = 1.1169662, gu = 1.4319285, ju = 0.00842892
2023/05/25 09:48:35 	cambi valute correnti: eu = 1.1751478, gu = 1.4589779, ju = 0.008282999
2023/05/25 09:48:36 ACQUISTATI gbp_usd dal valore di 1.1201105
2023/05/25 09:48:36 	cambi valute correnti: eu = 1.1672595, gu = 1.1315778, ju = 0.0076555945
2023/05/25 09:48:37 	cambi valute correnti: eu = 1.4999561, gu = 1.3567895, ju = 0.0063115424
2023/05/25 09:48:38 	cambi valute correnti: eu = 1.3593135, gu = 1.0084704, ju = 0.007583429
2023/05/25 09:48:39 ACQUISTATI gbp_usd dal valore di 1.1315778
2023/05/25 09:48:39 	cambi valute correnti: eu = 1.1321442, gu = 1.1209955, ju = 0.008232256
2023/05/25 09:48:40 	cambi valute correnti: eu = 1.3970907, gu = 1.3771405, ju = 0.008250034
2023/05/25 09:48:41 	cambi valute correnti: eu = 1.0478239, gu = 1.0050796, ju = 0.00749947
2023/05/25 09:48:42 ACQUISTATI gbp_usd dal valore di 1.1209955
2023/05/25 09:48:42 	cambi valute correnti: eu = 1.0629406, gu = 1.1632893, ju = 0.007295221
2023/05/25 09:48:43 	cambi valute correnti: eu = 1.4281111, gu = 1.4948357, ju = 0.0079836985
2023/05/25 09:48:44 	cambi valute correnti: eu = 1.339742, gu = 1.428214, ju = 0.008964181
2023/05/25 09:48:45 ACQUISTATI gbp_usd dal valore di 1.1632893
2023/05/25 09:48:45 	cambi valute correnti: eu = 1.3637974, gu = 1.3510185, ju = 0.0064040506
2023/05/25 09:48:46 	cambi valute correnti: eu = 1.3146775, gu = 1.3555107, ju = 0.007743306
2023/05/25 09:48:47 	cambi valute correnti: eu = 1.1423194, gu = 1.0362158, ju = 0.00873218
2023/05/25 09:48:48 	cambi valute correnti: eu = 1.1404585, gu = 1.0926083, ju = 0.007940724
2023/05/25 09:48:49 VENDUTI eur_usd dal valore di  1.3637974
2023/05/25 09:48:49 	cambi valute correnti: eu = 1.1720312, gu = 1.3889098, ju = 0.0062687765
2023/05/25 09:48:50 	cambi valute correnti: eu = 1.2184954, gu = 1.1632792, ju = 0.0071361503
2023/05/25 09:48:51 	cambi valute correnti: eu = 1.0952806, gu = 1.0398113, ju = 0.0061183465
2023/05/25 09:48:52 ACQUISTATI jpy_usd dal valore di 0.0062687765
2023/05/25 09:48:52 	cambi valute correnti: eu = 1.4562006, gu = 1.1920387, ju = 0.0068379464
2023/05/25 09:48:53 	cambi valute correnti: eu = 1.184744, gu = 1.0338442, ju = 0.006085027
2023/05/25 09:48:54 	cambi valute correnti: eu = 1.1405873, gu = 1.1415616, ju = 0.008710688
2023/05/25 09:48:55 ACQUISTATI jpy_usd dal valore di 0.0068379464
2023/05/25 09:48:55 	cambi valute correnti: eu = 1.0178804, gu = 1.0281926, ju = 0.007255404
2023/05/25 09:48:56 	cambi valute correnti: eu = 1.1178433, gu = 1.2029566, ju = 0.008434286
2023/05/25 09:48:57 	cambi valute correnti: eu = 1.2862812, gu = 1.464122, ju = 0.007617614
2023/05/25 09:48:58 ACQUISTATI jpy_usd dal valore di 0.007255404
2023/05/25 09:48:58 	cambi valute correnti: eu = 1.2668915, gu = 1.3140037, ju = 0.0073556066
2023/05/25 09:48:59 	cambi valute correnti: eu = 1.1536313, gu = 1.0255777, ju = 0.0068495874
2023/05/25 09:49:00 	cambi valute correnti: eu = 1.4771202, gu = 1.2436935, ju = 0.006812117
2023/05/25 09:49:01 ACQUISTATI jpy_usd dal valore di 0.0073556066
2023/05/25 09:49:01 	cambi valute correnti: eu = 1.4524262, gu = 1.4157887, ju = 0.0062013073
2023/05/25 09:49:02 	cambi valute correnti: eu = 1.1467671, gu = 1.3443596, ju = 0.006425267
2023/05/25 09:49:03 	cambi valute correnti: eu = 1.2351564, gu = 1.1011885, ju = 0.0060503054
2023/05/25 09:49:04 	cambi valute correnti: eu = 1.4543861, gu = 1.35691, ju = 0.0069618234
2023/05/25 09:49:05 VENDUTI eur_usd dal valore di  1.4524262
2023/05/25 09:49:05 	cambi valute correnti: eu = 1.4398787, gu = 1.3728697, ju = 0.0063622263
2023/05/25 09:49:06 	cambi valute correnti: eu = 1.3614022, gu = 1.1110493, ju = 0.007017688
2023/05/25 09:49:07 	cambi valute correnti: eu = 1.3958352, gu = 1.4312605, ju = 0.008918559
2023/05/25 09:49:08 ACQUISTATI jpy_usd dal valore di 0.0063622263
2023/05/25 09:49:08 	cambi valute correnti: eu = 1.2925277, gu = 1.1583073, ju = 0.0060373805
2023/05/25 09:49:09 	cambi valute correnti: eu = 1.0587523, gu = 1.0318178, ju = 0.0073924763
2023/05/25 09:49:10 	cambi valute correnti: eu = 1.0094571, gu = 1.0942079, ju = 0.006774799
2023/05/25 09:49:11 	cambi valute correnti: eu = 1.2334626, gu = 1.2981997, ju = 0.0060218875
2023/05/25 09:49:12 VENDUTI eur_usd dal valore di  1.2925277
2023/05/25 09:49:12 	cambi valute correnti: eu = 1.4424216, gu = 1.1128485, ju = 0.0082476
2023/05/25 09:49:13 	cambi valute correnti: eu = 1.2016089, gu = 1.1921993, ju = 0.008551516
2023/05/25 09:49:15 ACQUISTATI gbp_usd dal valore di 1.1128485
#### FINE SIMULAZIONE ####
```
