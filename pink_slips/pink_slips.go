package main
// this helped a lot (link)
import (
	"fmt"
	"io/ioutil"
	//"os"
	"time"
	"net/http"
)

// I'm somewhat following the tutorial from the  link in the readme

func main() {
	//us := "pink.slips@lwhs.org"
	//fmt.Print(string(teacher_p2.Body))
	fmt.Print("\n--SUCCESS--\n")
	// keep --SUCCESS-- AT THE END
	server := http.Server{ // makes server
		Addr: "127.0.0.1:8080",
	}
	fmt.Print()
	http.HandleFunc("/data", process) // handles student data
	server.ListenAndServe() // does the server thing idk
}
func process(w http.ResponseWriter, r *http.Request) {
	sbmt_data := evt_data{r.FormValue("t_name"),r.FormValue("name"), r.FormValue("t_email"), r.FormValue("s_email"), r.FormValue("date"), r.FormValue("class")}
	
	title := sbmt_data.prof_name + sbmt_data.st_name + time.Now().String() // this for testing and will be needed later
	teacher_p := &Page{title, []byte(sbmt_data.prof_name + "\n" + sbmt_data.prof_addr + "\n" + sbmt_data.st_name + "\n" + sbmt_data.st_addr + "\n" + sbmt_data.skp_class + "\n" + sbmt_data.evt_date)} // page being prepped for saving here is what the student has submitted and the teacher will be seeing
	teacher_p.save() // this saves
	//teacher_p2, _ := loadPage(teacher_p.Title)
	//reset(teacher_p.Title) // deletes the relevant files eventually should be put  in if statement or another  function ||| defer forces reset() to happen the end of the program
	
	//fmt.Print(sbmt_data.st_name, sbmt_data.evt_date)
	fmt.Fprintf(w, "You have successfully submitted your request!")
}

type evt_data struct {
	prof_name, st_name, prof_addr, st_addr, evt_date, skp_class string
}

type Page struct { // more or less constructor for pages in general?
	Title string
	Body  []byte
}

func (p *Page) save() error { //no input but creates a file idk i coppied lol
	filename := p.Title + ".txt"                    //info stored in txt
	return ioutil.WriteFile(filename, p.Body, 0600) // makes the txt file
}

//func loadPage(title string) (*Page, error) {
//	filename := title + ".txt"
//	body, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return nil, err
//	}
//	return &Page{Title: title, Body: body}, nil
//}

//func reset(group string)  error { // this function will delete resolved stuff so after the teacher says yes or no it is resolved and all related docs are purged||| also "group" is just the naming strategy i was thinking of, like teacher name + student name or sm
//	file1 := group + ".txt"
//	//file2 := group + ".html" //this is the page that is created when the student submits and that the teacher sees
//	return os.Remove(file1) // os.Remove(file2) // does the removal
//}
