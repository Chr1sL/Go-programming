package main
// this is the capitalization corrector it doesn't know waht to do with  punctiation tho :(
import "fmt"
import "os"
import "bufio"
import s "strings"

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  //^^^  https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
  fmt.Println(make_caps(text))
}

func make_caps(str string) string{
  new_string := []string {}

  for i, st := range str {
    if i == 0 {
      new_string = append(new_string, s.ToUpper(string(st)))
    } else {
      new_string = append(new_string, s.ToLower(string(st)))
    }
  }
  final_st := s.Join(new_string, "")
  return final_st
}
