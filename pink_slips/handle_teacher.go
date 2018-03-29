package main

import (
	"net/http"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"os"
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
	prof_rsp.log_rsp()
	defer prof_rsp.prune() // forces the pruning/ deletion of irrelavent files to happen last so it doesnt happen too soon
	// EMAIL STUDENT WITH RESPONSE OF THE TEACHER (YES/ NO)
	fmt.Fprintf(w, "Thank you for your response it has been recorded and your student  will be notified shortly!")
}

type rsp struct {
	reply, auth_id string
}

func (r *rsp)log_rsp() {
	file_txt := r.auth_id + ".txt"
	path, _ := filepath.Abs("pink_slips/data")
	data, err := ioutil.ReadFile(filepath.Join(path, file_txt))
	if err != nil {
		panic(err)
			fmt.Fprint(nil, "Oh no! this key is incorrect! Please try again.")
	}
	new_data := string(data) + "\n" + r.reply + "\n" + r.auth_id
	log_file, _ := os.OpenFile(filepath.Join(path, "log.txt"), os.O_APPEND|os.O_WRONLY, 0600)

	defer log_file.Close()
	log_file.WriteString(new_data)
}

func (r *rsp) prune() (error, error) {
	file_txt, file_html := r.auth_id + ".txt", r.auth_id + ".html"
	path, _ := filepath.Abs("pink_slips/data")
	return os.Remove(filepath.Join(path, file_txt)), os.Remove(filepath.Join(path, file_html))
}