package state

type StateMachine struct {
	states []*State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

// func (s StateMachine) ChangeState(state *state.State) {

// }

func (s *StateMachine) PushState(state State) {
	state.Init(s)
	s.states = append(s.states, &state)
}

func (s *StateMachine) PopState() (*State, bool) {
	stackLength := len(s.states)
	if stackLength == 0 {
		return nil, false
	}

	index := stackLength - 1
	el := s.states[index]
	s.states = s.states[:index]
	return el, true
}

func (s StateMachine) CurrentState() *State {
	stackLength := len(s.states)
	if stackLength == 0 {
		return nil
	}
	return s.states[stackLength-1]
}
