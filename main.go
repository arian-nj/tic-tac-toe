package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Table        *Table
	CursorX      int
	CursorY      int
	IsPlayerTurn bool
	CloseGame    bool
}

var CloseGameError = fmt.Errorf("game closed")

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return CloseGameError
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
	newMove := false
	if g.IsPlayerTurn == false {
		g.IsPlayerTurn = bot_move(g.Table)
		newMove = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.IsPlayerTurn && g.Table.Cells[g.CursorY][g.CursorX].Value == EmptyCell {
		g.IsPlayerTurn = false
		g.Table.Cells[g.CursorY][g.CursorX].Value = PlayerCell
		newMove = true
	}

	if newMove {
		who, isWon := g.Table.checkWin()
		if isWon {
			winner := "no one"
			if who == PlayerCell {
				winner = "player"
			} else {
				winner = "bot"
			}
			fmt.Println("winner is ", winner)
			return CloseGameError
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 160, 255, 255})
	txt := fmt.Sprintf("%.1f\nuse WASD to move cursor\nSpace to fill Cells", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, txt)

	for indexY, RowCells := range g.Table.Cells {
		for indexX, c := range RowCells {
			ishoverd := false
			if g.CursorX == indexX && g.CursorY == indexY {
				ishoverd = true
			}
			DrawCell(screen, indexX, indexY, c.Value, ishoverd)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
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
