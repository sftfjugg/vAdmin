package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var cfgMysql *viper.Viper
var cfgDrill *viper.Viper
var cfgApplication *viper.Viper
var cfgJwt *viper.Viper
var cfgLog *viper.Viper


//载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	cfgMysql = viper.Sub("settings.mysql")
	if cfgMysql == nil {
		panic("config not found settings.mysql")
	}
	MysqlConfig = InitMysql(cfgMysql)

	cfgDrill = viper.Sub("settings.drill")
	if cfgDrill == nil {
		panic("config not found settings.drill")
	}
	DrillConfig = InitDrill(cfgDrill)

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgJwt = viper.Sub("settings.jwt")
	if cfgJwt == nil {
		panic("config not found settings.jwt")
	}
	JwtConfig = InitJwt(cfgJwt)

	cfgLog = viper.Sub("settings.log")
	if cfgLog == nil {
		panic("config not found settings.log")
	}
	LogConfig = InitLog(cfgLog)
}


func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	viper.WriteConfig()
}
