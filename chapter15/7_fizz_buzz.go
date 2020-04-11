package chapter15

import (
	"fmt"
	"sync"
)

// FizzBuzzMultiThread implements the fizzbuzz algorithm in 3 go-routines
func FizzBuzzMultiThread(n int) {
	var fizzWaitGroup sync.WaitGroup
	var buzzWaitGroup sync.WaitGroup
	var numWaitGroup sync.WaitGroup

	for i := 0; i < n; i++ {
		fmt.Printf("%d = ", i)
		fizzWaitGroup.Add(1)
		go func(index int) {
			if index%3 == 0 {
				print("Fizz")
			}
			fizzWaitGroup.Done()
		}(i)
		fizzWaitGroup.Wait()

		buzzWaitGroup.Add(1)
		go func(index int) {
			if index%5 == 0 {
				print("Buzz")
			}

			buzzWaitGroup.Done()
		}(i)
		buzzWaitGroup.Wait()

		numWaitGroup.Add(1)
		go func(index int) {
			if index%3 != 0 && index%5 != 0 {
				print(index)
			}
			numWaitGroup.Done()
		}(i)
		numWaitGroup.Wait()
		print("\n")
	}

}
