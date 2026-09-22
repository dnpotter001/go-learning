package main

import "fmt"

// there is native enum type on go but you can implement the with other language features

type ServerState int //enum type

const (
	StateIdle ServerState = iota //this generate successive values automatically 0,1,2,3
	StateConnected
	StateError
	StateRetrying
)

var stateNames = map[ServerState]string{ //mapping the server state to a string
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// function for validating state transition
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknown state %s", s))
	}
}

// print the name from the enum
func (ss ServerState) String() string {
	return stateNames[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}
