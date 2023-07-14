package heat

const Size = 200

type Grid [Size][Size]float64

type Drawer interface {
	Draw(Grid)
}

const gamma = 0.25

type Simulation struct {
	grid *Grid
}

func NewSimulation() *Simulation {
	var grid Grid
	initialConditions(&grid)
	return &Simulation{&grid}
}

func initialConditions(grid *Grid) {
	fixedConditions(grid)
}

func fixedConditions(grid *Grid) {
	for i := int(Size/2.0 - 20.0); i < int(Size/2.0+20.0); i++ {
		grid[0][i] = 255.0
		grid[Size-1][i] = 255.0
		grid[i][0] = 255.0
		grid[i][Size-1] = 255.0
	}
	for i := -10; i <= 10; i++ {
		for j := -10; j <= 10; j++ {
			if i*i+j*j < 10*10 {
				center := int(Size / 2.0)
				grid[center+i][center+j] = 255.0
			}
		}
	}
	/*
		for i := -10; i <= 10; i++ {
			for j := -10; j <= 10; j++ {
				if i*i+j*j < 10*10 {
					center := int(Size / 4.0)
					grid[center+i][center+j] = 255.0
					center = int(Size / 4.0 * 3.0)
					grid[center+i][center+j] = 255.0
				}
			}
		}
	*/
}

func (s *Simulation) Progress(ticks int) {
	for i := 0; i < ticks; i++ {
		s.OneNextTick()
	}
}

func (s *Simulation) OneNextTick() {
	g := s.grid
	nextGrid := Grid{}

	// We should always start copying last boundary value
	nextGrid[0] = g[0]
	nextGrid[Size-1] = g[Size-1]
	for i := 0; i < Size; i++ {
		nextGrid[i][0] = g[i][0]
		nextGrid[i][Size-1] = g[i][Size-1]
	}

	for i := 1; i < Size-1; i++ {
		for j := 1; j < Size-1; j++ {
			// Centro2 = Gamma(Esquerda + Direita + Cima + Baixo - 4*Centro1) + Centro1
			nextGrid[i][j] = gamma*
				(g[i-1][j]+g[i+1][j]+g[i][j-1]+g[i][j+1]-4*g[i][j]) +
				g[i][j]
		}
	}

	fixedConditions(&nextGrid)
	s.grid = &nextGrid
}

func (s *Simulation) GetGrid() *Grid {
	return s.grid
}

func (s *Simulation) GetSize() int {
	return Size
}
