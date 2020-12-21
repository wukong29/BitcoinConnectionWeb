package models

import (
	"BitcoinConnectionWeb/db_mysql"
	"BitcoinConnectionWeb/util"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

func (u User) SaveUser() (int64, error) {
	//1、密码脱敏处理（hash加密）
	u.Password = util.MD5HashString(u.Password)
	//执行数据库操作
	fmt.Println("将要保存的手机号码：", u.Phone, "密码：", u.Password)
	result, err := db_mysql.Db.Exec("insert into user(phone, password)"+"values(?,?)",
		u.Phone, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//查询用户信息
func (u User) QueryUser() (*User, error) {
	//1、密码脱敏处理
	u.Password = util.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}