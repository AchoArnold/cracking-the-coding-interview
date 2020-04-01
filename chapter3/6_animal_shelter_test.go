package chapter3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewShelter(t *testing.T) {
	t.Run("New returns an empty shelter", func(t *testing.T) {
		shelter := GetNewShelter()
		assert.Equal(t, 0, shelter.Inventory.Len())
	})
}

func TestShelter_Enqueue(t *testing.T) {
	t.Run("Enqueue adds an animal to the shelter", func(t *testing.T) {
		shelter := GetNewShelter()
		shelter.Enqueue(AnimalNode{TypeCat, 33})

		assert.Equal(t, 1, shelter.Inventory.Len())
	})
}

func TestShelter_DequeueAny(t *testing.T) {
	t.Run("Dequeue removes an animal from the shelter", func(t *testing.T) {
		shelter := GetNewShelter()
		shelter.Enqueue(AnimalNode{TypeDog, 12})
		shelter.Enqueue(AnimalNode{TypeCat, 33})
		shelter.Enqueue(AnimalNode{TypeDog, 11})

		element, err := shelter.DequeueAny()

		assert.Nil(t, err)
		assert.True(t, assert.ObjectsAreEqualValues(AnimalNode{TypeDog, 11}, element))
		assert.Equal(t, 2, shelter.Inventory.Len())
	})

	t.Run("Dequeue returns an error if the inventory is empty", func(t *testing.T) {
		shelter := GetNewShelter()

		_, err := shelter.DequeueAny()

		assert.NotNil(t, err)
	})
}

func TestShelter_DequeueDog(t *testing.T) {
	t.Run("DequeueDog removes a dog from the shelter", func(t *testing.T) {
		shelter := GetNewShelter()
		shelter.Enqueue(AnimalNode{TypeCat, 33})
		shelter.Enqueue(AnimalNode{TypeDog, 12})
		shelter.Enqueue(AnimalNode{TypeCat, 11})

		element, err := shelter.DequeueDog()

		assert.Nil(t, err)
		assert.True(t, assert.ObjectsAreEqualValues(AnimalNode{TypeDog, 12}, element))
		assert.Equal(t, 2, shelter.Inventory.Len())
	})

	t.Run("DequeueDog returns an error if the inventory is empty", func(t *testing.T) {
		shelter := GetNewShelter()

		_, err := shelter.DequeueDog()

		assert.NotNil(t, err)
	})
}

func TestShelter_DequeueCat(t *testing.T) {
	t.Run("DequeueCat removes a dog from the shelter", func(t *testing.T) {
		shelter := GetNewShelter()
		shelter.Enqueue(AnimalNode{TypeCat, 33})
		shelter.Enqueue(AnimalNode{TypeDog, 12})
		shelter.Enqueue(AnimalNode{TypeCat, 11})

		element, err := shelter.DequeueCat()

		assert.Nil(t, err)
		assert.True(t, assert.ObjectsAreEqualValues(AnimalNode{TypeCat, 33}, element))
		assert.Equal(t, 2, shelter.Inventory.Len())
	})

	t.Run("DequeueCat returns an error if the inventory is empty", func(t *testing.T) {
		shelter := GetNewShelter()

		_, err := shelter.DequeueCat()

		assert.NotNil(t, err)
	})
}
