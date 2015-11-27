package prj

import (
	"flag"
	"fmt"
//	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/springCat/go_utils/dbx"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
//"github.com/elazarl/go-bindata-assetfs"
//"syscall"
)

func Start() (server *gin.Engine, db *xorm.Engine, conf *Conf) {

	conf = LoadDefaultConf()
	server = openGin(conf.Profile)
	db = OpenDb(conf.Db)


	err := os.Mkdir("logs", os.ModePerm)
	if err != nil {
		log.Println("err:",err)
	}
	return
}

/*********************gin  start*************************/
func openGin(profile string) *gin.Engine {

	if profile == Product {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
//	server.Use(static.ServeRoot("/", "./static"))
	ginLog, err := os.Create("logs/server.log")
	if err != nil {
		fmt.Println("err:", err)
	}
	server.Use(gin.LoggerWithWriter(ginLog))
	return server

}

/*********************gin  end*************************/

/*********************db  start*************************/
var db *xorm.Engine
var dbLog *os.File

func OpenDb(dbConf DbConf) *xorm.Engine {

	engine, err := xorm.NewEngine(DbUrl(dbConf))
	if err != nil {
		log.Println(err)
	}
	//sql log
	engine.ShowSQL = dbConf.ShowSql
	engine.ShowErr = true
	engine.ShowWarn = true

	engine.SetMaxIdleConns(dbConf.MaxIdleConns)
	engine.SetMaxOpenConns(dbConf.MaxOpenConns)

	dbLog, err = os.Create("logs/db.log")
	if err != nil {
		log.Println("err:", err)
	}
	engine.Logger = xorm.NewSimpleLogger(dbLog)
	//init tx
	dbx.InitTx(engine)

	db = engine
	return engine
}

func CloseDb() {
	err := db.Close()
	if err != nil {
		log.Println("err:", err)
	}
	err = dbLog.Close()
	if err != nil {
		log.Println("err:", err)
	}
}

/*********************db  end*************************/
/**
prj path
*/
func CurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

/**
LocalAddr
*/
func LocalAddr() string { //Get ip
	conn, err := net.Dial("udp", "baidu.com:80")
	defer conn.Close()

	if err != nil {
		return err.Error()
	}
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

/**
Addr in internet
*/
func PublicAddr() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	defer resp.Body.Close()
	if err != nil {
		log.Println("err:", err)
		return err.Error()
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err)
		return err.Error()
	}
	return string(result)
}

/**
server port
*/
var port = flag.String("p", "", "http serve port")

func Port() *string {
	flag.Parse()
	return port
}

/*********************profile  start*************************/
const (
	Local   string = "local"
	Dev     string = "dev"
	Product string = "product"
)

var conf = flag.String("conf", "", "http conf file")

func ConfPath() *string {
	flag.Parse()
	return conf
}

/*********************profile  end**************************/

/*********************db  start*************************/
func DbUrl(dbConf DbConf) (string, string) {
	switch dbConf.Type {
	case "postgres":
		return pgUrl(dbConf.User, dbConf.Pwd, dbConf.Address, dbConf.Database)
	case "mysql":
		return mysqlUrl(dbConf.User, dbConf.Pwd, dbConf.Address, dbConf.Database)

	default:
		panic("unsupport db type")
	}
}

func pgUrl(dbUser, dbPwd, dbAddr, dbName string) (string, string) {
	return "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPwd, dbAddr, dbName)
}

func mysqlUrl(dbUser, dbPwd, dbAddr, dbName string) (string, string) {
	return "mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8", dbUser, dbPwd, dbAddr, dbName)
}

/*********************db  end*************************/

/*********************version  start*************************/
/**
if >0 version1 > version2
if <0 version1 < version2
if =0 version1 = version2
 */
func VersionCompare(version1 string, version2 string) (int, error) {
	s1 := strings.Split(version1, ".")
	s2 := strings.Split(version2, ".")
	l1 := len(s1)
	l2 := len(s2)

	l := l1
	if l1 > l2 {
		l = l2
	}

	for i := 0; i < l; i++ {
		i1, err := strconv.Atoi(s1[i])
		if err != nil {
			return 0, err
		}

		i2, err := strconv.Atoi(s2[i])
		if err != nil {
			return 0, err
		}

		if result := i1 - i2; result != 0 {
			return result,nil
		}
	}
	return l1 - l2,nil
}

/*********************version  end*************************/
