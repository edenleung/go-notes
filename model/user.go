package library

import (
	"reflect"
	"strings"
	"xiaodim.com/xiaodi/project/library"

	//"xiaodim.com/xiaodi/project/library"
)

type UserInfo struct {
	Nickname string `db:"nickname"`
}

func (ui *UserInfo) fieldAndValues() (fields string, values []interface{}) {

	tl := reflect.TypeOf(ui).Elem()
	vl := reflect.ValueOf(ui).Elem()

	arr := []string{}
	for i :=0; i< tl.NumField(); i++ {
		arr = append(arr, tl.Field(i).Tag.Get("db"))
		values = append(values, vl.Field(i).Addr().Interface())
	}

	fields = "`"+ strings.Join(arr, "`,`") +"`"
	return
}

func GetUserInfo(id int64) (info UserInfo, err error){
	Db := library.Connect()
	fields, values := info.fieldAndValues()
	rows, err := Db.Query("select "+ fields +" from user where id=?", id)
	defer Db.Close()
	if err != nil {
		return
	}

	if rows.Next() {
		err = rows.Scan(values...)
	}

	return
}