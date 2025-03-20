package display

// DisplayManager is an interface that defines the methods that a display manager should implement
type DisplayManager interface {
	// Example from a tic tac toe game's needs
	ShowBoard(board [][]string)
	ShowTurnPrompt(player string)
	ShowGameResult(result string)
}
