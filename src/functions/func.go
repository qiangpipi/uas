package functions

import (
	. "logs"
	"net/http"
)

func UserAdd(w http.ResponseWriter, r *http.Request) {
	//UserAdd username
	//Parse username
	//user string
}

func Passwd(w http.ResponseWriter, r *http.Request) {
	//Passwd username/current password/new password/password again
	//Parse username/current password/new password/password again
	//user string
	//passwd string
	//newpass string
	//confirmpass string
}

func UserDel(w http.ResponseWriter, r *http.Request) {
	//UserDel username
	//Parse username
	//user string
}

func DataBackup(w http.ResponseWriter, r *http.Request) {
}

func DataRestore(w http.ResponseWriter, r *http.Request) {
}
