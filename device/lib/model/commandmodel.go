package model

import (
	"database/sql"
	//"encoding/json"
	"fmt"

	"../data"
)

type Command struct {
	Model
}

func CommandModelInit(db *sql.DB) (command Command) {
	command.db = new(sql.DB)
	command.db = db
	command.table = "tz_std_command"
	return
}

func (this *Command) SaveCommand(cmd data.Data) (id int64) {
	_, err := this.db.Prepare("INSERT INTO " + this.table + "(stdid, commandId, sendcontent) VALUES (?, ?, ?)")
	//	sendcontent, _ := json.Marshal(cmd.Body)
	//	res, _ := db.Exec(cmd.Device, cmd.Sign, string(sendcontent))
	//	id, _ = res.LastInsertId()
	fmt.Println(err)
	return
}
