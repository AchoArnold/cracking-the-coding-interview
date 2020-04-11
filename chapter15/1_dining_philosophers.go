package chapter15

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

// ChopStick is s the chopstick needed for dining
type ChopStick struct {
	Ctx  context.Context
	Lock *semaphore.Weighted
}

// NewChopStick creates a new chopstick
func NewChopStick(ctx context.Context) *ChopStick {
	return &ChopStick{
		Lock: semaphore.NewWeighted(1),
		Ctx:  ctx,
	}
}
func (cs *ChopStick) pickUp() {
	err := cs.Lock.Acquire(cs.Ctx, 1)
	if err != nil {
		log.Fatalf("Error aquireing lock %+v", err.Error())
	}
}

func (cs *ChopStick) putDown() {
	cs.Lock.Release(1)
}

// Philosopher is a person who can eat with a chopstick
type Philosopher struct {
	bites       int
	left, right *ChopStick
	name        string
}

// NewPhilosopher creates a new instance of a philosopher
func NewPhilosopher(left, right *ChopStick, name string) *Philosopher {
	return &Philosopher{
		bites: 10,
		left:  left,
		right: right,
		name:  name,
	}
}

func (p *Philosopher) eat(index int) {
	log.Printf("%s started eating n= %d\n", p.name, index)

	p.pickUp()
	p.chew()
	p.putDown()

	log.Printf("%s finished eating n=%d\n", p.name, index)
}

func (p *Philosopher) chew() {
	log.Printf("%s is chewing\n", p.name)
	time.Sleep(1 * time.Second)
}

func (p *Philosopher) pickUp() bool {
	// pick up left
	p.left.pickUp()
	log.Printf("%s picked up left chopstick\n", p.name)

	// pick up right
	p.right.pickUp()
	log.Printf("%s picked up right chopstick\n", p.name)

	return true
}

func (p *Philosopher) putDown() {
	p.right.putDown()
	p.left.putDown()
	log.Printf("%s put down both chopsticks\n", p.name)
}

func (p *Philosopher) run() {
	for i := 0; i < p.bites; i++ {
		p.eat(i)
	}
}
