package main

import (
    "fmt"
    "sync"
    "math/rand"
    "time"
)

// rappresnta il cliente di un autonoleggio
type client struct {
    name string
    car vehicle
}

// rappresenta un veicolo
type vehicle struct {
    vt vehicleType
}

// tipi di veicoli disponibili
type vehicleType int
const (
    SUV vehicleType = iota
    Berlina
    StationWagon
)

// metodo che torna una stringa contenete le info relative al cliente
func (cl *client) String() string {
    var ctn string // car type name
    switch cl.car.vt {
    case SUV:
        ctn = "SUV"
    case Berlina:
        ctn = "Berlina"
    case StationWagon:
        ctn = "Station Wagon"
    }
    return fmt.Sprintf("name: %s, ha noleggiato un veicolo di tipo: %s", cl.name, ctn)
}


// gestisce il noleggio auto dei clienti passati
func carRental(wg *sync.WaitGroup, data chan vehicleType, clients []client) {
    var sellers sync.WaitGroup
    for i := 0; i < len(clients); i++ {
        sellers.Add(1)
        go rent(&sellers, data, &clients[i])
    }
    sellers.Wait()
    close(data)
    wg.Done()
}

// associa un tipo di veicolo random al cliente
func rent(wg *sync.WaitGroup, data chan<- vehicleType, cl *client) {
    rnd := rand.Intn(3) // random tra 0 e 3
    cl.car = vehicle{vt: vehicleType(rnd)}
    data <- cl.car.vt
    fmt.Println(cl)
    wg.Done()
}

// conteggio auto nolleggiete per tipo e stampa resoconto finale
func accountant(wg *sync.WaitGroup, data <-chan vehicleType) {
    s, b, sw := 0, 0, 0
    for i := range(data) { // aspetta la chiusura del channel
        switch i { // filtra in base al tipo di veicolo
        case SUV:
            s++
        case Berlina:
            b++
        case StationWagon:
            sw++
        }
    }
    fmt.Println("numero totale di SUV =", s)
    fmt.Println("numero totale di Berline =", b)
    fmt.Println("numero totale di Station Wagon=", sw)
    wg.Done()
}

func main() {
    rand.Seed(time.Now().UnixNano()) // per la randomicità

    // alloco 10 clienti
    clients := make([]client, 10)
    for i := range(clients) {
        clients[i].name = string(i + 65) // 'A' = 65
    }

    var wg sync.WaitGroup
    data := make(chan vehicleType, 10) // per i nolleggi

    wg.Add(2)
    go accountant(&wg, data)
    go carRental(&wg, data, clients)
    wg.Wait()
}
