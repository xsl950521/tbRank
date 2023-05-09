package tbserversdk

import (
	"encoding/json"
	"fmt"
	mysqlserver "goredis/mysqlServer"
	rediserver "goredis/redis"
	"goredis/topsdk"
	"goredis/topsdk/defaultability"
	"goredis/topsdk/defaultability/domain"
	"goredis/topsdk/defaultability/request"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type returnSuccessData struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type returnErrorData struct {
	Success      bool   `json:"success"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type successData struct {
	Token string `json:"token"`
	Uid   string `json:"uid"`
}

const (
	appkey    = "33973648"
	appsecret = "2d415b3557f5eb4a04a02a25cbddb0b7"
	url       = "https://eco.taobao.com/router/rest" //游戏事件上报API
	appid     = 3000000073820645
)

var (
	client = topsdk.NewDefaultTopClient(appkey, appsecret, url, 20000, 20000)
)

type rankData struct {
	Value     int64  `json:"value"`
	TallyTime string `json:"tallytime"`
	Numid     int    `json:"numid"`
}

// func Send2TB(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if nil != err {
// 		logrus.Error("reqBody error", err.Error())
// 		return
// 	}
// 	logrus.Info(string(body))
// 	reqInfo := rankData{}

// 	err1 := json.Unmarshal(body, &reqInfo)
// 	if nil != err1 {
// 		logrus.Error("reqBody error", err1.Error())
// 		return
// 	}
// 	logrus.Info("appkey=", appkey, ",value=", reqInfo.Value)
// 	openid, _ := mysqlserver.GetNickName(reqInfo.Numid)
// 	ability := defaultability.NewDefaultability(&client)
// 	//接口请求参数
// 	taobaoIgoGameEventReportGameEventDTO := domain.TaobaoIgoGameEventReportGameEventDTO{}

// 	taobaoIgoGameEventReportGameEventDTO.SetAppId(appid)
// 	taobaoIgoGameEventReportGameEventDTO.SetAppKey(appkey)
// 	taobaoIgoGameEventReportGameEventDTO.SetOpenId(openid)
// 	taobaoIgoGameEventReportGameEventDTO.SetType("mj_challenge")
// 	taobaoIgoGameEventReportGameEventDTO.SetDelta(reqInfo.Value) //差值
// 	// taobaoIgoGameEventReportGameEventDTO.SetMissionId(123)
// 	taobaoIgoGameEventReportGameEventDTO.SetTime(time.Now().Unix())
// 	id := fmt.Sprintf("%v%v", openid, reqInfo.TallyTime)
// 	taobaoIgoGameEventReportGameEventDTO.SetId(id)
// 	/*
// 	   {"key":"value"}
// 	*/
// 	// taobaoIgoGameEventReportGameEventDTO.SetExtendInfo(make(map[string]interface{}))

// 	req := request.TaobaoIgoGameEventReportRequest{}
// 	req.SetEvent(taobaoIgoGameEventReportGameEventDTO)

// 	resp, err := ability.TaobaoIgoGameEventReport(&req)
// 	if err != nil {
// 		returnError(w, 200, err.Error())
// 		logrus.Error(err)
// 	} else {
// 		logrus.Info(resp.Body)
// 		returnSuccess(w, "success")
// 	}

// }

// const (
// 	api_award = "taobao.igo.game.event.report" //排名权益发奖API
// )

var (
	client_award = topsdk.NewDefaultTopClient(appkey, appsecret, url, 20000, 20000)
)

func PostTestAll() {
	logrus.Info("PostTest")
	rankinfo, err1 := rediserver.GetMatchTopNTest(100)
	if err1 != nil {
		logrus.Error("err1 in PostTest:", err1)
	}
	// var t time.Time = time.Now()
	//跨年有bug
	//特殊处理下
	timestr := "20230101"
	// timestr := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()-1)
	ability := defaultability.NewDefaultability(&client_award)

	// 接口请求参数
	taobaoIgoGameRankingAwardGameRankingDTO := domain.TaobaoIgoGameRankingAwardGameRankingDTO{}
	for _, z := range rankinfo {
		numid, _ := strconv.Atoi(z.Member.(string))
		olduserid, _ := mysqlserver.GetOldUserid(numid)
		var openid string
		res := strings.Split(olduserid, "scmj_")
		openid = res[len(res)-1]
		ranknum, _ := rediserver.GetRankNumTest(z)

		taobaoIgoGameRankingAwardGameRankingDTO.SetAppId(appid)
		taobaoIgoGameRankingAwardGameRankingDTO.SetAppKey(appkey)
		taobaoIgoGameRankingAwardGameRankingDTO.SetRanking(int64(ranknum))
		taobaoIgoGameRankingAwardGameRankingDTO.SetMissionId(10032)
		taobaoIgoGameRankingAwardGameRankingDTO.SetBizTime(timestr)
		id := fmt.Sprintf("%v%v", openid, time.Now().Unix())
		taobaoIgoGameRankingAwardGameRankingDTO.SetId(id)
		/*
		   {"key":"value"}
		*/
		// taobaoIgoGameRankingAwardGameRankingDTO.SetExtendInfo(make(map[string]interface{}))
		taobaoIgoGameRankingAwardGameRankingDTO.SetOpenId(openid)

		req := request.TaobaoIgoGameRankingAwardRequest{}
		req.SetDto(taobaoIgoGameRankingAwardGameRankingDTO)

		resp, err := ability.TaobaoIgoGameRankingAward(&req)
		if err != nil {
			logrus.Error("err=", err)
		} else {
			logrus.Info("Body=", resp.Body)
		}
	}
}

func PostTestOne() {
	logrus.Info("PostTest")
	numid, ranknum, _, err1 := rediserver.GetMatchNTest(101)
	if err1 != nil {
		logrus.Error("err1 in PostTest:", err1)
	}
	// var t time.Time = time.Now().AddDate(0, 0, -1) //昨天
	// timestr := t.Format("20060102")
	//跨年有bug
	//特殊处理下
	timestr := "20230101"

	// timestr := fmt.Sprintf("%d%d%d", t.Year(), t.Month(), t.Day()-1)
	ability := defaultability.NewDefaultability(&client_award)

	// 接口请求参数
	taobaoIgoGameRankingAwardGameRankingDTO := domain.TaobaoIgoGameRankingAwardGameRankingDTO{}
	// numid, _ := strconv.Atoi(z.Member.(string))
	olduserid, _ := mysqlserver.GetOldUserid(numid)
	var openid string
	res := strings.Split(olduserid, "scmj_")
	openid = res[len(res)-1]
	// ranknum, _ := rediserver.GetRankNumTest(z)

	taobaoIgoGameRankingAwardGameRankingDTO.SetAppId(appid)
	taobaoIgoGameRankingAwardGameRankingDTO.SetAppKey(appkey)
	taobaoIgoGameRankingAwardGameRankingDTO.SetRanking(int64(ranknum))
	taobaoIgoGameRankingAwardGameRankingDTO.SetMissionId(10032)
	taobaoIgoGameRankingAwardGameRankingDTO.SetBizTime(timestr)
	id := fmt.Sprintf("%v%v", openid, time.Now().Unix())
	taobaoIgoGameRankingAwardGameRankingDTO.SetId(id)
	/*
	   {"key":"value"}
	*/
	// taobaoIgoGameRankingAwardGameRankingDTO.SetExtendInfo(make(map[string]interface{}))
	taobaoIgoGameRankingAwardGameRankingDTO.SetOpenId(openid)

	req := request.TaobaoIgoGameRankingAwardRequest{}
	req.SetDto(taobaoIgoGameRankingAwardGameRankingDTO)

	resp, err := ability.TaobaoIgoGameRankingAward(&req)
	if err != nil {
		logrus.Error("err=", err)
	} else {
		logrus.Info("Body=", resp.Body)
	}
}

func RankPost2TB() {
	logrus.Info("RankPost2TB")
	rankinfo, err1 := rediserver.GetMatchTopN(100)
	if err1 != nil {
		logrus.Error("err1 in RankPost2TB:", err1)
	}
	var t time.Time = time.Now().AddDate(0, 0, -1) //昨天
	timestr := t.Format("20060102")
	ability := defaultability.NewDefaultability(&client_award)
	//接口请求参数
	taobaoIgoGameRankingAwardGameRankingDTO := domain.TaobaoIgoGameRankingAwardGameRankingDTO{}

	for _, z := range rankinfo {
		numid, _ := strconv.Atoi(z.Member.(string))
		olduserid, _ := mysqlserver.GetOldUserid(numid)
		var openid string
		res := strings.Split(olduserid, "scmj_")
		openid = res[len(res)-1]
		ranknum, _ := rediserver.GetRankNum(z)

		taobaoIgoGameRankingAwardGameRankingDTO.SetAppId(appid)
		taobaoIgoGameRankingAwardGameRankingDTO.SetAppKey(appkey)
		taobaoIgoGameRankingAwardGameRankingDTO.SetRanking(int64(ranknum))
		taobaoIgoGameRankingAwardGameRankingDTO.SetMissionId(10032)
		taobaoIgoGameRankingAwardGameRankingDTO.SetBizTime(timestr)
		id := fmt.Sprintf("%v%v", openid, time.Now().Unix())
		taobaoIgoGameRankingAwardGameRankingDTO.SetId(id)
		/*
		   {"key":"value"}
		*/
		// taobaoIgoGameRankingAwardGameRankingDTO.SetExtendInfo(make(map[string]interface{}))
		taobaoIgoGameRankingAwardGameRankingDTO.SetOpenId(openid)

		req := request.TaobaoIgoGameRankingAwardRequest{}
		req.SetDto(taobaoIgoGameRankingAwardGameRankingDTO)

		resp, err := ability.TaobaoIgoGameRankingAward(&req)
		if err != nil {
			logrus.Error("err=", err)
		} else {
			logrus.Info("Body=", resp.Body)
		}
	}
}

// Success 向 w 中写入成功信息
func ReturnSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	returnDa := returnSuccessData{
		Success: true,
		Data:    data,
	}
	if err := json.NewEncoder(w).Encode(returnDa); err != nil {
		logrus.Error("success json_encode err")
	}
}

// Error 向 w中写入错误信息
func ReturnError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	returnDa := returnErrorData{
		Success:      false,
		ErrorCode:    fmt.Sprint(code),
		ErrorMessage: msg,
	}
	if err := json.NewEncoder(w).Encode(returnDa); err != nil {
		logrus.Error("Error json_encode err")
	}
}
