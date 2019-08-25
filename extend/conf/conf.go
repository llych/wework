package conf

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/spf13/viper"
)

type weworkConf struct {
	CorpID  string `mapstructure:"CorpID"`
	Secret  string `mapstructure:"Secret"`
	AgentId string `mapstructure:"AgentId"`
}

var WeworkConf = &weworkConf{}

type logger struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

var LoggerConf = &logger{}

type cors struct {
	AllowAllOrigins  bool          `mapstructure:"allowAllOrigins"`
	AllowMethods     []string      `mapstructure:"allowMethods"`
	AllowHeaders     []string      `mapstructure:"allowHeaders"`
	ExposeHeaders    []string      `mapstructure:"exposeHeaders"`
	AllowCredentials bool          `mapstructure:"allowCredentials"`
	MaxAge           time.Duration `mapstructure:"maxAge"`
}

var CORSConf = &cors{}

type httpConf struct {
	Port   int    `mapstructure:"port"`
	Enable bool   `mapstructure:"enable"`
	Token  string `mapstructure:"token"`
}

var HttpConf = &httpConf{}

func Setup() {
	var (
		file []byte
		err  error
	)
	viper.SetConfigType("yaml")
	if file, err = ioutil.ReadFile("config/app.yaml"); err != nil {
		panic(err)
	}
	viper.ReadConfig(bytes.NewBuffer(file))
	viper.UnmarshalKey("wework", WeworkConf)
	viper.UnmarshalKey("logger", LoggerConf)
	viper.UnmarshalKey("cors", CORSConf)
	viper.UnmarshalKey("http", HttpConf)
}
