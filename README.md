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

## Homework 2 - noleggio automobili
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
2023/05/31 13:13:00 Cuoco: Tiramisù è stata cucinata
2023/05/31 13:13:01 Cuoco: Sacher è stata cucinata
2023/05/31 13:13:02 Cuoco: Cheesecake è stata cucinata
2023/05/31 13:13:03 Cuoco: Crostata è stata cucinata
2023/05/31 13:13:04 Guarnitore: Tiramisù è stata guarnita
2023/05/31 13:13:05 Cuoco: Meringata è stata cucinata
2023/05/31 13:13:08 Guarnitore: Sacher è stata guarnita
2023/05/31 13:13:12 Decoratore: Tiramisù è stata decorata
2023/05/31 13:13:12 Guarnitore: Cheesecake è stata guarnita
2023/05/31 13:13:16 Guarnitore: Crostata è stata guarnita
2023/05/31 13:13:20 Decoratore: Sacher è stata decorata
2023/05/31 13:13:20 Guarnitore: Meringata è stata guarnita
2023/05/31 13:13:28 Decoratore: Cheesecake è stata decorata
2023/05/31 13:13:36 Decoratore: Crostata è stata decorata
2023/05/31 13:13:44 Decoratore: Meringata è stata decorata
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
2023/05/28 17:16:23 	cambi valute correnti: EUR/USD = 1.090255, GBP/USD = 1.4743164, JPY/USD = 0.008893158
2023/05/28 17:16:24 	cambi valute correnti: EUR/USD = 1.237488, GBP/USD = 1.4857941, JPY/USD = 0.006895312
2023/05/28 17:16:25 	cambi valute correnti: EUR/USD = 1.0315063, GBP/USD = 1.3790376, JPY/USD = 0.008208969
2023/05/28 17:16:26 ACQUISTATI JPY/USD dal valore di 0.006895312
2023/05/28 17:16:26 	cambi valute correnti: EUR/USD = 1.1794821, GBP/USD = 1.151135, JPY/USD = 0.006642764
2023/05/28 17:16:27 	cambi valute correnti: EUR/USD = 1.1552612, GBP/USD = 1.3761103, JPY/USD = 0.007384676
2023/05/28 17:16:28 	cambi valute correnti: EUR/USD = 1.402172, GBP/USD = 1.0393479, JPY/USD = 0.008901599
2023/05/28 17:16:29 ACQUISTATI GBP/USD dal valore di 1.151135
2023/05/28 17:16:29 	cambi valute correnti: EUR/USD = 1.1276011, GBP/USD = 1.2371044, JPY/USD = 0.0069815367
2023/05/28 17:16:30 	cambi valute correnti: EUR/USD = 1.3131588, GBP/USD = 1.3325305, JPY/USD = 0.008038612
2023/05/28 17:16:31 	cambi valute correnti: EUR/USD = 1.482144, GBP/USD = 1.0112741, JPY/USD = 0.0072546685
2023/05/28 17:16:32 ACQUISTATI GBP/USD dal valore di 1.2371044
2023/05/28 17:16:32 	cambi valute correnti: EUR/USD = 1.2351931, GBP/USD = 1.1765357, JPY/USD = 0.008534676
2023/05/28 17:16:33 	cambi valute correnti: EUR/USD = 1.4945828, GBP/USD = 1.1504409, JPY/USD = 0.006278705
2023/05/28 17:16:34 	cambi valute correnti: EUR/USD = 1.3650668, GBP/USD = 1.3464973, JPY/USD = 0.008606613
2023/05/28 17:16:35 	cambi valute correnti: EUR/USD = 1.0021929, GBP/USD = 1.4970922, JPY/USD = 0.008541149
2023/05/28 17:16:36 VENDUTI EUR/USD dal valore di  1.2351931
2023/05/28 17:16:36 	cambi valute correnti: EUR/USD = 1.157764, GBP/USD = 1.4245199, JPY/USD = 0.0078030657
2023/05/28 17:16:37 	cambi valute correnti: EUR/USD = 1.3113414, GBP/USD = 1.1275012, JPY/USD = 0.007579696
2023/05/28 17:16:38 	cambi valute correnti: EUR/USD = 1.3999696, GBP/USD = 1.370509, JPY/USD = 0.007321366
2023/05/28 17:16:39 ACQUISTATI JPY/USD dal valore di 0.0078030657
2023/05/28 17:16:39 	cambi valute correnti: EUR/USD = 1.3341286, GBP/USD = 1.0479048, JPY/USD = 0.0070166457
2023/05/28 17:16:40 	cambi valute correnti: EUR/USD = 1.2257991, GBP/USD = 1.1473099, JPY/USD = 0.0072492273
2023/05/28 17:16:41 	cambi valute correnti: EUR/USD = 1.3766727, GBP/USD = 1.2012546, JPY/USD = 0.0071013644
2023/05/28 17:16:42 	cambi valute correnti: EUR/USD = 1.2570184, GBP/USD = 1.4235327, JPY/USD = 0.008791863
2023/05/28 17:16:43 VENDUTI EUR/USD dal valore di  1.3341286
2023/05/28 17:16:43 	cambi valute correnti: EUR/USD = 1.0744866, GBP/USD = 1.3924919, JPY/USD = 0.006564118
2023/05/28 17:16:44 	cambi valute correnti: EUR/USD = 1.2073885, GBP/USD = 1.0854852, JPY/USD = 0.0071486672
2023/05/28 17:16:45 	cambi valute correnti: EUR/USD = 1.4110394, GBP/USD = 1.1830978, JPY/USD = 0.0064241635
2023/05/28 17:16:46 ACQUISTATI JPY/USD dal valore di 0.006564118
2023/05/28 17:16:46 	cambi valute correnti: EUR/USD = 1.1708325, GBP/USD = 1.3056486, JPY/USD = 0.006407966
2023/05/28 17:16:47 	cambi valute correnti: EUR/USD = 1.4982027, GBP/USD = 1.4750392, JPY/USD = 0.0076837894
2023/05/28 17:16:48 	cambi valute correnti: EUR/USD = 1.3888762, GBP/USD = 1.0188625, JPY/USD = 0.0069191414
2023/05/28 17:16:49 ACQUISTATI JPY/USD dal valore di 0.006407966
2023/05/28 17:16:49 	cambi valute correnti: EUR/USD = 1.0986861, GBP/USD = 1.0482097, JPY/USD = 0.008523626
2023/05/28 17:16:50 	cambi valute correnti: EUR/USD = 1.1791203, GBP/USD = 1.113889, JPY/USD = 0.00709438
2023/05/28 17:16:51 	cambi valute correnti: EUR/USD = 1.2304968, GBP/USD = 1.170385, JPY/USD = 0.008257115
2023/05/28 17:16:52 ACQUISTATI GBP/USD dal valore di 1.0482097
2023/05/28 17:16:52 	cambi valute correnti: EUR/USD = 1.2795944, GBP/USD = 1.0292376, JPY/USD = 0.0089546265
2023/05/28 17:16:53 	cambi valute correnti: EUR/USD = 1.3914111, GBP/USD = 1.1808289, JPY/USD = 0.008152075
2023/05/28 17:16:54 	cambi valute correnti: EUR/USD = 1.4415332, GBP/USD = 1.3640448, JPY/USD = 0.007938985
2023/05/28 17:16:55 ACQUISTATI GBP/USD dal valore di 1.0292376
2023/05/28 17:16:55 	cambi valute correnti: EUR/USD = 1.1021636, GBP/USD = 1.0373666, JPY/USD = 0.008503059
2023/05/28 17:16:56 	cambi valute correnti: EUR/USD = 1.0640742, GBP/USD = 1.2637594, JPY/USD = 0.0062469225
2023/05/28 17:16:57 	cambi valute correnti: EUR/USD = 1.4595184, GBP/USD = 1.3394779, JPY/USD = 0.006463093
2023/05/28 17:16:58 ACQUISTATI GBP/USD dal valore di 1.0373666
2023/05/28 17:16:58 	cambi valute correnti: EUR/USD = 1.3339969, GBP/USD = 1.0632827, JPY/USD = 0.0075974315
2023/05/28 17:16:59 	cambi valute correnti: EUR/USD = 1.3748628, GBP/USD = 1.2397245, JPY/USD = 0.0064807627
2023/05/28 17:17:00 	cambi valute correnti: EUR/USD = 1.1265259, GBP/USD = 1.4988945, JPY/USD = 0.006678395
2023/05/28 17:17:01 ACQUISTATI JPY/USD dal valore di 0.0075974315
2023/05/28 17:17:01 	cambi valute correnti: EUR/USD = 1.360012, GBP/USD = 1.3278766, JPY/USD = 0.007559636
2023/05/28 17:17:02 	cambi valute correnti: EUR/USD = 1.4275522, GBP/USD = 1.3879255, JPY/USD = 0.007710783
2023/05/28 17:17:03 	cambi valute correnti: EUR/USD = 1.4252856, GBP/USD = 1.0999635, JPY/USD = 0.0063704294
2023/05/28 17:17:04 ACQUISTATI JPY/USD dal valore di 0.007559636
2023/05/28 17:17:04 	cambi valute correnti: EUR/USD = 1.402964, GBP/USD = 1.2821785, JPY/USD = 0.0064937556
2023/05/28 17:17:05 	cambi valute correnti: EUR/USD = 1.3376377, GBP/USD = 1.0567845, JPY/USD = 0.008561776
2023/05/28 17:17:06 	cambi valute correnti: EUR/USD = 1.4705496, GBP/USD = 1.0430784, JPY/USD = 0.006017899
2023/05/28 17:17:07 ACQUISTATI JPY/USD dal valore di 0.0064937556
2023/05/28 17:17:07 	cambi valute correnti: EUR/USD = 1.3135949, GBP/USD = 1.0924388, JPY/USD = 0.008615218
2023/05/28 17:17:08 	cambi valute correnti: EUR/USD = 1.1896889, GBP/USD = 1.2271626, JPY/USD = 0.007314072
2023/05/28 17:17:09 	cambi valute correnti: EUR/USD = 1.3838332, GBP/USD = 1.1859553, JPY/USD = 0.0089635635
2023/05/28 17:17:10 	cambi valute correnti: EUR/USD = 1.2746558, GBP/USD = 1.0694026, JPY/USD = 0.008337241
2023/05/28 17:17:11 VENDUTI EUR/USD dal valore di  1.3135949
2023/05/28 17:17:11 	cambi valute correnti: EUR/USD = 1.0617746, GBP/USD = 1.3947043, JPY/USD = 0.007993767
2023/05/28 17:17:12 	cambi valute correnti: EUR/USD = 1.4181266, GBP/USD = 1.3079457, JPY/USD = 0.008846561
2023/05/28 17:17:13 	cambi valute correnti: EUR/USD = 1.3047338, GBP/USD = 1.0250438, JPY/USD = 0.008553812
2023/05/28 17:17:14 ACQUISTATI JPY/USD dal valore di 0.007993767
2023/05/28 17:17:14 	cambi valute correnti: EUR/USD = 1.1489575, GBP/USD = 1.0705942, JPY/USD = 0.006530786
2023/05/28 17:17:15 	cambi valute correnti: EUR/USD = 1.0316505, GBP/USD = 1.0185449, JPY/USD = 0.007454509
2023/05/28 17:17:16 	cambi valute correnti: EUR/USD = 1.2792473, GBP/USD = 1.4476086, JPY/USD = 0.008748325
2023/05/28 17:17:17 ACQUISTATI GBP/USD dal valore di 1.0705942
2023/05/28 17:17:17 	cambi valute correnti: EUR/USD = 1.0635916, GBP/USD = 1.0111102, JPY/USD = 0.008943157
2023/05/28 17:17:18 	cambi valute correnti: EUR/USD = 1.3733578, GBP/USD = 1.1650856, JPY/USD = 0.006992888
2023/05/28 17:17:19 	cambi valute correnti: EUR/USD = 1.3749101, GBP/USD = 1.4246812, JPY/USD = 0.007283292
2023/05/28 17:17:20 ACQUISTATI GBP/USD dal valore di 1.0111102
2023/05/28 17:17:20 	cambi valute correnti: EUR/USD = 1.0827874, GBP/USD = 1.0577636, JPY/USD = 0.00708981
2023/05/28 17:17:21 	cambi valute correnti: EUR/USD = 1.4431722, GBP/USD = 1.4760121, JPY/USD = 0.008227829
2023/05/28 17:17:22 	cambi valute correnti: EUR/USD = 1.156206, GBP/USD = 1.3657992, JPY/USD = 0.006294121
2023/05/28 17:17:23 ACQUISTATI GBP/USD dal valore di 1.0577636
#### FINE SIMULAZIONE ####
```
