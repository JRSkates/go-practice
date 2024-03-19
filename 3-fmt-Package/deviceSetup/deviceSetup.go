package main

import (
  "fmt"
)

func main() {
  var name string

  fmt.Println("What is your name?")
  fmt.Scan(&name)

  fmt.Printf("Hello %v, nice to meet you!", name)
  fmt.Println(" ")

  var age int

  fmt.Println("What is your age?")
  fmt.Scan(&age)

  fmt.Printf("You are %d years old", age)
  fmt.Println(" ")

  newMessage := fmt.Sprintf("So your name is %v and you are %d years old, good to know!", name, age)

  fmt.Println(newMessage)
}