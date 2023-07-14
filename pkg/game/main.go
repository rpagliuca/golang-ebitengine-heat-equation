package game

import (
	"app/pkg/heat"
	_ "image/png"
	"log"

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

	g.Simulation.Progress(100)
	grid := g.Simulation.GetGrid()
	ratio := Size / g.Simulation.GetSize()

	for i := 0; i < g.Simulation.GetSize(); i++ {
		for j := 0; j < g.Simulation.GetSize(); j++ {
			clr := cm.At(grid[i][j] / 256.0)
			ebitenutil.DrawRect(screen, float64(ratio*i), float64(ratio*j), float64(ratio), float64(ratio), clr)
		}
	}
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
