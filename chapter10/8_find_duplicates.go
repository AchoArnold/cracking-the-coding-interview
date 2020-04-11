package chapter10

// BitSet is a set of bits
type BitSet struct {
	BitSet []int
}

// NewBitSet creates a new bit set
func NewBitSet(size int) BitSet {
	return BitSet{
		BitSet: make([]int, (size>>5)+1), // divide by 32
	}
}

func (bs BitSet) get(pos int) bool {
	wordNumber := pos >> 5  // divide by 32
	bitNumber := pos & 0x1F // mod 32

	return bs.BitSet[wordNumber]&(1<<bitNumber) != 0
}

func (bs BitSet) set(pos int) {
	wordNumber := pos >> 5  // divide by 32
	bitNumber := pos & 0x1F // mod 32

	bs.BitSet[wordNumber] |= 1 << bitNumber
}

// GetDuplicates returns the duplicates in array
func GetDuplicates(array []int) (result []int) {
	bs := NewBitSet(3200)
	for i := 0; i < len(array); i++ {
		num := array[i]
		num0 := num - 1 // bet set starts at 0, numbers start at 1
		if bs.get(num0) {
			result = append(result, num)
		} else {
			bs.set(num0)
		}
	}

	if len(result) > 0 {
		return result
	}

	return []int{}
}
