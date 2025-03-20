package console

// Sample Console implementation of the InputManager interface

type ConsoleInputManager struct{}

func (c *ConsoleInputManager) GetPlayerMove() (row, col int, valid bool) {
	// Get the player's move from the console
	return 0, 0, true
}
