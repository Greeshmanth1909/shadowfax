package board

func ClearPVTable(brd *S_Board) {
	for i := range brd.PvTable.PvTableEntries {
		brd.PvTable.PvTableEntries[i].PosKey = 0
		brd.PvTable.PvTableEntries[i].Move = 0
	}
}

func InitPvTable(brd *S_Board, num int) {
	var newP PvTable
	newP.PvTableEntries = make([]PvEntry, num)
	newP.NumEntries = num
	brd.PvTable = &newP
}
