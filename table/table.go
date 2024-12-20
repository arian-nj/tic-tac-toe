package table

import "github.com/arian-nj/tic-tac-toe/constants"

type Table struct {
	Cells []*Cell
}

func (t *Table) CordToIndex(xIndex, yIndex int) int {
	return yIndex*constants.TableSize + xIndex
}
func (t *Table) IndexToCord(index int) (xIndex, yIndex int) {
	xIndex = index % constants.TableSize
	yIndex = index / constants.TableSize
	return xIndex, yIndex
}

func NewTable(tableSize int) *Table {
	table := &Table{}
	for range tableSize {
		for range tableSize {
			table.Cells = append(table.Cells, &Cell{Value: EmptyCell})
		}
	}
	return table
}
func isAllSame(vals []int) bool {
	last := vals[0]
	if last == EmptyCell {
		return false
	}
	for _, v := range vals {
		if v != last {
			return false
		}
	}
	return true
}

// hardcoded only for 3*3
func (t *Table) CheckWin() (int, bool) {
	possibles := [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},

		[]int{0, 3, 6},
		[]int{1, 4, 7},
		[]int{2, 5, 8},

		[]int{0, 4, 8},
		[]int{2, 4, 6},
	}
	for _, pos := range possibles {
		items := []int{}
		for _, ci := range pos {
			items = append(items, t.Cells[ci].Value)
		}
		if isAllSame(items) {
			return items[0], true
		}
	}

	return 0, false
}
