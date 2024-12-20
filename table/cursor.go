package table

import (
	"github.com/arian-nj/tic-tac-toe/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func CursorMovement(cursor int) int {
	// if ebiten.IsKeyPressed(ebiten.KeyQ) {
	// 	return CloseGameError
	// }
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		cursor += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		cursor -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		cursor += constants.TableSize
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		cursor -= constants.TableSize
	}

	if cursor > constants.TableSize*constants.TableSize-1 {
		cursor = 0
	}
	if cursor < 0 {
		cursor = constants.TableSize - 1
	}
	return cursor
}
