package statemachine

import "github.com/hpdobrica/starbreach/state"

type StateMachine struct {
	states []*state.State
}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

// func (s StateMachine) ChangeState(state *state.State) {

// }

func (s *StateMachine) PushState(state state.State) {
	s.states = append(s.states, &state)
}

func (s *StateMachine) PopState() (*state.State, bool) {
	stackLength := len(s.states)
	if stackLength == 0 {
		return nil, false
	}

	index := stackLength - 1
	el := s.states[index]
	s.states = s.states[:index]
	return el, true
}

func (s StateMachine) CurrentState() *state.State {
	stackLength := len(s.states)
	if stackLength == 0 {
		return nil
	}
	return s.states[stackLength-1]
}
