package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)
func init() {
	var err error
	Cfg, err = ini.Load("dev.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'app.ini': %v", err)
	}
}



