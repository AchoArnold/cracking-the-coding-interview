package chapter15

import (
	"context"
	"sync"
	"testing"
)

func TestDiningPhilosophers(t *testing.T) {
	ctx := context.Background()
	leftChopStick := NewChopStick(ctx)
	rightChopStick := NewChopStick(ctx)

	var philosophers []*Philosopher
	for _, name := range []string{"0", "1", "2", "3", "4"} {
		philosophers = append(philosophers, NewPhilosopher(leftChopStick, rightChopStick, name))
	}

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go func(philosopher *Philosopher) {
			philosopher.run()
			wg.Done()
		}(philosopher)
	}

	wg.Wait()
}
