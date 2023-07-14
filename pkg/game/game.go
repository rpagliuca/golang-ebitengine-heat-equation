package game

import (
	"app/pkg/heat"
	"fmt"
	_ "image/png"
	"log"
	"math"

	"github.com/fogleman/colormap"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Simulation *heat.Simulation
}

func (g *Game) Update() error {
	return nil
}

const Size = 600.0

var cm = colormap.Magma

func (g *Game) Draw(screen *ebiten.Image) {

	g.Simulation.Progress(3)
	grid := g.Simulation.GetGrid()
	ratio := float64(Size) / float64(g.Simulation.GetSize())

	for i := 0; i < g.Simulation.GetSize(); i++ {
		for j := 0; j < g.Simulation.GetSize(); j++ {
			clr := cm.At(grid[i][j])
			ebitenutil.DrawRect(screen, ratio*float64(i), ratio*float64(j), ratio, ratio, clr)
		}
	}

	intMouseX, intMouseY := ebiten.CursorPosition()

	x := int(math.Floor(float64(intMouseX) / ratio))
	y := int(math.Floor(float64(intMouseY) / ratio))

	if mouseInScreen() {
		temp := 15.0 + (285.0)*grid[x][y]
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Temperatura no cursor: %.1f", temp))
	}

	if mouseInScreen() && ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		g.Simulation.AddSource(x, y)
	}

	fps := ebiten.ActualTPS()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.2f", fps), 0, Size-16)
}

func mouseInScreen() bool {
	mX, mY := ebiten.CursorPosition()
	if mX >= 0 && mY >= 0 && mX < Size && mY < Size {
		return true
	}
	return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Size, Size
}

func Run() {
	ebiten.SetWindowSize(Size, Size)
	if err := ebiten.RunGame(&Game{heat.NewSimulation()}); err != nil {
		log.Fatal(err)
	}
}
