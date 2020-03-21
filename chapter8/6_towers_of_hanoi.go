package chapter8

import (
	"container/list"
	"errors"
	"fmt"
	"log"
)

type Tower struct {
	Disks *list.List
	Index int
}

func NewTower(index int) *Tower  {
	return &Tower{
		Disks: list.New(),
		Index: index,
	}
}

func (tower *Tower)Add(d int) error {
	if tower.Disks.Len() != 0 && tower.Disks.Front().Value.(int) <= d {
		return errors.New(fmt.Sprintf("Error placing disk %d", d))
	}

	tower.Disks.PushFront(d)
	return nil
}

func (tower *Tower) MoveTopTo(t *Tower) {
	top := tower.Disks.Remove(tower.Disks.Front())

	_ = t.Add(top.(int))
}

func (tower *Tower) MoveDisks(n int, destination *Tower, buffer *Tower)  {
	if n <= 0 {
		return
	}

	// Move top n-1 disks from origin to buffer using the destination as buffer
	tower.MoveDisks(n-1, buffer, destination)

	// move top from origin to destination
	tower.MoveTopTo(destination)

	// move top n-1 disks from buffer to destination using origin as buffer
	buffer.MoveDisks(n-1, destination, tower)
}


func SolveForN3() {
	n:= 3
	var towers [3]*Tower

	for i:= 0; i< n-1; i++  {
		towers[i] = NewTower(i)
	}

	for i:= n -1; i >= 0; i-- {
		err := towers[0].Add(i)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	towers[0].MoveDisks(n, towers[2], towers[1])
}




