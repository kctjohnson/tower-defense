package console

// Sample Console implementation of the DisplayManager interface

type ConsoleDisplayManager struct{}

func (c *ConsoleDisplayManager) ShowBoard(board [][]string) {
	// Print the board to the console
}

func (c *ConsoleDisplayManager) ShowTurnPrompt(player string) {
	// Print the player's turn prompt to the console
}

func (c *ConsoleDisplayManager) ShowGameResult(result string) {
	// Print the game result to the console
}
