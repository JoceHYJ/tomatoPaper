package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type config struct {
	Server       server       `yaml:"server"`
	Db           db           `yaml:"db"`
	Log          log          `yaml:"log"`
	FileSettings fileSettings `yaml:"fileSettings"`
}

// server 项目端口配置
type server struct {
	Address string `yaml:"address"`
}

// db 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Loc      string `yaml:"loc"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type fileSettings struct {
	Uploader   string `yaml:"uploader"`
	Downloader string `yaml:"downloader"`
	FileHost   string `yaml:"fileHost"`
}

// log 日志配置
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

// init 初始化配置
func init() {
	yamlFile, err := os.ReadFile("/home/jocehyj/goWorkspace/src/tomatoPaper/common/config/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
