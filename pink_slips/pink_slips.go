package main
// this helped a lot (link)
import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"net/http"
	"crypto/md5"
)

// I'm somewhat following the tutorial from the  link in the readme

func main() {
	//us := "pink.slips@lwhs.org"
	names := evt_data{"sfdsf", "me", "teacher@lwhs.org", "st@gmail.com", "now", "this.class"}
	//st_data :=  // student_data{us, "some_teacher@lwhs.org", "sporty.student@gmail.com", "3/14/18", "12:40", "COMPUTING 2", "teacher+student+time.now"}
	mk_data(names)
	title := "ok_test" // this for testing and will be needed later
	teacher_p := &Page{title, []byte(title)} // page being prepped for saving here is what the student has submitted and the teacher will be seeing
	teacher_p.save() // this saves
	teacher_p2, _ := loadPage(teacher_p.Title)
	reset(teacher_p.Title) // deletes the relevant files eventually should be put  in if statement or another  function ||| defer forces reset() to happen the end of the program
	fmt.Print(string(teacher_p2.Body))
	fmt.Print("\n--SUCCESS--\n")
	// keep --SUCCESS-- AT THE END
	server := http.Server{ // makes server
		Addr: "127.0.0.1",
	}
	http.HandleFunc("/data", process) // handles student data and produces
	server.ListenAndServe() // does the server thing idk
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.ListenAndServe(":3000", nil)
	//http.Handle("/", http.FileServer(http.Dir("./index.html")))
	//http.ListenAndServe(":3001", nil)
}
func process(w http.ResponseWriter, r *http.Request) {
	sbmt_data := evt_data{r.FormValue("t_name"),r.FormValue("name"), r.FormValue("t_email"), r.FormValue("s_email"), r.FormValue("date"), r.FormValue("class")}
	
	t := sbmt_data.prof_name + sbmt_data.st_name + time.Now().String() // this for testing and will be needed later
	title, _ := md5.New().Write([]byte(t)) // makes password for teacher and makes title for related files
	teacher_p := &Page{string(title), []byte(sbmt_data.prof_name + "\n" + sbmt_data.prof_addr + "\n" + sbmt_data.st_name + "\n" + sbmt_data.st_addr + "\n" + sbmt_data.skp_class + "\n" + sbmt_data.evt_date) } // page being prepped for saving here is what the student has submitted and the teacher will be seeing
	teacher_p.save() // this save
	//teacher_p2, _ := loadPage(teacher_p.Title)
	//reset(teacher_p.Title) // deletes the relevant files eventually should be put  in if statement or another  function ||| defer forces reset() to happen the end of the program
	
	//fmt.Print(sbmt_data.st_name, sbmt_data.evt_date)
	fmt.Fprintf(w, "You have successfully submitted your request!")
}
type data interface { // this interface makes the struct above will
//	get_st_addr() string
//	get_prof_addr() string
//	get_evt_date() string
//	get_class() string
//	get_grade() string
//	get_team() string
	mk_group() string // this one should be working
} // none of the functions in the interface are defined yet so it won't work (which is why is commented)

type evt_data struct {
	prof_name, st_name, prof_addr, st_addr, evt_date, skp_class string
}

func (e evt_data) mk_group() string {
	pn := e.prof_name
	sn := e.st_name
	t := time.Now()
	return pn + sn + t.String()
}

func mk_data(d data) (interface{}){
	grp := d.mk_group()
	return grp
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
