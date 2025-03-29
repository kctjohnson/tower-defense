package teaui

import (
	"fmt"
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

	// Render the cursor
	cursorEnts := world.ComponentManager.GetAllEntitiesWithComponent(components.Cursor)
	if len(cursorEnts) == 1 {
		cursorEnt := cursorEnts[0]
		cursorPos, _ := componentAccess.GetPositionComponent(cursorEnt)
		x := int(math.Round(cursorPos.X))
		y := int(math.Round(cursorPos.Y))

		if x >= 0 && x < dm.buffer.Width && y >= 0 && y < dm.buffer.Height {
			dm.buffer.Cells[y][x] = Cell{
				Symbol: 'X',
				BG:     lipgloss.Color("#000000"),
				FG:     lipgloss.Color("#FF0000"),
			}
		}
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
	dm.writeString(0, 0, gameInfo.Message)
	dm.writeString(0, 1, fmt.Sprintf("Health: %0.2f", gameInfo.PlayerHealth))
	dm.writeString(0, 2, fmt.Sprintf("Money: %0.2f", gameInfo.PlayerMoney))
	dm.writeString(0, 3, fmt.Sprintf("Wave: %d", gameInfo.CurrentWave))
	dm.writeString(0, 4, fmt.Sprintf("Progress: %0.2f%%", gameInfo.WaveProgress*100))
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

func (dm *DisplayManager) writeString(x, y int, str string) {
	for i, r := range str {
		dm.buffer.Cells[y][x+i] = Cell{
			Symbol: r,
			BG:     lipgloss.Color("#000000"),
			FG:     lipgloss.Color("#CCCCCC"),
		}
	}
}
