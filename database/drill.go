package database

import (
	"bytes"
	_ "github.com/pinpt/go-drill"//加载drill
	"database/sql"
	"vAdmin/tools/config"
	"log"
	"strconv"
)

var Eloq *sql.DB

type Databasex interface {
	Open(dbType string, conn string) (db *sql.DB, err error)
}

type Drill struct {
}

func SetupDrill() {

	DbType = config.DrillConfig.Dbtype
	Host = config.DrillConfig.Host
	Port = config.DrillConfig.Port

	if DbType != "drill" {
		log.Println("db type unknow")
	}
	var err error

	conn := GetDrillConnect()

	log.Println(conn)

	var db Databasex
	if DbType == "drill" {
		db = new(Drill)
		Eloq, err = db.Open(DbType, conn)

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

func GetDrillConnect() string {
	var conn bytes.Buffer
	conn.WriteString("http://")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	return conn.String()
}

func (*Drill) Open(dbType string, conn string)  (db *sql.DB, err error) {
	// sql.Open("drill", "http://192.168.0.18:30700")
	eloq, err := sql.Open(dbType, conn)
	eloq.SetMaxIdleConns(5)
	return eloq, err
}


