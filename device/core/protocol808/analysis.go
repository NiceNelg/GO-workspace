package protocol808

import (
	"encoding/hex"
	"errors"

	"core/data"
)

/**
 * @Function 数据截取（不同协议需要重写，目前使用的是808协议）
 * @Auther Nelg
 */
func Cutpack(content []byte, buffer []byte) (dataArray [][]byte, inComplete []byte) {
	//判断buffer是否为空
	if len(buffer) <= 0 {
		isContinue := false
		cLen := len(content) - 1
		for index, value := range content {
			if value == 0x7e {
				isContinue = true
				if index+1 <= cLen && content[index+1] != 0x7e {
					content = content[index:]
				} else if value == 0x7e && index+1 <= cLen && content[index+1] == 0x7e {
					content = content[index+1:]
				}
				break
			}
		}
		if !isContinue {
			return
		}
	}
	//与上次未完整接收的buffer数据合并
	content = append(buffer, content...)
	maxIndex := len(content) - 1
	//数据切割
	tempArray := make([][]byte, 1)
	num := 0
	for index, value := range content {
		if value == 0x7e && index+1 <= maxIndex && content[index+1] != 0x7e {
			tempArray = append(tempArray, make([]byte, 0))
			tempArray[num] = append(tempArray[num], value)
		} else if value == 0x7e && index+1 <= maxIndex && content[index+1] == 0x7e {
			tempArray[num] = append(tempArray[num], value)
			num++
		} else {
			tempArray[num] = append(tempArray[num], value)
		}
	}
	//去除空命令数据
	notemptyArray := make([][]byte, 0)
	for _, cmd := range tempArray {
		if len(cmd) > 0 {
			notemptyArray = append(notemptyArray, cmd)
		}
	}
	//命令筛选
	discard := make([]int, 0)
	arrMaxIndex := len(notemptyArray) - 1
	for index, cmd := range notemptyArray {
		signNum := 0
		for _, value := range cmd {
			if value == 0x7e {
				signNum++
			}
		}
		//标识位少于2
		if signNum < 2 {
			//已到最后一个命令，代表命令不全，加入buffer
			if index == arrMaxIndex {
				inComplete = cmd
			}
			//需要丢弃的数据
			discard = append(discard, index)
		} else if signNum > 2 {
			//多个不完整的包与一个完整的包沾合在一起，取出完整的包数据
			//TODO 待完善
			times := 0
			for i, v := range cmd {
				if v == 0x7e {
					times++
				}
				if times > 1 {
					notemptyArray[index] = cmd[i:]
					break
				}
			}
		}
	}
	//删除无用命令数据
	for index, value := range notemptyArray {
		add := true
		for _, v := range discard {
			if v == index {
				add = false
			}
		}
		if add {
			dataArray = append(dataArray, value)
		}
	}
	return
}

/**
 * @Function 解析数据包结构体（不同协议需要重写，目前使用的是808协议）
 * @Auther Nelg
 */
func Resolvepack(cmd []byte) (data data.Data, err error) {
	if len(cmd) <= 0 {
		err = errors.New("data don't exist")
		return
	}
	//数据转义
	cmd = ReverseEscape(cmd)
	if len(cmd) <= 0 {
		err = errors.New("data escape fail")
		return
	}
	//异或校验
	xor := BuildBCC(cmd[1 : len(cmd)-2])
	if cmd[len(cmd)-2] != xor {
		err = errors.New("data BCC fail")
		return
	}
	//数据结构解析
	data.Sign = hex.EncodeToString(cmd[1:3])
	data.Attribute = hex.EncodeToString(cmd[3:5])
	data.Device = hex.EncodeToString(cmd[5:11])
	data.Sn = hex.EncodeToString(cmd[11:13])
	data.Content = hex.EncodeToString(cmd[13 : len(cmd)-2])
	return
}

/**
 * @Function 数据转义
 * @Auther Nelg
 */
func Escape(cmd []byte) (data []byte) {
	temp := make([]byte, 0)
	//转义
	cmdMaxIndex := len(cmd) - 1
	for index, value := range cmd {
		if index <= 0 || index >= cmdMaxIndex {
			temp = append(temp, value)
			continue
		}
		if value == 0x7e {
			temp = append(temp, 0x7d, 0x02)
		} else if value == 0x7d {
			temp = append(temp, 0x7d, 0x01)
		} else {
			temp = append(temp, value)
		}
	}
	data = temp
	return

}

/**
 * @Function 数据反转义
 * @Auther Nelg
 */
func ReverseEscape(cmd []byte) (data []byte) {
	if len(cmd) <= 0 {
		return
	}
	temp := make([]byte, 0)
	//反转义
	cmdMaxIndex := len(cmd) - 1
	for index := 0; index <= cmdMaxIndex; {
		if index <= 0 || index >= cmdMaxIndex {
			temp = append(temp, cmd[index])
			index++
			continue
		}
		if cmd[index] == 0x7d && index+1 < cmdMaxIndex && cmd[index+1] == 0x02 {
			temp = append(temp, 0x7e)
			index += 2
		} else if cmd[index] == 0x7d && index+1 < cmdMaxIndex && cmd[index+1] == 0x01 {
			temp = append(temp, 0x7d)
			index += 2
		} else {
			temp = append(temp, cmd[index])
			index++
		}
	}
	data = temp
	return
}

/**
 * @Function 计算校验和
 * @Auther Nelg
 */
func BuildBCC(content []byte) (xor byte) {
	xor = 0x00
	for _, value := range content {
		xor ^= value
	}
	return
}
