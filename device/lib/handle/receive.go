package handle

import (
	"fmt"
	"os"

	//"../data"
	"../protocol808"
)

/**
 * @Function 保存任务
 * @Auther Nelg
 */
func SaveTask(dataArray [][]byte) {
	if len(dataArray) <= 0 {
		return
	}
	//解析数据结构
	for _, cmd := range dataArray {
		data := protocol808.Resolvepack(cmd)
		fmt.Println(cmd)
		fmt.Println(data)
	}
	os.Exit(0)
}
