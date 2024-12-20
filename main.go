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
	Cursor       int
	IsPlayerTurn bool
	CloseGame    bool
}

var CloseGameError = fmt.Errorf("game closed")

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return CloseGameError
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.Cursor += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.Cursor -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.Cursor += TableSize
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.Cursor -= TableSize
	}

	if g.Cursor > len(g.Table.Cells)-1 {
		g.Cursor = 0
	}
	if g.Cursor < 0 {
		g.Cursor = TableSize - 1
	}
	newMove := false
	if g.IsPlayerTurn == false {
		g.IsPlayerTurn = bot_move(g.Table)
		newMove = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.IsPlayerTurn && g.Table.Cells[g.Cursor].Value == EmptyCell {
		g.IsPlayerTurn = false
		g.Table.Cells[g.Cursor].Value = PlayerCell
		newMove = true
	}
	_ = newMove

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

	for ind, c := range g.Table.Cells {
		ishoverd := false
		if g.Cursor == ind {
			ishoverd = true
		}
		xin, yin := g.Table.IndexToCord(ind)
		DrawCell(screen, xin, yin, c.Value, ishoverd)

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
	g.Cursor = 0
	g.IsPlayerTurn = true
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
