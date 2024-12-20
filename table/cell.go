package table

import (
	"image/color"

	"github.com/arian-nj/tic-tac-toe/constants"
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
	xPos := float32((indexX * (constants.CubeSize + constants.CubePadding)) + constants.TableStartX)
	yPos := float32((indexY * (constants.CubeSize + constants.CubePadding)) + constants.TableStartY)
	if ishoverd {
		col = color.RGBA{120, 255, 120, 255}
	}

	if value == EmptyCell {
		vector.StrokeRect(screen,
			xPos, yPos,
			constants.CubeSize, constants.CubeSize, constants.CubeBorderSize,
			col, true)
	} else {
		if value == PlayerCell {
			col = color.RGBA{0, 0, 100, 255}
		} else if value == BotCell {
			col = color.RGBA{40, 0, 0, 255}
		}

		vector.DrawFilledRect(screen,
			xPos, yPos,
			constants.CubeSize, constants.CubeSize,
			col, true)
	}
	if ishoverd {
		col = color.RGBA{120, 255, 120, 255}
	}
	vector.StrokeRect(screen,
		xPos, yPos,
		constants.CubeSize, constants.CubeSize, constants.CubeBorderSize,
		col, true)
}
