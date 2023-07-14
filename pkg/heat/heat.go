package heat

const Size = 100.0

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
	// Everything starts with zero temperature
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			grid[i][j] = 0.0
		}
	}

	fixedConditions(grid)
}

func copyBoundariesFromNeighbors(grid *Grid) {
	grid[0] = grid[1]
	grid[Size-1] = grid[Size-2]
	for i := 0; i < Size; i++ {
		grid[i][0] = grid[i][1]
		grid[i][Size-1] = grid[i][Size-2]
	}
}

func fixedConditions(grid *Grid) {

	// Walls should be fixed
	temp := 0.0
	for i := 0; i < Size; i++ {
		grid[0][i] = temp
		grid[Size-1][i] = temp
		grid[i][0] = temp
		grid[i][Size-1] = temp
	}

	unit := int(Size * 0.05)

	for i := int(Size/2.0 - 4*unit); i < int(Size/2.0+4*unit); i++ {
		//grid[0][i] = 255.0
		//grid[Size-1][i] = 255.0
		//grid[i][0] = 255.0
		//grid[i][Size-1] = 255.0
	}
	for i := -unit; i <= unit; i++ {
		for j := -unit; j <= unit; j++ {
			if i*i+j*j < unit*unit {
				center := int(Size / 2.0)
				_ = center
				//grid[center+i][center+j] = 255.0
			}
		}
	}
	for i := -unit; i <= unit; i++ {
		for j := -unit; j <= unit; j++ {
			if i*i+j*j < unit*unit {
				center := int(Size / 4.0)
				grid[center+i][center+j] = 255.0
				center = int(Size / 4.0 * 3.0)
				grid[center+i][center+j] = 255.0
				_ = center
			}
		}
	}
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

	// Copy boundaries from neighbors before applying fixedConditions
	copyBoundariesFromNeighbors(&nextGrid)
	fixedConditions(&nextGrid)
	s.grid = &nextGrid
}

func (s *Simulation) GetGrid() *Grid {
	return s.grid
}

func (s *Simulation) GetSize() int {
	return Size
}
