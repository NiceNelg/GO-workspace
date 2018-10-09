package handleunit

import (
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
	cmd.Body = make(map[string]string, 1)
	cmd.Body["authkey"] = string(content[:8])
	cmd.Body["std_type"] = string(content[8:16])
	cmd.Body["car_type"] = string(content[16:18])
	cmd.Body["CCID"] = string(content[18:58])
	cmd.Body["now_version_len"] = string(content[58:60])
	//	versionLen, _ := strconv.Atoi(string(cmd.Body["now_version_len"]))
	//	versionLen = versionLen*2 + 60
	//	cmd.Body["now_version"] = string(content[60:versionLen])
	cmd.Body["now_version"] = string(content[60:])
	authcheck = new(Authcheck)
	authcheck.Data = cmd
	authcheck.storedCache = true
	authcheck.storedDatabase = true
	return
}

/**
 * @Function 业务处理
 * @Auther Nelg
 */
func (this *Authcheck) HandleBusiness() {

}

/**
 * @Function 数据下发处理
 * @Auther Nelg
 */
func (this *Authcheck) HandleSend() {

}
