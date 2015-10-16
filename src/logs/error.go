package logs

import (
	. "fmt"
)

func Error(err ...interface{}) {
	Println(err)
}
