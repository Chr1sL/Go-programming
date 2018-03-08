package main
// this is the capitalization corrector it doesn't know waht to do with  punctiation tho :(
import "fmt"
import "os"
import "bufio"
import s "strings"
import sc "strconv"

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter text: ")
  text, _ := reader.ReadString('\n')
  fmt.Println(text)
  fmt.Println(make_caps(text))
}


func make_caps(str string) string{
  // a = len(str)
  new_string := ""
  for i, st := range str {
    r_st := sc.QuoteRune(st)
    if i == 0 {
      new_string += s.ToUpper(string(r_st))
    } else {
      new_string += s.ToLower(string(r_st))
    }
    // else if st == "." {
    //   new_string += s.ToUpper(st)
    // }
  }
  // return s.ToUpper(str)
  return new_string
}
