package toml

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

// DecodeFile 解释不同环境配置结构
// 参数:
//   - filepath 配置文件路径
//   - config 解释目标结构体
func DecodeFile(filepath string, config interface{}) error {

	// 读取配置文件, 解决跑测试的时候找不到配置文件的问题，最多往上找10层目录
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath); err == nil {
			break
		} else {
			filepath = "../" + filepath
		}
	}

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return errors.New("config file is not exist")
	}

	_, err = toml.DecodeFile(filepath, config)
	return err
}
