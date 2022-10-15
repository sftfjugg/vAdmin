package config

import "github.com/spf13/viper"

type Database struct {
	Dbtype   string
	Host     string
	Port     int
	Name string
	Username string
	Password string
	MaxIdleConns int
	MaxOpenConns int
}

func InitMysql(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		Dbtype:   cfg.GetString("dbType"),
		Host:     cfg.GetString("host"),
		Name: cfg.GetString("name"),
		Username: cfg.GetString("username"),
		Password: cfg.GetString("password"),
		MaxIdleConns: cfg.GetInt("MaxIdleConns"),
		MaxOpenConns: cfg.GetInt("MaxOpenConns"),
	}
}

func InitDrill(cfg *viper.Viper) *Database {
	return &Database{
		Port:     cfg.GetInt("port"),
		Dbtype:   cfg.GetString("dbType"),
		Host:     cfg.GetString("host"),
	}
}

var MysqlConfig = new(Database)

var DrillConfig = new(Database)

