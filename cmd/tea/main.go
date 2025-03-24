package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"ecstemplate/internal/game"
	"ecstemplate/internal/game/ui/teaui"
)

type TickMsg struct {
	Time time.Time
}

type GameModel struct {
	game      *game.Game
	lastTick  time.Time
	frameRate time.Duration
}

func NewGameModel() *GameModel {
	g := game.NewGame()
	g.Initialize(80, 10)
	return &GameModel{
		game:      g,
		lastTick:  time.Now(),
		frameRate: time.Second / 60,
	}
}

func (m GameModel) tick() tea.Cmd {
	return tea.Tick(m.frameRate, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

func (m GameModel) Init() tea.Cmd {
	return m.tick()
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		// Calculate time delta
		delta := msg.Time.Sub(m.lastTick).Seconds()
		m.lastTick = msg.Time

		// Update the game
		m.game.Update(delta)

		return m, m.tick()

	case tea.KeyMsg:
		// Queue the key in the input manager
		inputManager := m.game.GetInputManager().(*teaui.InputManager)
		inputManager.QueueKey(msg.String())

		// Check for quit
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		return m, nil

	case tea.WindowSizeMsg:
		displayManager := m.game.GetDisplayManager().(*teaui.DisplayManager)
		displayManager.Resize(msg.Width, msg.Height)
	}

	return m, nil
}

func (m GameModel) View() string {
	displayManager := m.game.GetDisplayManager().(*teaui.DisplayManager)
	buffer := displayManager.GetBuffer()
	return buffer.String()
}

func main() {
	p := tea.NewProgram(NewGameModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
