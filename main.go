package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	eatWgroup sync.WaitGroup              // variable para esperar a la ejecucion de todos los hilos
	total     int            = 5          // Cuantos filosofos quieres ejecutar
	veces     int            = 5          // Cuantas veces quieres que coma cada filosofo antes de poder preguntar por mas
	color     string         = "\033[34m" // simple color para resaltar
	reset     string         = "\033[0m"  // regresar el color de la terminal
)

// Espacio para patron de diseño
type philprototype interface {
	Clone() *philosopher
}

type forkprototype interface {
	Clone() *fork
}

func (p *philosopher) Clone() *philosopher {
	return &philosopher{
		id:        p.id,
		leftFork:  p.leftFork,
		rightFork: p.rightFork,
	}
}

// Funcion para acortar la sentencia de espera :3
func wait() {
	// Espera 1 segundo
	time.Sleep(50 * time.Millisecond)
}

// Son los tenedores :3
type fork struct{ sync.Mutex }

func (f *fork) Clone() *fork {
	return &fork{
		Mutex: f.Mutex,
	}
}

// "Objeto" filosofo que recibe tres objetos o argumentos
type philosopher struct {
	id                  int
	leftFork, rightFork *fork
}

// funciones "propias" del "Objeto" filosofo
func (p philosopher) eat() {
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Printf("Filosofo #%d comiendo con tenedores %d y %d\n", p.id, p.id, ((p.id + 1) % (total)))
	wait()

	p.rightFork.Unlock()
	p.leftFork.Unlock()

	say("Terminando de comer", p.id)
	wait()
}

func (p philosopher) think() {
	say("Pensando sobre la vida", p.id)
}

func (p philosopher) start() {
	defer eatWgroup.Done()
	for j := 0; j < veces; j++ {
		// ---------------------------Intercambiar los comentarios para limitar o no las veces de ejecucion
		//for true {
		p.think()
		p.eat()
	}
}

func say(action string, id int) {
	fmt.Printf("Filosofo #%d esta %s\n", id, action)
}

func main() {
	// Create forks
	forke := new(fork)
	forks := make([]*fork, total)
	for i := 0; i < total; i++ {
		forks[i] = forke.Clone()
	}

	continuar := 1

	// secuencia para preguntar si repetir la ejecucion del programa
	for continuar == 1 {
		// Se crea el filosofo y se le establecen dos tenedores
		philosophers := make([]*philosopher, total)
		for i := 0; i < total; i++ {
			philosophers[i] = &philosopher{
				id: i, leftFork: forks[i], rightFork: forks[(i+1)%total]}
			eatWgroup.Add(1)
			go philosophers[i].start()
		}
		eatWgroup.Wait()
		fmt.Printf("%s¿Deseas continuar?(1:si/0:no): %s", color, reset)
		fmt.Scan(&continuar)
	}
}
