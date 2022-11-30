package props

import (
	"fmt"
	"sync"
	"time"
)

var (
	times int = 5
)

// Interfaces de prototype

type forkprototype interface {
	Clone() *Fork
}

type philprototype interface {
	Clone() *Philosopher
}

// Funciones de clonado (descendientes de las interfaces anteriores)

func (f *Fork) Clone() *Fork {
	return &Fork{
		Mutex: f.Mutex,
	}
}

func (p *Philosopher) Clone() *Philosopher {
	return &Philosopher{
		Id:    p.Id,
		LFork: p.LFork,
		RFork: p.RFork,
	}
}

// Estructuras de Fork y Philosopher

type Fork struct{ sync.Mutex }

type Philosopher struct {
	Id           int
	LFork, RFork *Fork
}

// Funciones de Philosopher

func (p Philosopher) eat() {
	p.LFork.Lock()
	p.RFork.Lock()

	say("Comiendo", p.Id)
	wait()

	p.RFork.Unlock()
	p.LFork.Unlock()

	say("Terminando de comer", p.Id)
	wait()
}

func (p Philosopher) think() {
	say("Pensando sobre la vida", p.Id)
}

func say(action string, id int) {
	fmt.Printf("Filosofo #%d esta %s\n", id, action)
}

func (p Philosopher) Start() {
	for j := 0; j < times; j++ {
		p.think()
		p.eat()
	}
}

// Funciones para acortar
func wait() {
	// Espera 1 segundo
	time.Sleep(50 * time.Millisecond)
}
