package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Rappresnta il cliente di un autonoleggio.
type Client struct {
	Name string  // nome cliente
	Car  Vehicle // veicolo prenotato
}

// Stampa le informazioni relative ad un cliente: nome e tipo di veicolo nollegiato.
func (cl *Client) String() string {
	var ctn string // car type name
	switch cl.Car.Vt {
	case SUV:
		ctn = "SUV"
	case Berlina:
		ctn = "Berlina"
	case StationWagon:
		ctn = "Station Wagon"
	default:
		ctn = ""
	}
	return fmt.Sprintf("name: %s, ha noleggiato un veicolo di tipo: %s", cl.Name, ctn)
}

// Rappresenta un veicolo.
type Vehicle struct {
	Vt VehicleType // tipologia del veicolo
}

// Un vehicleType indica la tipologia di un veicolo.
type VehicleType int

const (
	SUV VehicleType = iota
	Berlina
	StationWagon
)

// Gestisce il nolleggio di veicoli di un concessionario: ad ogni cliente viene
// associata una tipologia di veicolo in maniera randomica.
// veicolo random le tipologie disponibili ai clienti passati.
// wg viene utilizzato per sincronizzare la go rouitine.
// data è canale in cui vengono inviate le tipologie di veicoli che sono state
// assegnate al cliente.
// clients sono i clienti a cui vengono associati randomicamente le tiopologie di
// veicoli.
func CarRental(wg *sync.WaitGroup, data chan VehicleType, clients []Client) {
	defer wg.Done()

	var rentalAgents sync.WaitGroup
	for i := 0; i < len(clients); i++ {
		rentalAgents.Add(1)
		go setRandomVehicleType(&rentalAgents, data, &clients[i])
	}

	rentalAgents.Wait()
	close(data)
}

func setRandomVehicleType(wg *sync.WaitGroup, data chan<- VehicleType, cl *Client) {
	defer wg.Done()

	rnd := rand.Intn(3) // random tra [0, 3[
	cl.Car = Vehicle{Vt: VehicleType(rnd)}
	data <- cl.Car.Vt

	fmt.Println(cl)
}

// Conteggia le tipologie di veicoli che sono stati nollegiati e stampa il resoconto
// finale.
// wg viene utilizzato per sincronizzare la go rouitine.
// data è il canale in cui vengono inviate le tipologie di veicoli.
func CountVehicleTypes(wg *sync.WaitGroup, data <-chan VehicleType) {
	defer wg.Done()

	s, b, sw := 0, 0, 0 // contatori
	for i := range data {
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
	clients := make([]Client, 10)
	for i := range clients {
		clients[i].Name = string('A' + i)
	}

	data := make(chan VehicleType, 10) // per i nolleggi

	var wg sync.WaitGroup
	wg.Add(2)
	go CountVehicleTypes(&wg, data)
	go CarRental(&wg, data, clients)
	wg.Wait()
}
