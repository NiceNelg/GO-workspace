package handleunit

import (
	"time"

	"../../data"
)

type Authcheck struct {
	HandUnit
}

/**
 * @Function 数据内容解析
 * @Auther Nelg
 */
func AuthcheckInit(cmd data.Data) (authcheck *Authcheck) {
	content := []rune(cmd.Content)
	cmd.Body = make(map[string]string, 6)
	cmd.Body["authkey"] = string(content[:8])
	cmd.Body["std_type"] = string(content[8:16])
	cmd.Body["car_type"] = string(content[16:18])
	cmd.Body["CCID"] = string(content[18:58])
	cmd.Body["now_version_len"] = string(content[58:60])
	cmd.Body["now_version"] = string(content[60:])
	authcheck = new(Authcheck)
	authcheck.Data = cmd
	return
}

/**
 * @Function 业务处理	需要数据下发的时候请返回Hand结构体
 * @Auther Nelg
 */
func (this *Authcheck) HandleBusiness() (sendCmd Hand) {
	authCheck := &Authcheck{}
	authCheck.Sign = "8102"
	authCheck.Sn = this.Sn
	authCheck.Device = this.Device
	authCheck.Body = make(map[string]string, 4)
	authCheck.Body["ack_sn"] = this.Sn
	authCheck.Body["ack_sign"] = this.Sign
	authCheck.Body["result"] = "00"
	//获取当前时间
	authCheck.Body["time"] = time.Now().Format("060102150405")
	authCheck.StoredDatabase = false
	authCheck.StoredCache = true
	sendCmd = authCheck
	return
}

/**
 * @Function 数据下发处理
 * @Auther Nelg
 */
//func (this *Authcheck) HandleSend() {
//}
