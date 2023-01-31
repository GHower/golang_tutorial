package internal

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

type BaseConf struct {
	DebugMode   string `mapstructure:"debug_mode"`
	Timezone    string `mapstructure:"timezone"`
	ApiSlowTime int    `mapstructure:"api_slow_time"`
}
type DbConf struct{}

type GlobalConf struct {
	BaseConf *BaseConf `mapstructure:"base"`
	//DbConf   DbConf   `yaml:"db"`
}

// ViperConfMap viper 反映射文件所在
type ViperConfMap map[string]string

/* 全局变量 */
var ConfBase *BaseConf
var VConfMap ViperConfMap

func init() {
	VConfMap = make(ViperConfMap)
}

// InitViperConf viper初始化配置
func InitViperConf() error {
	f, err := os.Open(ConfEnvPath + "/")
	if err != nil {
		return err
	}
	// 通常配置目录中的目录项不会太多,1024个基本够用
	fileList, err := f.ReadDir(1024)
	if err != nil {
		return err
	}
	for _, file := range fileList {
		if !file.IsDir() {
			fullPath := ConfEnvPath + "/" + file.Name()
			bts, err := ioutil.ReadFile(fullPath)
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType("yaml")
			_ = v.ReadConfig(bytes.NewBuffer(bts))
			keys := v.AllKeys()
			fmt.Println(keys)
			//pathArr := strings.Split(file.Name(), ".")
			//if ViperConfMap == nil {
			//	ViperConfMap = make(map[string]*viper.Viper)
			//}
			for _, key := range keys {
				ks := strings.Split(key, ".")
				VConfMap[ks[0]] = fullPath
			}
		}
	}
	//fmt.Println(VConfMap)
	return nil

}
func InitBaseConf() error {
	ConfGlobal := &GlobalConf{}
	err := ParseConfig(VConfMap["base"], ConfGlobal)
	ConfBase = ConfGlobal.BaseConf

	fmt.Println(ConfBase)
	return err
}
func (v *ViperConfMap) GetConfTopKeys() []string {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]string, len(VConfMap))
	for k := range VConfMap {
		keys[j] = k
		j++
	}
	return keys
}
