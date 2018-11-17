package dbops

import (
	"database/sql"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddUserCredential(userName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO user(username, pwd) VALUES(?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(userName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(userName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM user WHERE username=?")
	if err != nil {
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(userName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUserCredential(userName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM user WHERE username=? AND pwd=?")
	if err != nil {
		return err
	}
	stmtDel.Exec(userName, pwd)
	stmtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	stmtIns, err := dbConn.Prepare("INSERT INTO user(id, author_id, name, display_time, createtime) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	t := time.Now()
	_, err = stmtIns.Exec(vid, aid, name, t, t)
	stmtIns.Close()
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{
		Id:          vid,
		AuthorId:    aid,
		Name:        name,
		DisplayTime: t.Format("2006-01-02 15:04:05"),
	}

	return res, err
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_time FROM video_info WHERE id=?")

	var aid int
	var dct int64
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()
	t := time.Unix(dct, 0).Format("2006-01-02 15:04:05")
	res := &defs.VideoInfo{
		Id:          vid,
		AuthorId:    aid,
		Name:        name,
		DisplayTime: t,
	}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
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

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.username, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

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

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
