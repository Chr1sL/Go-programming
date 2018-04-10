package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"net/http"
	"log"
	"crypto/md5"
	"path/filepath"
	"strings"
	"net/smtp"
	_ "bytes"
	"crypto/tls"
)

func main() {
	server := http.Server{ // makes server
		Addr: ":8080",
	}
	http.HandleFunc("/data", process) // handles student data
	server.ListenAndServe() // does the server thing idk
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func process(w http.ResponseWriter, r *http.Request) {
	sbmt_data := evt_data{r.FormValue("t_name"),r.FormValue("name"), r.FormValue("t_email"), r.FormValue("s_email"), r.FormValue("date"), r.FormValue("class"), r.FormValue("block"), r.FormValue("time")}

		if strings.HasSuffix(sbmt_data.prof_addr, "@lwhs.org") == true { // prevents students from sending email to self
			t := sbmt_data.prof_name + sbmt_data.st_name + time.Now().String() // this for testing and will be needed later
			t2 := md5.New()
			t2.Write([]byte(t))
			title := fmt.Sprintf("%x", t2.Sum(nil))// makes password for teacher and makes title for related files
			st_data := &Page{title, []byte(sbmt_data.prof_name + "\n" + sbmt_data.prof_addr + "\n" + sbmt_data.st_name + "\n" + sbmt_data.st_addr + "\n" + sbmt_data.skp_class + "\n" + sbmt_data.class_block+ "\n" + sbmt_data.evt_date + "\n" + sbmt_data.out_time)} // page being prepped for saving here is what the student has submitted and the teacher will be seeing
			st_data.save() // this saves the data submitted to a txt and will make a n HTML file too

			defer sbmt_data.send_mail("Hello " + sbmt_data.prof_name +", \n \n I have a sports competition on " + sbmt_data.evt_date + " and was hoping it would be ok if I missed class/ left early that day (I leave at "+ sbmt_data.out_time + "). Of course I will make up any course material that I missed in class. \n \n Best wishes, \n \t"+ sbmt_data.st_name + "\n \n http://192.168.1.22:3000/data/", title+".html")
			defer sbmt_data.send_mail(strings.Join([]string{"Hello, \n \n Here is your anticipated absence receipt:", sbmt_data.st_name, sbmt_data.class_block, sbmt_data.skp_class, sbmt_data.evt_date}, "\n"), "")

			log.Print( title, "\n", string(st_data.Body))
			fmt.Fprint(w, "You have successfully submitted your request!")
			} else {
				fmt.Fprint(w, "Oh no! The teacher email address you submitted was not recognized! Please go back and try again.")
		}

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
	body := strings.Replace(string(p.Body), "\n", "</p><p>", -1) // "http://[your ip addr]/prof?r
	html_content := []byte("<!DOCTYPE html><html><head><meta charset=\"utf-8\"><title>"+p.Title+"</title></head><body><p>"+body+"</p><form action=\"http://192.168.1.22:8081/prof?rsp=&id=&thread=666\" method=\"post\" enctype=\"application/x-www-form-urlencoded\"><input type=\"radio\" name=\"rsp\" value=\"yes\" required/>yes<br><input type=\"radio\" name=\"rsp\" value=\"no\" required/>no<br><table><tr><td>Computer Generated signature: </td><td><input type = \"text\" name = \"id\" required/></td></tr></table><input type=\"submit\" value=\"submit\" /></form></body></html>") //file byte

	path, _ := filepath.Abs("pink_slips/data") // set path
	return ioutil.WriteFile(filepath.Join(path, file_txt), p.Body, 0755), ioutil.WriteFile(filepath.Join(path,file_html), html_content, 0755)//writes both the record txt and the html file
}
// blue print for the code bellow came from https://hackernoon.com/golang-sendmail-sending-mail-through-net-smtp-package-5cadbe2670e0
func (e *evt_data) send_mail(body, file_addr string) {
	mail := Mail{}
	mail.senderId = "lwhs.pinkslips@gmail.com"
	mail.toIds = []string{e.st_addr, e.prof_addr}
	mail.subject = "Request to miss class on " + e.evt_date
	mail.body = body + file_addr
	messageBody := mail.BuildMessage()
	
	smtpServer := SmtpServer{host: "smtp.gmail.com", port: "465"}
	
	log.Println(smtpServer.host)
	//build an auth
	auth := smtp.PlainAuth("", mail.senderId, "4153334021", smtpServer.host)
	
	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName: smtpServer.host,
	}
	
	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}
	
	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}
	
	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}
	
	// step 2: add all from and to
	if err = client.Mail(mail.senderId); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}
	
	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}
	
	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}
	
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	
	client.Quit()
	
	log.Println("Mail sent successfully")
}

type SmtpServer struct {
	host string
	port string
}

func (s *SmtpServer) ServerName() string {
	return s.host + ":" + s.port
}

type Mail struct {
	senderId string
	toIds    []string
	subject  string
	body     string
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderId)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}
	
	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body
	
	return message
}
