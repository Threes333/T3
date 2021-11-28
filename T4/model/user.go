package model

import (
	"T4/util"
	"log"
)

// User 用户对象，定义了用户的基础信息
type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

// @title 用户注册
// @description 接收用户名和密码进行有效验证并在数据库进行存储
// @param username,password string "要进行注册的用户名与密码"
// @return code int "状态码"
func Register(username, password string) int {
	var id int
	str := "select id from user where username = ?"
	stmt, err := DB.Prepare(str)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	row := stmt.QueryRow(username)
	_ = row.Scan(&id)
	if id == 0 {
		//该用户名可用y6
		str = "insert into user(username,password) values (?,?)"
		stmt, err = DB.Prepare(str)
		if err != nil {
			log.Println(err)
			return emsg.Error
		}
		_, err = stmt.Exec(username, password)
		if err != nil {
			log.Println(err)
			return emsg.Error
		}
		return emsg.Success
	}
	return emsg.UsernameExist
}

// @title 用户登录
// @description 接收用户名和密码进行有效验证
// @param username,password string "要进行登录的用户名与密码"
// @return code int "状态码"
func Login(username, password string) int {
	var pw string
	str := "select password from user where username = ?"
	stmt, err := DB.Prepare(str)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	err = stmt.QueryRow(username).Scan(&pw)
	if pw == "" {
		return emsg.UsernameNoExist
	} else if pw != password {
		return emsg.PasswordWrong
	}
	return emsg.Success
}
