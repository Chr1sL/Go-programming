package main

// this is the capitalization corrector it only know periods tho no other puncuation
import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	//^^^  https://stackoverflow.com/questions/20895552/how-to-read-input-from-console-line
	// fmt.Println(make_caps(text))
	text_a := s.Split(text, ". ")
	out_a := []string{}

	for _, p := range text_a {
		part := make_caps(p)
		out_a = append(out_a, part)
	}
	out_st := s.Join(out_a, ". ")
	fmt.Print(out_st)
}

func make_caps(str string) string {
	new_string := []string{}

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
