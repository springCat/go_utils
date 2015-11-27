package prj

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type DbConf struct {
	Type         string `json:"type"`
	User         string `json:"user"`
	Pwd          string `json:"pwd"`
	Address      string `json:"address"`
	Database     string `json:"database"`
	ShowSql      bool   `json:"show_sql"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
}

type Conf struct {
	Db      DbConf `json:"db"`
	IP      string `json:"server_ip"`
	Port    string `json:"port"`
	CdnUrl  string `json:"cdn_url"`
	Profile string `json:"profile"`
}

func LoadDefaultConf() *Conf {
	path := ConfPath()
	return LoadConf(path)
}

func SaveDefaultConf(conf *Conf) {
	path := ConfPath()
	SaveConf(path, conf)
}

func LoadConf(path *string) *Conf {
	var conf Conf
	var bs []byte
	var err error
	//read file
	if *path == "" {
		bs, err = LoadWebRootFileStr("resource", "app.json")
		if err != nil {
			panic(err)
		}
	} else {
		bs, err = LoadOsFileStr(*path)
		if err != nil {
			panic(err)
		}
	}

	//json decode
	err = json.Unmarshal(bs, &conf)
	if err != nil {
		panic(err)
	}

	//if conf ip is empty auto get it
	if conf.IP == "" {
		conf.IP = LocalAddr()
	}

	// flag -p port has the first class
	port := *(Port())
	if port != "" {
		conf.Port = port
	} else {
		if conf.Port == "" {
			conf.Port = "3000"
		}
	}

	return &conf
}

func SaveConf(path *string, conf *Conf) error {

	//json encode
	data,err := json.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	//save file
	if *path == "" {
		err = SaveWebRootFile("resource", "app.json",data)
		if err != nil {
			return err
		}
	} else {
		err = SaveOsFile(*path,data)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadWebRootFileStr(path, name string) ([]byte, error) {
	f, err := http.Dir(path).Open(name)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func SaveWebRootFile(d, name string,data []byte) error {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return errors.New("http: invalid character in file path")
	}
	dir := string(d)
	if dir == "" {
		dir = "."
	}
	f, err := os.Create(filepath.Join(dir, filepath.FromSlash(path.Clean("/"+name))))
	if err != nil {
		return  err
	}
	_,err = f.Write(data)
	return err
}

func LoadOsFileStr(name string) ([]byte, error) {
	f, err := os.Open(name)
	defer f.Close()

	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func SaveOsFile(name string, data []byte) error {

	return ioutil.WriteFile(name, data, 0666)
}
