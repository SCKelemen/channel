package main 

import (
  "fmt"
)


func main() {
  dat := make(chan string)
  som := make(chan string)
 
  go func() { dat<-generateData()}()  
  go func() { som<-generateSome()}()

  for i := 0; i < 2; i++ {
    select {
      case data := <-dat:
        fmt.Println(data)
      case some := <-som:
        fmt.Println(some)
    }
  }

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

func generateData() string {
  return "data"
}

func generateSome() string {
  return "some"
}


type Command int 

const (
  INVALID Command = iota
  ABORT
  INCREMENT
  DECREMENT
)

type State struct {
  Value int
}

func setup(c <-chan Command, done chan<- bool, initial State) {
  state := initial
  for command := range c {
    fmt.Println(command)
  
    if command == ABORT {
      done <- true
      close(done)
      return 
    }
 
    if command == INCREMENT {
      state.Value = state.Value + 1
    }
  
    if command == DECREMENT {
      state.Value = state.Value - 1
    } 

    fmt.Printf("state: %v\n", state.Value)
  }

}



