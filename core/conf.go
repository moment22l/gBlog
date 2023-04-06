package core

import (
	"gBlog/config"
	"gBlog/global"
	"io/fs"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

const ConfigFile = "conf.yaml"

// InitConfig 初始化配置
func InitConfig() {
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

// 修改配置文件
func ModifyConf() error {
	byteData, err := yaml.Marshal(global.Conf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("修改配置文件成功")
	return nil
}
