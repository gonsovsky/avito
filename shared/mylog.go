package shared

import (
	"fmt"
	"log"
	"os"
)

var logfile *os.File

func init() {
	logfile, _ := os.OpenFile("z_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.SetOutput(logfile)
}

func Log(v ...interface{}) {
	fmt.Println(v)
	log.Print(v)
}
