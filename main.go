package main

import (
	"goredis/config"
	httpserver "goredis/httpServer"
	mysqlserver "goredis/mysqlServer"
	rediserver "goredis/redis"
	"os"

	mowcli "github.com/jawher/mow.cli"
	"github.com/sirupsen/logrus"
)

var (
	cli        = mowcli.App("tbRank", "淘宝比赛排行榜")
	version    = "0.0.1"
	lastTime   = "2022.12.15"
	modifiedBy = "xushilong"
)

func main() {
	cli.Command("start", "运行服务", func(cmd *mowcli.Cmd) {
		cmd.Action = func() {
			//配置文件初始化
			configInit()
			//redis初始化
			rediserver.RedisInit()
			//roomgameid初始化
			rediserver.GameidInit()
			//mysql初始化
			mysqlserver.InitMysql()
			//定时任务
			// timetask.EveryDay4Time()
			//http端口启动
			httpserver.Httpstart()
			//test
			test()
		}
	})
	if err := cli.Run(os.Args); nil != err {
		logrus.Error("启动失败")
	}

}

func configInit() {
	config.Init()
}

func test() {
	rediserver.TestRedis()
}
