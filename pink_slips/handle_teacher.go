package main

import (
	"net/http"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"os"
	"strings"
	"log"
	"net/smtp"
	"crypto/tls"
)

func main () {
	server := http.Server{ // makes server
		Addr: ":8081",
	}
	http.HandleFunc("/prof", process_rsp) // handles student data
	server.ListenAndServe() // does the server thing idk
}

func process_rsp(w http.ResponseWriter, r *http.Request)  {
	// SEND THE STUDENT AN EMAIL THAT HAS THE TEACHER'S RESPONSE
	prof_rsp := rsp{r.FormValue("rsp"), r.FormValue("id")}
	f_data_arr := prof_rsp.log_rsp()
	//prof_rsp.log_rsp()
	defer prof_rsp.prune() // forces the pruning/ deletion of irrelavent files to happen last so it doesnt happen too soon
	final_data := all_data{f_data_arr[0], f_data_arr[1], f_data_arr[2], f_data_arr[3], f_data_arr[4], f_data_arr[5], f_data_arr[6], f_data_arr[7], f_data_arr[8]} // the final set of data  with the teacher's response
	final_data.send_mail()
	
	fmt.Fprint(w, "Thank you for your response it has been recorded and your student will be notified shortly")
}

type rsp struct {
	reply, auth_id string
}

type all_data struct {
	prof_name, st_name, prof_addr, st_addr, evt_date, skp_class, class_block, out_time, reply string
}

func (r *rsp)log_rsp() []string {
	file_txt := r.auth_id + ".txt"
	path, _ := filepath.Abs("pink_slips/data")
	data, _ := ioutil.ReadFile(filepath.Join(path, file_txt))
	//if err != nil {
	//	panic(err)
	//		fmt.Fprint(nil, "Oh no! this key is incorrect! Please try again.")
	//}
	new_data := string(data) + "\n" + r.reply + "\n" + r.auth_id
	log_file, _ := os.OpenFile(filepath.Join(path, "log.txt"), os.O_APPEND|os.O_WRONLY, 0600)

	defer log_file.Close()
	log_file.WriteString(new_data + "\n")
	//return strings.Split(new_data, "\n")
	return strings.Split(new_data, "\n")
}

func (r *rsp) prune() (error, error) {
	file_txt, file_html := r.auth_id + ".txt", r.auth_id + ".html"
	path, _ := filepath.Abs("pink_slips/data")
	return os.Remove(filepath.Join(path, file_txt)), os.Remove(filepath.Join(path, file_html))
}

func (ad *all_data) send_mail() {
	mail := Mail{}
	mail.senderId = "lwhs.pinkslips@gmail.com"
	mail.toIds = []string{ad.st_addr, ad.prof_addr} //TODO add the email for molly french
	mail.subject = "Request to miss class on " + ad.evt_date
	mail.body = strings.Join([]string{"Hello, \n \n Here is your anticipated absence receipt:", ad.st_name, ad.class_block, ad.skp_class, ad.evt_date, "teacher's reply: "+ ad.reply}, "\n")
	
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
