package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// I'm somewhat following the tutorial from the  link in the readme

func main() {
	//us := "pink.slips@lwhs.org"
	names := group_name{"my.teacher", "me"}
	//st_data :=  // student_data{us, "some_teacher@lwhs.org", "sporty.student@gmail.com", "3/14/18", "12:40", "COMPUTING 2", "teacher+student+time.now"}
	mk_data(names)
	title := "ok_test" // this for testing and will be needed later
	teacher_p := &Page{title, []byte(title)} // page being prepped for saving here is what the student has submitted and the teacher will be seeing
	teacher_p.save() // this saves
	teacher_p2, _ := loadPage(teacher_p.Title)
	defer reset(teacher_p.Title) // deletes the relevant files eventually should be put  in if statement or another  function ||| defer forces reset() to happen the end of the program
	fmt.Print(string(teacher_p2.Body))
	fmt.Print("\n--SUCCESS--\n")
	// keep --SUCCESS-- AT THE END
}

//type student_data struct { // change to an interface with functions to get this stuff vvv
//	from_addr string // this is from us we can figure that later also is a place holder will be deleted
//	prof_addr string // this for the teacher molly is a constant
//	stdt_addr string // student email address
//	evt_date  string // date of the event
//	n_time      string //? idk if this should be a string
//	skp_class string // the  class being skipped
//	grp_name  string // teacherName+studentName+time.now
//}

// that ^^^ is a struct (object/ dictionary) for all the data from the form to go to [and it works]
type data interface { // this interface makes the struct above will
//	get_st_addr() string
//	get_prof_addr() string
//	get_evt_date() string
//	get_class() string
	mk_group() string // this one should be working
} // none of the functions in the interface are defined yet so it won't work (which is why is commented)

type group_name struct {
	prof_name, st_name string
}
func (g group_name) mk_group() string {
	pn := g.prof_name
	sn := g.st_name
	t := time.Now()
	return pn + sn + t.String()
}

func mk_data(d data) (interface{}, interface{}){
	grp := d.mk_group()
	return grp, 0
}

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

func reset(group string)  error { // this function will delete resolved stuff so after the teacher says yes or no it is resolved and all related docs are purged||| also "group" is just the naming strategy i was thinking of, like teacher name + student name or sm
	file1 := group + ".txt"
	//file2 := group + ".html" //this is the page that is created when the student submits and that the teacher sees
	return os.Remove(file1) // os.Remove(file2) // does the removal
}
