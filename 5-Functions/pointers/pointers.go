package main

import "fmt"

func main() {
	star := "Polaris"
  planetNumber := 3
	
	starAddress := &star
  fmt.Println("The address of star is", starAddress)
  
  var planetPointer *int
  planetPointer = &planetNumber
  fmt.Println("The address of planetNumber is", planetPointer)
}
