package main

// I'm somewhat following the tutorial from the  link in the readme
import (
	"fmt"
	// "io/ioutil"
)

func main() {
	us := "pink.slips@lwhs.org"
	dflt_msg := "this is a test default message"
	test_mail := email{us, "some_teacher@lwhs.org", "student@gmail.com", dflt_msg, "3/14/18", "1:30", "COMPUTING 2"}
	fmt.Print(test_mail)
}

type email struct {
	from_addr string // this is from us we can figure that later
	prof_addr string // this for the teacher molly is a constant
	stdt_addr string // student email address
	msg_cntnt string // !this must be enterpritted as a string no matter what! also we should develope a default message
	evt_date  string // date of the event
	time      string //? idk if this should be a string
	skp_class string // the  class being skipped
}

// that ^^^ is a struct (object/ dictionary) for all the data from the form to go to {and it works}
