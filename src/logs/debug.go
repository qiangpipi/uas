package logs

import (
	. "fmt"
)

var D *bool

func Debug(debug ...interface{}) {
	if *D {
		Println(info)
	}
}
