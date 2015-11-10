package functions

import (
	"db"
	"io/ioutil"
	. "logs"
	"net/http"
	"os"
	"strings"
)

func readHomepage(fn string) []byte {
	var fd []byte
	file, err := os.Open(fn)
	defer file.Close()
	if err != nil {
		Error("function.Homepage: fail to open main page")
		fd = []byte("Fail to open main page")
	} else {
		fd, _ = ioutil.ReadAll(file)
	}
	return fd
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	//Security check
	//Reponse the home page
	fd := readHomepage("data/main.html")
	w.Write(fd)
}
func UserAdd(w http.ResponseWriter, r *http.Request) {
	//useradd?username=<>
	//Parse username
	username := ""
	if err := r.ParseForm(); err != nil {
	}
	for k, v := range r.Form {
		Debug("functions.UserAdd: ", k, " = ", v)
		if k == "username" {
			username = strings.Join(v, "")
		}
	}
	if username == "" || len(r.Form) > 1 {
		//Response error paras
		Debug("functions.UserAdd: username = ", username)
		Error("Bad para")
	} else {
		//Add user in database
		Info("Adding user in database")
		if err := db.UserAdd(username); err != nil {
		} else {
			//Response to client
			w.Write([]byte("OK"))
		}
	}
}

func Passwd(w http.ResponseWriter, r *http.Request) {
	//passwd?username=<>&password=<>&newpass=<>&newpassagain=<>
	//Parse username=<>&password=<>&newpass=<>&newpassagain=<>
	username := ""
	password := ""
	newpass := ""
	newpassagain := ""
	fd := readHomepage("data/main.html")
	if err := r.ParseForm(); err != nil {
	}
	for k, v := range r.Form {
		Debug(k, " = ", v)
		if k == "username" {
			username = strings.Join(v, "")
		}
		if k == "password" {
			password = strings.Join(v, "")
		}
		if k == "newpass" {
			newpass = strings.Join(v, "")
		}
		if k == "newpassagain" {
			newpassagain = strings.Join(v, "")
		}
	}
	w.Write(fd)
	if len(r.Form) != 4 {
		Debug(len(r.Form))
		w.Write([]byte("<font color=\"red\">Bad para and please try again</font>"))
	} else if username == "" {
		w.Write([]byte("<font color=\"red\">Username IS blank and please try again</font>"))
	} else if password == "" {
		w.Write([]byte("<font color=\"red\">Password IS blank and please try again</font>"))
	} else if newpass == "" {
		w.Write([]byte("<font color=\"red\">Newpass IS blank and please try again</font>"))
	} else if newpass != newpassagain {
		w.Write([]byte("<font color=\"red\">Newpass not equal newpassagain and please try again</font>"))
	} else {
		//Change password in database
		Debug("Changing password in database")
		passcheck := db.PasswordCheck(username, password)
		if passcheck == "OK" {
			changecheck := db.Passwd(username, newpass)
			if changecheck == "OK" {
				w.Write([]byte("<font color=\"green\">Password changed for <strong>" + username + "</strong></font>"))
			} else {
				w.Write([]byte("<font color=\"red\">" + changecheck + "</strong></font>"))
			}
		} else if passcheck == "Current password check failed" {
			w.Write([]byte("<font color=\"red\">" + passcheck + " <strong>" + username + "</strong></font>"))
		} else {
			w.Write([]byte("<font color=\"red\">Unknown error!</strong></font>"))
		}
	}
}

func UserDel(w http.ResponseWriter, r *http.Request) {
	//userdel?username=<>
	//Parse username
}

func DataBackup(w http.ResponseWriter, r *http.Request) {
	//databackup?TBD
	//Parse TBD
}

func DataRestore(w http.ResponseWriter, r *http.Request) {
	//datarestore?TBD
	//Parse TBD
}
