package db

import (
	"os/exec"
	"strings"
	. "uas/logs"
)

var BindDn string
var BindPwd string
var BaseDn string
var GroupSuffix string
var UserSuffix string

func commExec(cmd string, arg ...string) (output string, err error) {
	output = ""
	Debug(cmd, arg)
	stdout, err := exec.Command(cmd, arg...).Output()
	Debug("db.commExec: stdout:", stdout)
	if err != nil {
		Debug("db.commExec: Command() error:", err)
	} else {
		output = string(stdout)
	}
	return output, err
}

func whoami(username, password string) (who string, err error) {
	Debug("db.whoami: username:", username,
		", passwd: ", password)
	command := "ldapwhoami"
	arg1 := "-x"
	arg2 := "-w" + password
	arg3 := "-Duid=" + username + "," + UserSuffix + "," + BaseDn
	who, err = commExec(command, arg1, arg2, arg3)
	who = strings.Trim(who, "\n")
	return who, err
}

func ldapsetpasswd(username, password string) (res string, err error) {
	Debug("db.ldapsetpasswd: username:", username,
		", passwd: ", password)
	command := "ldapsetpasswd"
	res, err = commExec(command, username, password)
	res = strings.Trim(res, "\n")
	return res, err
}

//func ldapadduser(username string) (res string, err error) {
//  Debug("db.ldapuseradd: username:", username)
//	command := "ldapadduser"
//}

func UserAdd(username string) string {
	Debug("db.UserAdd: username: ", username)
	res := "OK"
	return res
}

func PasswordCheck(username, password string) string {
	Debug("db.PasswordCheck: username:", username,
		", passwd: ", password)
	res := "OK"
	dn, err := whoami(username, password)
	Debug("db.PasswordCheck: dn1:", dn)
	if err != nil {
		if err.Error() == "exit status 49" {
			res = "Current password check failed"
		} else {
			res = "Unknown error"
		}
	} else {
		if dn == "dn:uid="+username+","+UserSuffix+","+BaseDn {
			res = "OK"
		} else {
			res = "Current password check failed"
		}
	}
	return res
}

func Passwd(username, password string) string {
	Debug("db.Passwd: username: ", username,
		", passwd: ", password)
	res := "OK"
	msg, err := ldapsetpasswd(username, password)
	Debug("db.Passwd: msg:", msg)
	if err != nil {
		Error("ldapsetpasswd failed")
		res = "ldapsetpasswd failed"
	} else {
		if strings.Contains(msg, "Successfully set encoded password for user uid="+username) {
			res = "OK"
		} else {
			res = "Password change failed"
		}
	}
	return res
}
