package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	EmptyCell = iota
	PlayerCell
	BotCell
)

type Cell struct {
	Value int
}

func DrawCell(screen *ebiten.Image, indexX, indexY int, value int, ishoverd bool) {
	var col color.Color = color.RGBA{0, 0, 0, 255}
	xPos := float32((indexX * (CubeSize + CubePadding)) + TableStartX)
	yPos := float32((indexY * (CubeSize + CubePadding)) + TableStartY)
	if ishoverd {
		col = color.RGBA{120, 255, 120, 255}
	}

	if value == EmptyCell {
		vector.StrokeRect(screen,
			xPos, yPos,
			CubeSize, CubeSize, CubeBorderSize,
			col, true)
	} else {
		if value == PlayerCell {
			col = color.RGBA{0, 0, 100, 255}
		} else if value == BotCell {
			col = color.RGBA{40, 0, 0, 255}
		}

		vector.DrawFilledRect(screen,
			xPos, yPos,
			CubeSize, CubeSize,
			col, true)
	}
	if ishoverd {
		col = color.RGBA{120, 255, 120, 255}
	}
	vector.StrokeRect(screen,
		xPos, yPos,
		CubeSize, CubeSize, CubeBorderSize,
		col, true)
}

type Table struct {
	Cells []*Cell
}

func (t *Table) CordToIndex(xIndex, yIndex int) int {
	return yIndex*TableSize + xIndex
}
func (t *Table) IndexToCord(index int) (xIndex, yIndex int) {
	xIndex = index % TableSize
	yIndex = index / TableSize
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
func (t *Table) checkWin() (int, bool) {
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
