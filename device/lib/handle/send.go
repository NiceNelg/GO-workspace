package handle

import (
	"encoding/hex"
	"encoding/json"
	"net"
	"time"

	"../data"
	"github.com/garyburd/redigo/redis"
)

/**
 * @Function 发送协程
 * @Auther Nelg
 */
func (this *Handle) Send(Device *string, conn *net.TCPConn) {
	for {
		//客户端还没上传设备号
		if *Device == "" {
			time.Sleep(time.Second)
			continue
		}
		//从redis中取出连接
		redisCli := this.redisPool.Get()
		failrCmd := make([]string, 1)
		for redisData, err := redis.String(redisCli.Do("rpop", this.sendList+"_"+*Device)); err == nil; {
			//解析json数据
			var cmd data.Data
			err := json.Unmarshal([]byte(redisData), &cmd)
			if cmd.Content == "" {
				continue
			}
			//TODO 可按需重写标志位
			data, err := hex.DecodeString(cmd.Content)
			//添加标识头尾
			data = append(data, 0x7e)
			data = append([]byte{0x7e}, data...)
			//发送数据，检测发送时间是否已到
			if time.Now().Unix()-cmd.Sendtime >= this.sendtime {
				_, err = conn.Write(data)
				if err != nil {
					//发送失败，检测发送次数是否已超出
					cmd.Sendnum++
					if cmd.Sendnum < this.sendnum {
						cmd.Sendtime = time.Now().Unix()
						addCmd, _ := json.Marshal(cmd)
						failrCmd = append(failrCmd, string(addCmd))
					}
					for _, cmd := range failrCmd {
						redisCli.Do("lpush", this.sendList+"_"+*Device, cmd)
					}
					return
				}
			} else if cmd.Sendnum < this.sendnum {
				failrCmd = append(failrCmd, redisData)
			}
		}
		for _, cmd := range failrCmd {
			redisCli.Do("lpush", this.sendList+"_"+*Device, cmd)
		}
		//归还redis连接到redis连接池
		redisCli.Close()
		time.Sleep(1 * time.Second)
	}
	return
}
