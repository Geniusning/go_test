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
