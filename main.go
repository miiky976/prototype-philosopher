package main

import (
	"fmt"
	"sync"

	"github.com/miiky976/prototype-philosopher/props"
)

var (
	eatWgroup sync.WaitGroup              // variable para esperar a la ejecucion de todos los hilos
	total     int            = 5          // Cuantos filosofos quieres ejecutar
	color     string         = "\033[34m" // simple color para resaltar
	reset     string         = "\033[0m"  // regresar el color de la terminal
)

func main() {
	// Creando forks
	forke := new(props.Fork)
	forks := make([]*props.Fork, total)
	for i := 0; i < total; i++ {
		forks[i] = forke.Clone()
	}

	continuar := 1

	// secuencia para preguntar si repetir la ejecucion del programa
	for continuar == 1 {
		// Se crea el filosofo y se le establecen dos tenedores
		philosophers := make([]*props.Philosopher, total)
		for i := 0; i < total; i++ {
			philosophers[i] = &props.Philosopher{
				Id: i, LFork: forks[i], RFork: forks[(i+1)%total]}
			eatWgroup.Add(1)
			go func(i int) {
				defer eatWgroup.Done()
				philosophers[i].Start()
			}(i)
		}
		eatWgroup.Wait()
		fmt.Printf("%s¿Deseas continuar?(1:si/0:no): %s", color, reset)
		fmt.Scan(&continuar)
	}
}
