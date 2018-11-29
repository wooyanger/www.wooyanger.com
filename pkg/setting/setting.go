package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"www.wooyanger.com/pkg/logs"
)

var (
	HttpHost	string
	HttpPort	int64

	ListenAddr	string
	Cfg			*ini.File
	Error		error
)

func init() {
	Cfg, Error = ini.Load("./config/server.ini")
	if Error != nil {
		logs.Fatalf("%v", Error)
	}
	ServerSec := Cfg.Section("server")
	HttpHost = ServerSec.Key("HOST").MustString("127.0.0.1")
	HttpPort = ServerSec.Key("PORT").MustInt64(1110)
	ListenAddr = fmt.Sprintf("%s:%d", HttpHost, HttpPort)
}