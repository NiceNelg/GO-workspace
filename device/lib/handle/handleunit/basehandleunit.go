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
	//保存指令
	id := commandModel.SaveCommand(this.Data)
	sn := strconv.FormatInt(id, 16)
	l := 4 - len(sn)
	if l > 0 {
		var zero string
		for i := 0; i < l; i++ {
			zero += "0"
		}
		sn = zero + sn
	}
	this.Sn = sn
	//更新指令流水号
	commandModel.SaveSn(id, this.Sn)
	return
}
