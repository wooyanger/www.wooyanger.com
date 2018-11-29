package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"www.wooyanger.com/pkg/logs"
)

// 定义全局变量
var (
	HttpHost	string
	HttpPort	int64

	DbHost		string
	DbPort		int64
	DbUser		string
	DbPass		string
	DbName		string

	DbDsn		string
	ListenAddr	string
	Cfg			*ini.File
	Error		error
)

// 初始化
func init() {
	Cfg, Error = ini.Load("./config/server.ini")
	if Error != nil {
		logs.Fatalf("%v", Error)
	}
	ServerSec := Cfg.Section("server")
	DatabaseSec := Cfg.Section("database")
	HttpHost = ServerSec.Key("HOST").MustString("127.0.0.1")
	HttpPort = ServerSec.Key("PORT").MustInt64(1110)
	ListenAddr = fmt.Sprintf("%s:%d", HttpHost, HttpPort)
	DbHost = DatabaseSec.Key("HOST").MustString("127.0.0.1")
	DbPort = DatabaseSec.Key("Port").MustInt64(3306)
	DbUser = DatabaseSec.Key("USER").MustString("root")
	DbPass = DatabaseSec.Key("PASS").MustString("root")
	DbName = DatabaseSec.Key("NAME").MustString("www.wooyanger.com")
	DbDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName)
}
