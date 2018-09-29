package data

/**
 * 数据头部
 */
type DataHeader struct {
	//命令标识
	sign string
	//消息属性
	attribute string
	//设备编号
	device string
	//命令流水
	sn string
}

/**
 * 数据包
 */
type Data struct {
	//数据头部
	DataHeader
	//数据体（已解析）
	body map[string][string]
	//未解包数据
	content interface{}
}

