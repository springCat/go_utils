package db

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

var db *xorm.Engine
var f *os.File

func Open() *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", dbConf.DbUrl)
	if err != nil {
		log.Println(err)
	}
	//sql log
	engine.ShowSQL = dbConf.ShowSql
	engine.ShowErr = true
	engine.ShowWarn = true

	engine.SetMaxIdleConns(dbConf.MaxIdleConns)
	engine.SetMaxOpenConns(dbConf.MaxOpenConns)

	f, err = os.Create("./sql.log")
	if err != nil {
		log.Println("err:",err)
	}
	engine.Logger = xorm.NewSimpleLogger(f)
	db = engine
	return engine
}

func Close(){
	err := db.Close()
	if err != nil {
		log.Println("err:",err)
	}
	err = f.Close()
	if err != nil {
		log.Println("err:",err)
	}
}

var dbConf DbConf

type DbConf struct {
	DbUrl        string `json:"db_url"`
	ShowSql      bool   `json:"db_show_sql"`
	MaxIdleConns int    `json:"db_maxIdleConns"`
	MaxOpenConns int    `json:"db_maxOpenConns"`
}

func Conf(dbConfFile string) {
	f, err := os.Open(dbConfFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(f).Decode(&dbConf)
	if err != nil {
		log.Fatal(err)
	}
}

type Executor func(*xorm.Session) error

func (self Executor) Run(session *xorm.Session) error{
	return self(session)
}

func Tx(f interface{}) bool {

	session := db.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		log.Println(err)
		return false
	}

	err = Executor((f.(func(*xorm.Session) error))).Run(session)
	if err != nil {
		log.Println(err)
		session.Rollback()
		return  false
	}

	err = session.Commit()
	if err != nil {
		log.Println(err)
		return  false
	}

	return true
}