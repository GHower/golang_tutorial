package internal

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ConfEnvPath string //配置文件夹
	ConfEnv     string //配置环境名
)

// ParseConfPath 解析配置目录，将读取路径和环境记录在全局变量
// 如：configDir=conf/dev/base.json 	ConfEnvPath=conf/dev	ConfEnv=dev
// 如：configDir=conf/base.json		ConfEnvPath=conf		ConfEnv=conf
// fixme: 监听全局环境变量改变，直接修改项目所用的配置，比如从注册中心获得配置时
func ParseConfPath(configDir string) error {
	path := strings.Split(configDir, "/")
	if len(path) < 1 {
		return errors.New("configDir error. eg: conf/dev/")
	}
	prefix := strings.Join(path[:len(path)-1], "/")
	ConfEnvPath = prefix
	ConfEnv = path[len(path)-2]
	return nil
}

func GetConfEnv() string {
	return ConfEnv
}

func GetConfPath(fileName string) string {
	return ConfEnvPath + "/" + fileName + ".yml"
}

func GetConfFilePath(fileName string) string {
	return ConfEnvPath + "/" + fileName
}
func ParseConfig(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open config %v fail, %v", path, err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read config fail, %v", err)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	_ = v.ReadConfig(bytes.NewBuffer(data))
	bb := &BaseConf{}
	fmt.Println(v.UnmarshalKey("base", bb))
	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("Parse config fail, config:%v, err:%v", string(data), err)
	}
	return nil
}
