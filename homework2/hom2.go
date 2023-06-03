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
    car vehicle // veicolo prenotato
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
    default:
        ctn = ""
    }
    return fmt.Sprintf("name: %s, ha noleggiato un veicolo di tipo: %s", cl.name, ctn)
}

// gestisce il noleggio auto dei clienti passati
func carRental(wg *sync.WaitGroup, data chan vehicleType, clients []client) {
    defer wg.Done()
    var rentalAgents sync.WaitGroup
    for i := 0; i < len(clients); i++ {
        rentalAgents.Add(1)
        go rent(&rentalAgents, data, &clients[i])
    }
    rentalAgents.Wait()
    close(data)
}

// associa un tipo di veicolo random al cliente
func rent(wg *sync.WaitGroup, data chan<- vehicleType, cl *client) {
    defer wg.Done()
    rnd := rand.Intn(3) // random tra [0, 3[
    cl.car = vehicle{vt: vehicleType(rnd)}
    data <- cl.car.vt
    fmt.Println(cl)
}

// conteggio auto nolleggiete per tipo e stampa resoconto finale
func accountant(wg *sync.WaitGroup, data <-chan vehicleType) {
    defer wg.Done()
    s, b, sw := 0, 0, 0 // contatori
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
    fmt.Println("\n#### RESOCONTO FINALE ####",
                "\nnumero totale di SUV =", s, 
                "\nnumero totale di Berline =", b, 
                "\nnumero totale di Station Wagon =", sw)
}

func main() {
    rand.Seed(time.Now().UnixNano()) // per la randomicità

    // alloco 10 clienti
    clients := make([]client, 10)
    for i := range(clients) {
        clients[i].name = string('A' + i) 
    }

    data := make(chan vehicleType, 10) // per i nolleggi

    var wg sync.WaitGroup
    wg.Add(2)
    go accountant(&wg, data)
    go carRental(&wg, data, clients)
    wg.Wait()
}
