package chapter3

import (
	"container/list"
	"errors"
)

type Shelter struct {
	Inventory *list.List
}

type AnimalNode struct {
	Type  string
	Value int
}

const TypeDog = "DOG"
const TypeCat = "CAT"

func GetNewShelter() *Shelter {
	return &Shelter{list.New()}
}

func (shelter *Shelter) Enqueue(node AnimalNode) {
	shelter.Inventory.PushBack(node)
}

func (shelter *Shelter) IsEmpty() bool {
	return shelter == nil || *shelter == Shelter{} || shelter.Inventory.Front() == nil
}

func (shelter *Shelter) DequeueAny() (AnimalNode, error) {
	if shelter.IsEmpty() {
		return AnimalNode{}, errors.New("cannot dequeue from an empty list")
	}

	element := shelter.Inventory.Back()
	shelter.Inventory.Remove(element)

	return element.Value.(AnimalNode), nil
}

func (shelter *Shelter) DequeueDog() (AnimalNode, error) {
	if shelter.IsEmpty() {
		return AnimalNode{}, errors.New("cannot dequeue from an empty list")
	}

	element := shelter.Inventory.Front()
	for element != nil {
		if element.Value.(AnimalNode).Type == TypeDog {
			shelter.Inventory.Remove(element)
			return element.Value.(AnimalNode), nil
		}

		element = element.Next()
	}

	return AnimalNode{}, errors.New("there is no dog in the list")
}

func (shelter *Shelter) DequeueCat() (AnimalNode, error) {
	if shelter.IsEmpty() {
		return AnimalNode{}, errors.New("cannot dequeue from an empty list")
	}

	element := shelter.Inventory.Front()
	for element != nil {
		if element.Value.(AnimalNode).Type == TypeCat {
			shelter.Inventory.Remove(element)
			return element.Value.(AnimalNode), nil
		}
	}

	return AnimalNode{}, errors.New("there is no cat in the list")
}
