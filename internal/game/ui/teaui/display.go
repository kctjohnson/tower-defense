package teaui

import (
	"math"

	"github.com/charmbracelet/lipgloss"

	"ecstemplate/internal/display"
	"ecstemplate/internal/game/components"
	"ecstemplate/pkg/ecs"
)

type Cell struct {
	Symbol rune
	BG     lipgloss.Color
	FG     lipgloss.Color
}

type Buffer struct {
	Cells         [][]Cell
	Width, Height int
}

func (b Buffer) String() string {
	var out string
	for y := range b.Cells {
		for x := range b.Cells[y] {
			cell := b.Cells[y][x]
			out += lipgloss.NewStyle().
				Background(cell.BG).
				Foreground(cell.FG).
				Render(string(cell.Symbol))
		}
		out += "\n"
	}
	return out
}

type DisplayManager struct {
	buffer *Buffer
}

func (dm *DisplayManager) Initialize(width, height int) error {
	dm.buffer = &Buffer{
		Cells:  make([][]Cell, height),
		Width:  width,
		Height: height,
	}

	for i := range dm.buffer.Cells {
		dm.buffer.Cells[i] = make([]Cell, width)
	}

	return nil
}

func (dm *DisplayManager) Clear() {
	for y := range dm.buffer.Height {
		for x := range dm.buffer.Width {
			dm.buffer.Cells[y][x] = Cell{Symbol: ' '}
		}
	}
}

func (dm *DisplayManager) Render(
	world *ecs.World,
	componentAccess *components.ComponentAccess,
) {
	// Render the path points for now
	paths := world.ComponentManager.GetAllEntitiesWithComponents(
		[]ecs.ComponentType{
			components.Path,
		},
	)

	for _, path := range paths {
		pathComp, _ := componentAccess.GetPathComponent(path)
		for _, point := range pathComp.Waypoints {
			x := int(math.Round(point.X))
			y := int(math.Round(point.Y))

			if x < 0 || x >= dm.buffer.Width || y < 0 || y >= dm.buffer.Height {
				continue
			}

			dm.buffer.Cells[y][x] = Cell{
				Symbol: '*',
				BG:     lipgloss.Color("#000000"),
				FG:     lipgloss.Color("#337733"),
			}
		}
	}

	// Render the entities that have rendering and a position
	renderables := world.ComponentManager.GetAllEntitiesWithComponents(
		[]ecs.ComponentType{
			components.Renderable,
			components.Position,
		},
	)

	for _, renderable := range renderables {
		rend, _ := componentAccess.GetRenderableComponent(renderable)
		pos, _ := componentAccess.GetPositionComponent(renderable)
		dm.RenderEntity(renderable, pos, rend)
	}
}

func (dm *DisplayManager) RenderEntity(
	entity ecs.Entity,
	position *components.PositionComponent,
	renderable *components.RenderableComponent,
) {
	x := int(math.Round(position.X))
	y := int(math.Round(position.Y))

	if x < 0 || x >= dm.buffer.Width || y < 0 || y >= dm.buffer.Height {
		return
	}

	dm.buffer.Cells[y][x] = Cell{
		Symbol: rune(renderable.Symbol[0]),
		BG:     lipgloss.Color("#000000"),
		FG:     lipgloss.Color("#CCCCCC"),
	}
}

func (dm *DisplayManager) RenderUI(gameInfo display.GameInfo) {
	// Just display it over the top for now
}

func (dm *DisplayManager) Update() {
	// No-op, bubbletea reads the buffer directly
}

func (dm *DisplayManager) Shutdown() {
	// No-op since we're not doing any allocating or channel closing
}

func (dm *DisplayManager) GetBuffer() *Buffer {
	return dm.buffer
}

func (dm *DisplayManager) Resize(width, height int) {
	dm.buffer.Width = width
	dm.buffer.Height = height
	dm.buffer.Cells = make([][]Cell, height)
	for i := range dm.buffer.Cells {
		dm.buffer.Cells[i] = make([]Cell, width)
	}
	dm.Clear()
}
