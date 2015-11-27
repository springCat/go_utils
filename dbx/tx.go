package dbx

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
)

var db *xorm.Engine

type Executor func(*xorm.Session) error

func (self Executor) Run(session *xorm.Session) error {
	return self(session)
}

func InitTx(engine *xorm.Engine)  {
	db = engine
}

func Tx(f interface{}) bool {
	return DbTx(db,f)
}


func DbTx(engine *xorm.Engine,f interface{}) bool {

	session := engine.NewSession()
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
		return false
	}

	err = session.Commit()
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}