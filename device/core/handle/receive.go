package handle

import (
	"encoding/json"

	"core/protocol808"
)

/**
 * @Function 保存任务
 * @Auther Nelg
 */
func (this *Handle) SaveTask(deviceId *string, dataArray [][]byte) {
	if len(dataArray) <= 0 {
		return
	}
	redisCli := this.redisPool.Get()
	defer redisCli.Close()
	//解析数据结构
	for _, cmd := range dataArray {
		//TODO 解析数据结构（不同协议需要更换包）
		data, err := protocol808.Resolvepack(cmd)
		if err != nil {
			continue
		}
		if *deviceId == "" {
			*deviceId = data.Device
		}
		//存入redis
		jsonData, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			continue
		}
		redisCli.Do("lpush", this.handleList, string(jsonData))
	}
}
