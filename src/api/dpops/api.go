package dpops

import (
	"api/defs"
	"api/utils"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// AddUserCredential 新增用户更
func AddUserCredential(loginName string, pwd string) error {
	stmtins, err := dbconn.Prepare("INSERT INTO users (login_name,pwd) VALUES(?,?)")
	if err != nil {
		return err
	}

	_, err = stmtins.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtins.Close()
	return nil
}

// GetUserCredential 获取用户
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbconn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}

//DeleteUser 删除用户
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbconn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("delete User %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

// AddNewVideo 新增视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//create uid

	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05")
	stmtIns, err := dbconn.Prepare("INSERT INTO video_info (id,author_id,name,display_time) VALUE(?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)

	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		ID:           vid,
		AuthorID:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}

	defer stmtIns.Close()

	return res, nil

}

//GetVideoInfo 获得视频信息
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbconn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{ID: vid, AuthorID: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

//DeleteVideoInfo 删除视频
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbconn.Prepare("SELECT FROM video_info WHERE id=?")
	if err != nil {
		log.Printf("DeleteVideoInfo err=%v", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()

	return nil
}

//AddNewComment 新增评论
func AddNewComment(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbconn.Prepare("INSERT INTO comment (id,video_id,author_id,content) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()

	return nil
}

//ListComments list comments
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbconn.Prepare(`SELECT comment.id,users.login_name,comment.content FROM comment INNER JOIN users
	ON comment.author_id=users.id WHERE comment.video_id=?AND comment.time>FROM_UNIXTIME(?)AND comment.time<= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{ID: id, VideoID: vid, Author: name, Content: content}
		res = append(res, c)
	}

	defer stmtOut.Close()

	return res, nil
}
