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

  commands := make(chan Command)
  done := make(chan bool)
  go setup(commands, done)
  
  commands <- HELLO
  commands <- GOODBYE
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
  HELLO
  GOODBYE
)


func setup(c <-chan Command, done chan<- bool) {
  for command := range c {
    fmt.Println(command)
    if command == ABORT {
      done <- true
      close(done)
      return 
    }
  }
}



