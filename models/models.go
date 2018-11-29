package models

import (
	"github.com/go-xorm/xorm"
	"www.wooyanger.com/pkg/logs"
	"www.wooyanger.com/pkg/setting"
)

var (
	x		*xorm.Engine
	e		error
	tables	[]interface{}
)

func init()  {
	tables = append(tables, new(User))
	x, e = xorm.NewEngine("mysql", setting.DbDsn)
	if e != nil {
		logs.Fatalf("%v", e)
	}
	if e = x.Ping();e != nil {
		logs.Fatalf("%v", e)
	}
}