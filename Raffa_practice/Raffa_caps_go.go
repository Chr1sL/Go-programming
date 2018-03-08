package main
// this is the capitalization corrector it doesn't know waht to do with  punctiation tho :(
import "fmt"
import "os"
import "bufio"
import s "strings"
// import sc "strconv"
// import b "bytes"

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  // fmt.Println(text)
  fmt.Println(make_caps(text))
}

func make_caps(str string) string{
  new_string := []string {}

  for i, st := range str {
    // r_st := sc.QuoteRune(st)
    if i == 0 {
      new_string = append(new_string, s.ToUpper(string(st)))
    } else {
      new_string = append(new_string, s.ToLower(string(st)))
    }
  }
  final_st := s.Join(new_string, "")
  return final_st
}
