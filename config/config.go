package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func init() {
	// 读取
	f, err := os.Open(configPath)
	defer f.Close()
	if err != nil {
		log.Errorf("Init config Error %v", err)
		return
	}
	// 解构
	err = yaml.NewDecoder(f).Decode(&Conf)
	if err != nil {
		log.Errorf("Decode Conf Error %v", err)
		panic(err)
	}
}

type Storage struct {
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Host   string `yaml:"host"`
}

/*
type Tag struct {
	Required string `yaml:"required"`
	Min      string `yaml:"min"`
	Max      string `yaml:"max"`
}

type Field struct {
	Name            string `yaml:"name"`
	Nickname        string `yaml:"nickname"`
	UserName        string `yaml:"username"`
	PassWord        string `yaml:"password"`
	PasswordConfirm string `yaml:"passwordconfirm"`
}
*/
var (
	Conf       = Config{}
	configPath = "./gxu_pointsmall/config/config.yaml"
)

type Config struct {
	Storages map[string]Storage `yaml:"storages"`
	//Tag      map[string]Tag     `yaml:"tag"`
	//Field    map[string]Tag     `yaml:"field"`
}
