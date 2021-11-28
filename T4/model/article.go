package model

import (
	"T4/util"
	"log"
)

// Article 文章对象，定义了文章的基础信息
type Article struct {
	Id      int    `json:"id"`      //文章id
	Author  string `json:"author"`  //文章作者
	Title   string `json:"title"`   //文章标题
	Context string `json:"context"` //文章内容
	Likes   int    `json:"likes"`   //文章点赞数
}

// @title 发布文章
// @description 接收文章对象在数据库进行存储
// @param msg *Article "要存储的文章对象"
// @return code int "状态码"
func PostArticle(msg *Article) int {
	str := "insert into article (author,title,context) values (?,?,?)"
	stmt, err := DB.Prepare(str)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	_, err = stmt.Exec(msg.Author, msg.Title, msg.Context)
	if err != nil {
		log.Println(err)
		return emsg.PostArticleFailed
	}
	return emsg.Success
}

// @title 获取文章对象
// @description 接收文章id在数据库中查找返回文章内容
// @param id int "要查询的文章id"
// @return msg *Article "文章id对应的文章内容" code int "状态码"
func GetArticle(id int) (*Article, int) {
	str := "select * from article where id = ?"
	stmt, err := DB.Prepare(str)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	row := stmt.QueryRow(id)
	var msg Article
	err = row.Scan(&msg.Id, &msg.Author, &msg.Title, &msg.Context, &msg.Likes)
	if err != nil {
		log.Println(err)
		return nil, emsg.GetArticleFailed
	}
	return &msg, emsg.Success
}

// @title 删除文章
// @description 接收文章id将其删除
// @param id int "要删除的文章id"
// @return code int "状态码"
func DeleteArticle(id int) int {
	str := "delete from article where id = ?"
	stmt, err := DB.Prepare(str)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return emsg.DeleteArticleFailed
	}
	return emsg.Success
}

// @title 点赞文章
// @description 接收文章id将其点赞数加一
// @param username string "记录的点赞的用户"(未实现), id int "要点赞的文章id"
// @return code int "状态码"
func LikeArticle(username string, id int) int {
	str := "update article set likes = likes + 1 where id = ?;"
	stmt, err := DB.Prepare(str)
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return emsg.LikeArticleFailed
	}
	return emsg.Success
}
