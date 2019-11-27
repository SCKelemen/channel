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
  Counter int 
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
         state.Counter = state.Counter + 1
       case DECREMENT:
         state.Value = state.Value - 1
         state.Counter = state.Counter + 1 
    }

    printState(state)
  }
}



func printState(state State) {
   fmt.Printf("clock: %v\tstate: %v\n", state.Counter,  state.Value)
}

func printCommand(command Command) {
  fmt.Printf("command: %s\n", command)
}


func merge(a State, b State) State {
  return a 
}




type CommandMux struct {
  
}


var commandMux map[Command]func(param IState)
func increment(param IState)  {
   param.SetValue(param.GetValue() + 1)
   param.SetCounter(param.GetCounter() + 1)
   //return param 
}
func decrement(param IState)   {
   param.SetValue(param.GetValue() - 1)
   param.SetCounter(param.GetCounter() + 1)

//   return param
}



func init() {
 commandMux = make(map[Command]func(param IState))
 commandMux[INCREMENT] = increment 
 commandMux[DECREMENT] = decrement 
}

type IState interface {
  GetValue() int 
  SetValue(i int) int 
  GetCounter() int 
  SetCounter(i int) int 
}


func (s State) GetValue() int  {
  return s.Value 
}

func (s State) SetValue(i int) int {
  s.Value = i
  return s.Value
}

func (s State) GetCounter() int {
  return s.Counter
}

func (s State) SetCounter(i int) int {
  s.Counter = i
  return s.Counter
}
