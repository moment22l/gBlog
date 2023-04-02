package core

import (
	"gBlog/config"
	"gBlog/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func InitConfig() {
	const ConfigFile = "conf.yaml"
	configFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		global.Log.Error("open config file failed")
		return
	}
	conf := &config.Config{}
	err = yaml.Unmarshal(configFile, conf)
	if err != nil {
		global.Log.Error("parse config file failed")
		return
	}
	global.Conf = conf
}
