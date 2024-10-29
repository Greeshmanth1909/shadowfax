package board

var Square120to64 [120]int
var Square64to120 [64]int

// InitSquares120 initialises an array of length 120 to have indexes that correspond to a 64 int array
func InitSquares120() {
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
}

// InitSquares64 initializes the 64 int array with values corresponding to 120 int array
func InitSquares64() {
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
}

// FRtoSq120 takes file and rank of a respective square and returns its 64 int array index
func FRtoSq120(file File, rank Rank) int {
	file120 := int(file) + 1
	rank120 := int(rank) + 2
	index := rank120*10 + file120
	return Square120to64[index]
}

// InitFileRankArrays initializes file and rank arrays
func InitFileRankArrays() {
	for i := range FileArr {
		FileArr[i] = File(OFFBOARD)
		RankArr[i] = Rank(OFFBOARD)
	}

	for _, val := range Square64to120 {
		mod := val % 10
		switch mod {
		case 0:
			FileArr[val] = File(OFFBOARD)
		case 1:
			FileArr[val] = FILE_A
		case 2:
			FileArr[val] = FILE_B
		case 3:
			FileArr[val] = FILE_C
		case 4:
			FileArr[val] = FILE_D
		case 5:
			FileArr[val] = FILE_E
		case 6:
			FileArr[val] = FILE_F
		case 7:
			FileArr[val] = FILE_G
		case 8:
			FileArr[val] = FILE_H
		default:
		}
	}

	for _, val := range Square64to120 {
		mod := val / 10
		switch mod {
		case 2:
			RankArr[val] = RANK_1
		case 3:
			RankArr[val] = RANK_2
		case 4:
			RankArr[val] = RANK_3
		case 5:
			RankArr[val] = RANK_4
		case 6:
			RankArr[val] = RANK_5
		case 7:
			RankArr[val] = RANK_6
		case 8:
			RankArr[val] = RANK_7
		case 9:
			RankArr[val] = RANK_8
		default:
		}
	}
}
