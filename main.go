package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Cell struct {
	Value int
}

type Table struct {
	Cells [][]Cell
}

type Game struct {
	Table        *Table
	CursorX      int
	CursorY      int
	IsPlayerTurn bool
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return fmt.Errorf("game closed")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.CursorX += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.CursorX -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.CursorY += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.CursorY -= 1
	}

	if g.CursorX > TableSize-1 {
		g.CursorX = 0
	}
	if g.CursorX < 0 {
		g.CursorX = TableSize - 1
	}

	if g.CursorY > TableSize-1 {
		g.CursorY = 0
	}
	if g.CursorY < 0 {
		g.CursorY = TableSize - 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.IsPlayerTurn {
		g.IsPlayerTurn = false

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 160, 255, 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.1f", ebiten.ActualTPS()))

	for indexY, RowCells := range g.Table.Cells {
		for indexX, _ := range RowCells {
			var col color.Color = color.Black
			if g.IsPlayerTurn {
				if g.CursorX == indexX && g.CursorY == indexY {
					col = color.RGBA{255, 255, 255, 1}
				}
			}

			vector.StrokeRect(screen,
				float32((indexX*(CubeSize+CubePadding))+TableStartX), float32((indexY*(CubeSize+CubePadding))+TableStartY),
				CubeSize, CubeSize, CubeBorderSize,
				col, true)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
func NewTable(tableSize int) *Table {
	table := &Table{}
	for range tableSize {
		rowCell := []Cell{}
		for range tableSize {
			rowCell = append(rowCell, Cell{Value: 0})
		}
		table.Cells = append(table.Cells, rowCell)
	}
	return table
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := &Game{}
	g.Table = NewTable(TableSize)
	g.CursorX = 0
	g.CursorY = 0
	g.IsPlayerTurn = true
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
