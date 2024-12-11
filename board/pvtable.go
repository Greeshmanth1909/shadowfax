package board

func InitPvTable(brd *S_Board) {
	var new PvTable
	brd.PvTable = &new
}

/* I'm not sure weather I should manually clear pv table or let the GC handle it. I'll let the GC handle it for now
   i.e. call intipvTable when it needs clearing
*/

func SetPvTable(brd *S_Board, pos uint64, move uint32) {
	var pvE PvEntry
	pvE.Move = move
	pvE.PosKey = pos
	brd.PvTable.PvTableEntries[pos] = pvE
}
