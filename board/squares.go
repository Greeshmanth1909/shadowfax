package board

var Square120to64 [120]int
var Square64to120 [64]int

// InitSquares120 initialises an array of length 120 to have indexes that correspond to a 64 int array
func InitSquares120() [120]int {
	count := 0
	for i := range Square120to64 {
		if i <= 19 || i >= 100 {
			Square120to64[i] = 65
		} else if i%10 == 0 || (i+1)%10 == 0 {
			Square120to64[i] = 65
		} else {
			Square120to64[i] = count
			count++
		}
	}
	return Square120to64
}

// InitSquares64 initializes the 64 int array with values corresponding to 120 int array
func InitSquares64() [64]int {
	value := 21
	for i := range Square64to120 {
		if i%8 == 0 && i != 0 {
			value += 2
			Square64to120[i] = value
			value++
		} else {
			Square64to120[i] = value
			value++
		}
	}
	return Square64to120
}

// FRtoSq120 takes file and rank of a respective square and returns its 64 int array index
func FRtoSq120(file File, rank Rank) int {
	file120 := int(file) + 1
	rank120 := int(rank) + 2
	index := rank120*10 + file120
	return Square120to64[index]
}
