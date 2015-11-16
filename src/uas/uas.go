package main

import (
	"functions"
	. "logs"
	"net/http"
)

func main() {
	functions.ReadConf()
	functions.CheckDb()
	http.HandleFunc("/", functions.Homepage)
	http.HandleFunc("/useradd", functions.UserAdd)
	//URL: useradd?username=<>
	http.HandleFunc("/passwd", functions.Passwd)
	//URL: passwd?username=<>&password=<>&newpass=<>&newpassagain=<>
	http.HandleFunc("/userdel", functions.UserDel)
	//URL: userdel?username=<>
	http.HandleFunc("/databackup", functions.DataBackup)
	//URL: databackup?TBD
	http.HandleFunc("/datarestore", functions.DataRestore)
	//URL: datarestore?TBD
	srv := &http.Server{Addr: "0.0.0.0:9999", Handler: nil}
	err := srv.ListenAndServe()
	if err != nil {
		Error("Start http server failed")
	}
}
