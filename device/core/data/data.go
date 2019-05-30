package data

/**
 * @Struct 数据头部
 * @Auther Nelg
 * @Date 2019.05.30
 */
type DataHeader struct {
	/*命令标识*/
	Sign string `json:"sign"`
	/*消息属性*/
	Attribute string `json:"attribute"`
	/*设备编号*/
	Device string `json:"device"`
	/*命令流水*/
	Sn string `json:"sn"`
}

/**
 * @Struct 数据包
 * @Auther Nelg
 * @Date 2019.05.30
 */
type Data struct {
	/*数据头部*/
	DataHeader
	/*数据体（已解析）*/
	Body map[string]string `json:"-"`
	/*未解包数据（不含协议头及协议尾，对于16进制数据已转换成可视化数据）*/
	Content string `json:"content"`
	/*发送时间*/
	Sendtime int64 `json:"sendtime"`
	/*发送次数*/
	Sendnum int `json:"sendnum"`
}
