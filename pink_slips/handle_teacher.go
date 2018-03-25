package main

import (
	"net/http"
	"fmt"
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
	
	fmt.Fprintf(w, "Thank you for your response it has been recorded and sent to your student!")
}

type rsp struct {
	reply, auth_id string
}

func (r *rsp)log_rsp(){
	file_txt := r.auth_id + "txt"
	f, err := os.OpenFile(file_txt, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	text := r.reply
	if _, err = f.WriteString("\n"+text); err != nil {
		panic(err)
	}
}

// make func prune()