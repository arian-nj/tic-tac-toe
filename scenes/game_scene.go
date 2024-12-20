package scenes

import (
	"fmt"
	"image/color"

	"github.com/arian-nj/tic-tac-toe/constants"
	"github.com/arian-nj/tic-tac-toe/table"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameScene struct {
	loaded       bool
	Table        *table.Table
	Cursor       int
	IsPlayerTurn bool
	CloseGame    bool
}

func NewGameScene() *GameScene {
	return &GameScene{
		Table:        &table.Table{},
		Cursor:       0,
		IsPlayerTurn: true,
		loaded:       false,
	}
}

func (g *GameScene) IsLoaded() bool {
	return g.loaded
}

// Draw implements Scene.
func (g *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 160, 255, 255})
	txt := fmt.Sprintf("%.1f\nuse WASD to move cursor\nSpace to fill Cells", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, txt)

	for ind, c := range g.Table.Cells {
		ishoverd := false
		if g.Cursor == ind {
			ishoverd = true
		}
		xin, yin := g.Table.IndexToCord(ind)
		table.DrawCell(screen, xin, yin, c.Value, ishoverd)

	}
}

// FirstLoad implements Scene.
func (g *GameScene) FirstLoad() {
	g.Table = table.NewTable(constants.TableSize)
	g.Cursor = 0
	g.IsPlayerTurn = true
	g.loaded = true
}

// OnEnter implements Scene.
func (g *GameScene) OnEnter() {
}

// OnExit implements Scene.
func (g *GameScene) OnExit() {
}

// Update implements Scene.
func (g *GameScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ExitSceneId
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return PauseSceneId
	}

	g.Cursor = table.CursorMovement(g.Cursor)

	newMove := false
	if g.IsPlayerTurn == false {
		g.IsPlayerTurn = table.BotMove(g.Table)
		newMove = true
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && g.IsPlayerTurn && g.Table.Cells[g.Cursor].Value == table.EmptyCell {
		g.IsPlayerTurn = false
		g.Table.Cells[g.Cursor].Value = table.PlayerCell
		newMove = true
	}

	if newMove {
		who, isWon := g.Table.CheckWin()
		if isWon {
			winner := "no one"
			if who == table.PlayerCell {
				winner = "player"
			} else {
				winner = "bot"
			}
			fmt.Println("winner is ", winner)
			return ExitSceneId
		}
	}
	return GameSceneId
}

var _ Scene = (*GameScene)(nil)
