package timetask

import (
	"goredis/config"
	"goredis/define"
	mysqlserver "goredis/mysqlServer"
	rediserver "goredis/redis"
	tbserversdk "goredis/tbServerSdk"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

var toprank define.PlayerRankInfoSql
var pRank int

func postData2Sql() {
	pRank = 100
	var rankmap = make(map[int]define.PlayerRankInfoSql)
	rankinfo, err := rediserver.GetMatchTopN(pRank)
	if err != nil {
		return
	}
	for _, z := range rankinfo {
		ranknum, _ := rediserver.GetRankNum(z)
		numid, _ := strconv.Atoi(z.Member.(string))
		toprank.Numid = numid
		toprank.Rank = ranknum
		toprank.Score = z.Score
		rankmap[ranknum] = toprank
	}
	if len(rankmap) >= 1 {
		mysqlserver.MatchInfoPost(rankmap)
	} else {
		logrus.Warn("排行榜为空:", len(rankmap))
	}

}

// var timeconfig config.CronTimeConfig

func EveryDay4Time() {
	timeconfig := config.Config_CronTime
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	s := gocron.NewScheduler(timezone)
	// 每日4点上抛排行榜数据 前100名 给淘宝和我们自己的mysql
	logrus.Info("sql time=", timeconfig.Time_Post_SQL, ",p2tb time=", timeconfig.Time_Post_TB)
	s.Every(1).Day().At(timeconfig.Time_Post_SQL).Do(func() {
		go postData2Sql()
	})
	s.Every(1).Day().At(timeconfig.Time_REFRESH_DATA).Do(func() {
		go rediserver.MatchRoundAdd()
	})
	s.Every(1).Day().At(timeconfig.Time_Post_TB).Do(func() {
		go tbserversdk.RankPost2TB()
	})
	// s.Every(1).Day().At(timeconfig.Time_Post_TB).Do(func() {
	// 	go tbserversdk.PostTestAll()
	// })
	// s.Every(1).Second().Do(func() {
	// 	go cron2()
	// })

	s.StartAsync()
}
