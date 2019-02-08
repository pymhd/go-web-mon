package main

import (
	"sync"
	"time"
)

type State struct {
	isActive bool
	when     time.Time
}

type Map map[string]*State

type StateKeeper struct {
	sync.Mutex
	Map
}

func (s *StateKeeper) getState(k string) State {
	s.Lock()
	defer s.Unlock()

	return *s.Map[k]
}

func (s *StateKeeper) setState(k string, b bool, w time.Time) {
	s.Lock()
	defer s.Unlock()

	s.Map[k].isActive = b
	s.Map[k].when = w
}

func CreateKeeper(wrs ...WebResource) *StateKeeper {
	Keeper := new(StateKeeper)
	m := make(Map, 0)
	for _, w := range wrs {
		m[w.Name] = new(State)
	}
	Keeper.Map = m
	return Keeper
}
