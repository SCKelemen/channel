package main

import "fmt"

func main() {
 fmt.Println("sanity check")

 salts := make(chan int)
 ids := make(chan int)

 go salt(salts)
 go id(ids)

 salt, id := <-salts, <-ids

 fmt.Println(salt, id, salt+id)
}


func salt(c chan int) {
 c <- 42
}

func id(c chan int) {
 c <- 100 
}
