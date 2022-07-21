package game

type StateManager struct {
	currentState State
	pendingState State
}

func (manager *StateManager) SetState(s State) {
	manager.pendingState = s
}

func (manager *StateManager) UpdateState() {
	if manager.pendingState == nil {
		manager.currentState.Destroy()
		manager.currentState = manager.pendingState
		manager.currentState.Initialize()
		manager.pendingState = nil
	}
}
