package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuScene struct {
	loaded bool
}

func NewMenuScene() *MenuScene {
	return &MenuScene{
		loaded: false,
	}
}
func (s *MenuScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255})
	ebitenutil.DebugPrint(screen, "Press ENTER to Start")
}

func (s *MenuScene) FirstLoad() {
	s.loaded = true
}

func (s *MenuScene) IsLoaded() bool {
	return s.loaded
}

// OnEnter implements Scene.
func (s *MenuScene) OnEnter() {
}

// OnExit implements Scene.
func (s *MenuScene) OnExit() {
}

// Update implements Scene.
func (s *MenuScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return GameSceneId
	}
	return MenuSceneId
}

var _ Scene = (*MenuScene)(nil)
