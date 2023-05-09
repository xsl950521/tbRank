package config

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

type DB struct {
	DbUser string `json:"db_user"`
	DbPass string `json:"db_pass"`
	DbAddr string `json:"db_hostip"`
	DbName string `json:"db_name"`
	DbPort string `json:"db_port"`
}

type RedisDB struct {
	RedisUser string `json:"db_user"`
	RedisPass string `json:"db_pass"`
	RedisAddr string `json:"db_hostip"`
	RedisName int    `json:"db_name"`
	RedisPort int    `json:"db_port"`
}

type HttpConfig struct {
	HostIp string `json:"host_ip"`
	Port   int    `json:"port"`
}

type GameConfig struct {
	RoomGameid int `json:"roomgameid"`
}

type CronTimeConfig struct {
	Time_Post_SQL     string `json:"post_mysql"`
	Time_Post_TB      string `json:"post_tb"`
	Time_REFRESH_DATA string `json:"refresh_data"`
}

var Config_Sql DB

var Config_Redis RedisDB

var Config_Http HttpConfig

var Config_Game GameConfig

var Config_CronTime CronTimeConfig

func Init() {
	httpInit()
	sqlInit()
	redisInit()
	gameInit()
	crontimeInit()
}
func httpInit() {
	//打开文件
	file, _ := os.Open("./conf/httpconfig.json")
	//关闭文件
	defer file.Close()
	// NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&Config_Http)
	if err != nil {
		panic(err)
	}
	logrus.Info("Config_Http:", Config_Http)
}

func sqlInit() {
	//打开文件
	file, _ := os.Open("./conf/dbconfig.json")
	//关闭文件
	defer file.Close()
	// NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&Config_Sql)
	if err != nil {
		panic(err)
	}
	logrus.Info("Config_Sql:", Config_Sql)
}

func redisInit() {
	//打开文件
	file, _ := os.Open("./conf/redisconfig.json")
	//关闭文件
	defer file.Close()
	// NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoderredis := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	errredis := decoderredis.Decode(&Config_Redis)
	if errredis != nil {
		panic(errredis)
	}
	logrus.Info("Config_Redis:", Config_Redis)
}

func gameInit() {
	//打开文件
	file, _ := os.Open("./conf/gameconfig.json")
	//关闭文件
	defer file.Close()
	// NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoderredis := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	errredis := decoderredis.Decode(&Config_Game)
	if errredis != nil {
		panic(errredis)
	}
	logrus.Info("Config_Game:", Config_Game)
}

func crontimeInit() {
	//打开文件
	file, _ := os.Open("./conf/corntime.json")
	//关闭文件
	defer file.Close()
	// NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoderredis := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	errredis := decoderredis.Decode(&Config_CronTime)
	if errredis != nil {
		panic(errredis)
	}
	logrus.Info("Config_CronTime:", Config_CronTime)
}
