package Temp

import (
	"log"
	"github.com/quexer/utee"
)

func Println(str string) {
	log.Println("print",str)
	utee.Chk(nil)
}
