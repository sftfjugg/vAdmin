package database

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"vAdmin/tools/config"
	"log"
	"strconv"
)

var Eloquent *gorm.DB

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	MaxIdleConns int
	MaxOpenConns int

)

func SetupMysql() {

	DbType = config.MysqlConfig.Dbtype
	Host = config.MysqlConfig.Host
	Port = config.MysqlConfig.Port
	Name = config.MysqlConfig.Name
	Username = config.MysqlConfig.Username
	Password = config.MysqlConfig.Password
	MaxIdleConns = config.MysqlConfig.MaxIdleConns
	MaxOpenConns = config.MysqlConfig.MaxOpenConns

	if DbType != "mysql" {
		log.Println("db type unknow")
	}
	var err error

	conn := GetMysqlConnect()

	log.Println(conn)

	var db Database
	if DbType == "mysql" {
		db = new(Mysql)
		Eloquent, err = db.Open(DbType, conn)

	} else {
		panic("db type unknow")
	}
	if err != nil {
		log.Fatalf("%s connect error %v", DbType, err)
	} else {
		log.Printf("%s connect success!", DbType)
	}


	if Eloquent.Error != nil {
		log.Fatalf("database error %v", Eloquent.Error)
	}

	Eloquent.LogMode(true)
}

func GetMysqlConnect() string {
	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=1000ms")
	return conn.String()
}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

type SqlLite struct {
}

func (*SqlLite) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}
