package handleunit

import (
	"database/sql"
	"fmt"
	"strconv"

	"../../data"
	"../../model"
)

type Hand interface {
	HandleBusiness() Hand
	HandleSend()
	SaveToSendList()
	SaveToDatabase(db *sql.DB)
}

type HandUnit struct {
	data.Data
	//是否存储到缓存
	StoredCache bool
	//是否记录到数据库
	StoredDatabase bool
}

func (this *HandUnit) SaveToSendList() {

}

func (this *HandUnit) SaveToDatabase(db *sql.DB) {
	if !this.StoredDatabase {
		return
	}
	commandModel := model.CommandModelInit(db)
	this.Sn = strconv.FormatInt(commandModel.SaveCommand(this.Data), 10)
	fmt.Println(this.Sn)
}
