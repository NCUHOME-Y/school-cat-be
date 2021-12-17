package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)


var (

	AppMode  string
	HttpPort string

	//DB     string
	DBHost string
	DBPort string
	DBUser string
	DBPWD  string
	DBName string

	JwtKey string
)



func init (){
/*	fmt.Println("选择配置文件：")
	var ConfigFileName string
	n, err := fmt.Scanln(&ConfigFileName)
	if err != nil || n != 1 {
		return
	}
	file,err := ini.Load("./conf/"+ConfigFileName+".ini")*/

	file,err := ini.Load("./config/conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误",err)
	}
	LoadSever(file)
	LoadData(file)
	LoadKey(file)
}


func LoadSever (file *ini.File) {
	AppMode = file.Section("sever").Key("AppMode").MustString("debug")
	HttpPort = file.Section("sever").Key("HttpPort").MustString(":8080")

}

func LoadData (file *ini.File)  {
	//DB = file.Section("database").Key("db").MustString("mysql")
	DBHost = file.Section("database").Key("dbHost").MustString("localhost")
	DBPort = file.Section("database").Key("dbPort").MustString("3306")
	DBUser = file.Section("database").Key("dbUser").MustString("root")
	DBPWD = file.Section("database").Key("dbPWD").MustString("nil")
	DBName = file.Section("database").Key("dbName").MustString("test")

}
func LoadKey (file *ini.File){
	JwtKey = file.Section("KEY").Key("JwtKey").MustString("hello")
}
