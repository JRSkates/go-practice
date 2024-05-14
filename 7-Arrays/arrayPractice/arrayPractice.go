package main

import "fmt"

func main() {
    var practiceArray [5]string
    practiceArray = [5]string{"this", "is", "a", "practice", "array"}
    fmt.Println(practiceArray)

    practiceArray[2] = "the"
    fmt.Println(practiceArray)
}
