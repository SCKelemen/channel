package main 

import (
  "fmt"
 "strconv"
)


func main() {
  initial := State{Value: 0}
  commands := make(chan Command)
  done := make(chan bool)
  go setup(commands, done, initial)
  
  commands <- INCREMENT
  commands <- INCREMENT
  commands <- INCREMENT
  commands <- DECREMENT
  commands <- INVALID
  commands <- ABORT
  
  <-done 
}



type Command int 

const (
  INVALID Command = iota
  ABORT
  INCREMENT
  DECREMENT
)


var commandss = [...]string{
	INVALID:     "ILLEGAL",
	ABORT:       "ABORT",
	INCREMENT:   "INCREMENT",
	DECREMENT:   "DECREMENT",
}

func (command Command) String() string {
	s := ""
	if 0 <= command && command < Command(len(commandss)) {
		s = commandss[command]
	}
	if s == "" {
		s = "command(" + strconv.Itoa(int(command)) + ")"
	}
	return s
}

type State struct {
  Value int
}

func setup(c <-chan Command, done chan<- bool, initial State) {
  state := initial
  printState(state)

  for command := range c {

    printCommand(command)

    switch command {
      case ABORT:
         done <- true
         close(done)
         return 
       case INCREMENT:
         state.Value = state.Value + 1
       case DECREMENT:
         state.Value = state.Value - 1
     }

    printState(state)
  }
}



func printState(state State) {
   fmt.Printf("state: %v\n", state.Value)
}

func printCommand(command Command) {
  fmt.Printf("command: %s\n", command)
}
