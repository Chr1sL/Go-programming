package main
// this helped a lot (link)
import (
	"fmt"
	"io/ioutil"
	"time"
	"net/http"
	"log"
	"crypto/md5"
	"path/filepath"
	"strings"
)

// I'm somewhat following the tutorial from the  link in the readme
//https://golang.org/doc/articles/wiki/#tmp_6
func main() {
	//us := "pink.slips@lwhs.org"
	server := http.Server{ // makes server
		Addr: ":8080",
	}
	http.HandleFunc("/data", process) // handles student data
	server.ListenAndServe() // does the server thing idk
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func process(w http.ResponseWriter, r *http.Request) {
	sbmt_data := evt_data{r.FormValue("t_name"),r.FormValue("name"), r.FormValue("t_email"), r.FormValue("s_email"), r.FormValue("date"), r.FormValue("class"), r.FormValue("block"), r.FormValue("time")}

	t := sbmt_data.prof_name + sbmt_data.st_name + time.Now().String() // this for testing and will be needed later
	t2 := md5.New()
	t2.Write([]byte(t))
	title := fmt.Sprintf("%x", t2.Sum(nil))// makes password for teacher and makes title for related files
	st_data := &Page{title, []byte(sbmt_data.prof_name + "\n" + sbmt_data.prof_addr + "\n" + sbmt_data.st_name + "\n" + sbmt_data.st_addr + "\n" + sbmt_data.skp_class + "\n" + sbmt_data.class_block+ "\n" + sbmt_data.evt_date + "\n" + sbmt_data.out_time)} // page being prepped for saving here is what the student has submitted and the teacher will be seeing
	st_data.save() // this saves the data submitted to a txt and will make a n HTML file too
	
	// WE NEED TO SEND THIS DATA TO THE STUDENT AND THE TEACHER (FOR THE STUDENT IT IS A CONFIRMATION)

	//reset(teacher_p.Title) // deletes the relevant files eventually should be put  in if statement or another  function ||| defer forces reset() to happen the end of the program
	log.Print( title, "\n", string(st_data.Body))
	fmt.Fprintf(w, "You have successfully submitted your request!")
}

type evt_data struct {
	prof_name, st_name, prof_addr, st_addr, evt_date, skp_class, class_block, out_time string
}


type Page struct { // more or less constructor for pages in general?
	Title string
	Body  []byte
}

func (p *Page) save() (error, error) { //no input but creates a file idk i coppied lol
	file_txt := p.Title + ".txt" //info stored in txt
	file_html := p.Title +".html"
	body := strings.Replace(string(p.Body), "\n", "</p><p>", -1)
	html_content := []byte("<!DOCTYPE html><html><head><meta charset=\"utf-8\"><title>"+p.Title+"</title></head><body><p>"+body+"</p><form action=\"http://192.168.1.19:8081/prof?rsp=&id=&thread=666\" method=\"post\" enctype=\"application/x-www-form-urlencoded\"><input type=\"radio\" name=\"rsp\" value=\"yes\" required/>yes<br><input type=\"radio\" name=\"rsp\" value=\"no\" required/>no<br><table><tr><td>Computer Generated signature: </td><td><input type = \"text\" name = \"name\" required/></td></tr></table><input type=\"submit\" value=\"submit\" /></form></body></html>") //file byte

	path, _ := filepath.Abs("pink_slips/data") // set path
	return ioutil.WriteFile(filepath.Join(path, file_txt), p.Body, 0755), ioutil.WriteFile(filepath.Join(path,file_html), html_content, 0755)//writes both the record txt and the html file
}