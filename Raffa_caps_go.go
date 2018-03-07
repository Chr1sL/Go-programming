package main

import "fmt"
import "os"
import "bufio"
import s "strings"

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  fmt.Println(text)
  fmt.Println(make_caps(text))
}

func make_caps(str string) string{
  new_string := ""
  for i, st := range str {
    if i == 0 {
      new_string = new_string + s.ToUpper(st string)
    }
    elif == "." {
      new_string = new_string + s.ToUpper(st string)
    }
    else {
      new_string += st
    }
  }
  // return s.ToUpper(str)
  return new_string
}
