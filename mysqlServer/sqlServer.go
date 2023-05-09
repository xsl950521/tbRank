package mysqlserver

import (
	"bytes"
	"fmt"
	"goredis/config"
	"goredis/define"
	"strings"

	rediserver "goredis/redis"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMysql() {
	// config.Init()
	logrus.Info("init mysql")
	dbConfig := config.Config_Sql
	db_user := dbConfig.DbUser
	db_pass := dbConfig.DbPass
	db_name := dbConfig.DbName
	db_addr := dbConfig.DbAddr
	db_port := dbConfig.DbPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=latin1&parseTime=True&loc=Local", db_user, db_pass, db_addr, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil != err {
		panic(err)
	}
	Db = db
}

func MatchInfoPost(rankinfo map[int]define.PlayerRankInfoSql) {
	// now := time.Now().Format("2006-01-02 15:04:05")
	var buffer bytes.Buffer
	gameconfig := config.Config_Game
	roomgameid := gameconfig.RoomGameid
	roundm := rediserver.GetMatchRound(roomgameid)
	sql := "insert into `game_tb_rank` (`numid`, `score`, `nickname`, `ranknum`, `round` ) values"
	if _, err := buffer.WriteString(sql); err != nil {
		logrus.Error("sql error", err.Error())
		return
	}

	n := 0
	logrus.Info("len=", len(rankinfo))
	for _, v := range rankinfo {
		nickname, _ := GetNickName(v.Numid)
		if n == len(rankinfo)-1 {
			buffer.WriteString(fmt.Sprintf("('%d','%f','%s','%d','%d');", v.Numid, v.Score, nickname, v.Rank, roundm))
		} else {
			buffer.WriteString(fmt.Sprintf("('%d','%f','%s','%d','%d'),", v.Numid, v.Score, nickname, v.Rank, roundm))
		}
		n++
	}

	err := Db.Exec(buffer.String()).Error
	if nil != err {
		logrus.Error("sql error", err.Error())
		return
	} else {
		logrus.Info("pushToDB success")
	}
}

type game_users struct {
	NickName  string `gorm:"column:nickname"`
	OldUserid string `gorm:"column:olduserid"`
}

func GetNickName(numid int) (nickname string, e error) {
	infos := game_users{}
	err := Db.Debug().Where("numid = ?", numid).First(&infos).Error
	if nil != err {
		logrus.Error("GetNickName error", err.Error())
		return "", err
	} else {
		infos.NickName = rediserver.ConvertToString(infos.NickName, "gbk", "utf8")
		return infos.NickName, nil
	}
}

func GetOldUserid(numid int) (olduserid string, e error) {
	infos := game_users{}
	err := Db.Debug().Where("numid = ?", numid).First(&infos).Error
	if nil != err {
		logrus.Error("GetOpenid error", err.Error())
		return "", err
	} else {
		return infos.OldUserid, nil
	}
}

func GetOpenid(numid int) (openid string, e error) {
	olduserid, err1 := GetOldUserid(numid)
	if err1 != nil {
		logrus.Error("err1=", err1)
		return "", err1
	}
	res := strings.Split(olduserid, "scmj_")
	openid = res[len(res)-1]
	return openid, nil
	// infos := tb_serrion_data{}
	// err := Db.Debug().Where("userid = ?", olduserid).First(&infos).Error
	// if nil != err {
	// 	logrus.Error("GetOpenid error", err.Error())
	// 	return "", err
	// } else {
	// 	return infos.Openid, nil
	// }
}
