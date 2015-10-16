package main

import (
	"functions"
	. "logs"
	"net/http"
)

func main() {
	http.HandleFunc("/useradd", functions.UserAdd)
	http.HandleFunc("/passwd", functions.Passwd)
	http.HandleFunc("/userdel", functions.UserDel)
	http.HandleFunc("/databackup", functions.DataBackup)
	http.HandleFunc("/datarestore", functions.DataRestore)
	srv := &http.Server{Addr: "0.0.0.0:8888", Handler: nil}
	err := srv.ListenAndServe()
	if err != nil {
		Error("Start http server failed")
	}
}
