package internal

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// InitModules 模块初始化，输入配置文件目录，并给出目录中需要作为内部使用的配置文件
// FIXME: 目前全部配置写在app.yml一个中
func InitModules(configDir string, modules []string) error {
	// 命令行的config优先级最高
	conf := flag.String("config", configDir, "input config file like ./conf/dev/")
	flag.Parse()
	if *conf == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Println("------------------------------------------------------------------------")
	log.Printf("[Info] 使用的配置目录是：%s\n", *conf)
	log.Println("[Info] 加载中....")

	// 解析配置文件目录,也是验证
	if err := ParseConfPath(*conf); err != nil {
		return err
	}
	// 初始化配置文件
	if err := InitViperConf(); err != nil {
		return err
	}
	viperConfTopKeys := VConfMap.GetConfTopKeys()
	// 加载base信息
	if ContainsStrings("base", viperConfTopKeys) {
		if err := InitBaseConf(); err != nil {
			fmt.Printf("[ERROR] InitBaseConf:%s\n", err.Error())
		}
	}

	// 加载http信息
	return nil
}

func ContainsStrings(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// Destroy 公共销毁函数
func Destroy() {
	log.Println("------------------------------------------------------------------------")
	log.Printf("[INFO] %s\n", " start destroy resources.")
	// fixme: 系统内用到的链接等需要关闭的操作，在这里进行统一关闭
	// db.close()
	// pool.close()
	log.Printf("[INFO] %s\n", " success destroy resources.")
}
