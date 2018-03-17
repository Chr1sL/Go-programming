package main

// I'm somewhat following the tutorial from the  link in the readme
import (
	"fmt"
	"io/ioutil"
)

func main() {
	us := "pink.slips@lwhs.org"
	test_mail := email{us, "some_teacher@lwhs.org", "student@gmail.com", "3/14/18", "1:30", "COMPUTING 2"}
	teacher_p := &Page{"s_2_t_Title", []byte(test_mail.evt_date + "\n"+  test_mail.from_addr + "\n"+  test_mail.prof_addr + "\n"+  test_mail.skp_class + "\n"+  test_mail.time)} // bage being prepped for saving here is what the student has submitted and the teacher will be seeing
	teacher_p.save() // this saves
	teacher_p2, _ := loadPage(teacher_p.Title)
	fmt.Print(string(teacher_p2.Body), "\n--SUCCESS--\n")// keep --SUCCESS-- AT THE END
}

type email struct {
	from_addr string // this is from us we can figure that later
	prof_addr string // this for the teacher molly is a constant
	stdt_addr string // student email address
	evt_date  string // date of the event
	time      string //? idk if this should be a string
	skp_class string // the  class being skipped
}

// that ^^^ is a struct (object/ dictionary) for all the data from the form to go to {and it works}

// func get_info() struct {
//}
type Page struct { // more or less constructor for pages in general?
	Title string
	Body  []byte
}

func (p *Page) save() error { //no input but creates a file idk i coppied lol
	filename := p.Title + ".txt"                    //info stored in txt
	return ioutil.WriteFile(filename, p.Body, 0600) // makes the txt file
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
