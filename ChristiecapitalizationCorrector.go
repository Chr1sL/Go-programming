package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

//this function fixes the capitalization of strings
func main() {
	//gets input from user: https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text to manipulate: ")
    text, _ := reader.ReadString('\n')
    //makes all the strings in text lowercase
    text = strings.ToLower(text)
    //splits strings into individual elements in array
    stringSlice := strings.Split(text, "")

    //makes the first element in the array capital
    stringSlice[0] = strings.ToUpper(stringSlice[0])

    //for loop runs through the array of strings and finds which should
    //be capitalized
    for i := 0; i < len(stringSlice); i++ {
        //if the element comes after a period/exclamation/question mark
        //the loop will make it capitalized
        if stringSlice[i] == "." || stringSlice[i] == "!" || stringSlice[i] == "?" {
          if stringSlice[i + 1] == " " {
            stringSlice[i + 2] = strings.ToUpper(stringSlice[i + 2])
          }
        }
	  }

    //joins the elements into a sentences and prints it
    fmt.Println(strings.Join(stringSlice, ""))
}
