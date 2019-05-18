package dbops

import (
	"GOFFPmgSystem/api/defs"
	"GOFFPmgSystem/api/utils"
	"database/sql"
	"fmt"
	"log"
)


func AddUserCredential(loginName string,pwd string) error {
	stmt, e := dbConn.Prepare("INSERT INTO `users` (`login_name`,`pwd`) VALUE (?,?)")
	defer func() {
		stmt.Close()
	}()
	if e != nil {
		return e
	}
	pwd = utils.Md5String(pwd)
	_, e = stmt.Exec(loginName, pwd)
	return e
}

func GetUserCredential(loginName string) (string,error)  {
	stmt, e := dbConn.Prepare("SELECT `pwd` FROM `users` WHERE `login_name` = ?")
	defer func() {
		stmt.Close()
	}()
	if e != nil {
		log.Printf("%s",e)
		return "",e
	}
	var pwd string
	e = stmt.QueryRow(loginName).Scan(&pwd)

	// 对空返回 容错
	if e!=nil || e!=sql.ErrNoRows {
		return pwd,nil
	}
	return pwd,e
}

func DeleteUser(loginName string,pwd string) error {
	stmt, e := dbConn.Prepare("DELETE FROM `users` WHERE `login_name` = ? AND `pwd` = ?")
	defer func() {
		stmt.Close()
	}()
	if e != nil{
		log.Printf("%s",e)
		return e
	}
	pwd = utils.Md5String(pwd)
	_, e = stmt.Exec(loginName, pwd)
	return e
}

func AddNewVideo(aid int,name string) (*defs.VideoInfo,error) {
	s, e := utils.NewUUIDSimplicity()
	if e != nil {
		return nil,e
	}
	time := utils.GetCurrentTime()
	stmt, e := dbConn.Prepare(`INSERT INTO video_info (video_id, author_id, name, create_time) VALUES(?, ?, ?, ?)`)

	fmt.Println(time)
	fmt.Println(s)
	defer func() {stmt.Close()}()
	if e!=nil {
		fmt.Println("dbConn.Prepare")
		return nil,e
	}
	_, e = stmt.Exec(s, aid, name, time)
	if e!=nil {
		return nil,e
	}
	info := &defs.VideoInfo{VideoId: s, AuthorId: aid, Name: name, CreateTime: time}

	return info,nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo,error)  {
	stmt, e := dbConn.Prepare("SELECT `author_id`,`name`,`create_time` FROM `video_info` WHERE `video_id` = ?")
	defer func() {
		stmt.Close()
	}()
	if e != nil {
		return nil,e
	}
	var aid int
	var ctime string
	var name string

	e = stmt.QueryRow(vid).Scan(&aid, &name, &ctime)
	if e != nil && err != sql.ErrNoRows {
		return nil,err
	}
	if err == sql.ErrNoRows {
		return nil,nil
	}

	info := &defs.VideoInfo{VideoId: vid, AuthorId: aid, Name: name, CreateTime: ctime}
	return info,nil
}

func DeleteVideoInfo(vid string,aid int) error {
	stmt, e := dbConn.Prepare("DELETE FROM `video_info` WHERE `video_id` = ? AND `author_id` = ?")
	defer func() {
		stmt.Close()
	}()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(vid, aid)
	return e
}

func AddNewComment(vid string,aid int,content string) error {
	stmt, e := dbConn.Prepare("INSERT INTO `comments`(`video_id`,`author_id`,`content`,`create_time`) VALUE(?,?,?,?)")
	if e != nil {
		return e
	}
	time := utils.GetCurrentTime()
	_, e = stmt.Exec(vid, aid, content, time)
	return e
}

func ListComments(vid string,from,to int) ([]*defs.Comment,error) {
	//stmt, e := dbConn.Prepare(`SELECT comments.id,user.Login_name,comments.content FROM comments
	//	INNER JOIN users ON comments.author_id = users.id
	//	WHERE comment.video_id = ? AND comment.create_time > ? AND comment.create_time <= ?
	//	ORDER BY comment.create_time DESC
	//	`)
	stmt, e := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? 
		ORDER BY comments.create_time DESC 
		`)
	if e != nil {
		fmt.Println(e.Error())
		return nil,e
	}
	//多条查询定义存储容器
	var res []*defs.Comment

	//rows, e := stmt.Query(vid, from, to)
	rows, e := stmt.Query(vid)
	if e != nil && e != sql.ErrNoRows {
		fmt.Println(e.Error())
		return nil,e
	}

	//fmt.Printf("%v,%T",rows,rows)

	for rows.Next() {
		var id,name,content string
		if e := rows.Scan(&id, &name, &content);e != nil {
			return res,e
		}
		comment := &defs.Comment{Id: id, VideoId: vid, AuthorName: name, Content: content}
		res = append(res,comment)
	}
	return res,nil
}
