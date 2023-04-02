package core

import (
	"gBlog/config"
	"gBlog/global"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func InitConfig() {
	const ConfigFile = "conf.yaml"
	configFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Println("读取配置文件失败")
		return
	}
	conf := &config.Config{}
	err = yaml.Unmarshal(configFile, conf)
	if err != nil {
		log.Println("解析配置文件失败")
		return
	}
	global.Conf = conf
}
