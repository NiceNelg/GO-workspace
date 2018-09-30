package data

/**
 * 数据头部
 */
type DataHeader struct {
	//命令标识
	Sign string
	//消息属性
	Attribute string
	//设备编号
	Device string
	//命令流水
	Sn string
}

/**
 * 数据包
 */
type Data struct {
	//数据头部
	DataHeader
	//数据体（已解析）
	Body map[string][]string
	//未解包数据（不含协议头及协议尾，对于16进制数据已转换成可视化数据）
	Content string
}
