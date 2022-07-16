package game

var gameState *Game = nil

func GetGameState() (state *Game) {
	return gameState
}

func SetGameState(state *Game) {
	gameState = state
}
