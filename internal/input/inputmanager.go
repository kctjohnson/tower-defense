package input

// InputManager is an interface that defines the methods that an input manager should implement
type InputManager interface {
	// Example from a tic tac toe game's needs
	GetPlayerMove() (row, col int, valid bool)
}
