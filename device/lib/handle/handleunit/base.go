package handleunit

import (
	"fmt"

	"../../data"
)

type Hand interface {
	HandleBusiness()
	HandleSend()
	SaveToSendList()
	SaveToDatabase()
}

type HandUnit struct {
	data.Data
	//是否存储到缓存
	storedCache bool
	//是否记录到数据库
	storedDatabase bool
}

func (this *HandUnit) SaveToSendList() {

}

func (this *HandUnit) SaveToDatabase() {
	if !this.storedDatabase {
		return
	}
	fmt.Println("HH")
}
