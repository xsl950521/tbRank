package httpserver

import (
	"encoding/json"
	"fmt"
	"goredis/define"
	mysqlserver "goredis/mysqlServer"
	rediserver "goredis/redis"
	tbserversdk "goredis/tbServerSdk"
	"goredis/topsdk"
	"goredis/topsdk/defaultability"
	"goredis/topsdk/defaultability/domain"
	"goredis/topsdk/defaultability/request"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type MatchRankInfo struct {
	Numid      int `json:"numid"`
	MatchScore int `json:"matchscore"`
	RoundNow   int `json:"round"`
	RoomGameid int `json:"roomgameid"`
	Type       int `json:"type"` //type=1 进榜,type=2 踢出榜单
}

type GetRankParam struct {
	Type      int    `json:"type"`
	TopN      int    `json:"topN"`
	Day       string `json:"day"`
	TopMore_1 int    `json:"topmore1"`
	TopMore_2 int    `json:"topmore2"`
	TopMore_3 int    `json:"topmore3"`
	Numid     int    `json:"numid"`
}

type BackValue struct {
	Code   int                           `json:"code"`
	Msg    map[int]define.PlayerRankInfo `json:"msg"`
	ErrMsg string                        `json:"errmsg"`
}

// getRankHandle 返回码
const (
	SUCCESS         = 0    //成功
	JSONREADFAIL    = 1001 //请求参数解析失败
	BODYREADFAIL    = 1002 //头信息入去失败
	GETRANKINFOFAIL = 1003 //获取排名信息失败
)

var (
	PlayerRankInfo define.PlayerRankInfo
)

func Httpstart() {
	http.HandleFunc("/addrank", addRankHandle)
	http.HandleFunc("/getrank", getRankHandle)
	http.HandleFunc("/posttopsdk", posttopsdkRankHandle)
	// http.HandleFunc("/csp/random", csp_random_Handler)
	http.ListenAndServe(":9016", nil)
}

type rankData struct {
	Value     int64 `json:"value"`
	TallyTime int   `json:"tallytime"`
	Numid     int   `json:"numid"`
}

const (
	appkey    = "33973648"
	appsecret = "2d415b3557f5eb4a04a02a25cbddb0b7"
	url_tb    = "https://eco.taobao.com/router/rest" //游戏事件上报API
	appid     = 3000000073820645
)

var (
	client = topsdk.NewDefaultTopClient(appkey, appsecret, url_tb, 20000, 20000)
)

func posttopsdkRankHandle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		logrus.Error("reqBody error", err.Error())
		return
	}
	logrus.Info(string(body))
	reqInfo := rankData{}

	err1 := json.Unmarshal(body, &reqInfo)
	if nil != err1 {
		logrus.Error("reqBody error", err1.Error())
		return
	}
	logrus.Info("appkey=", appkey, ",value=", reqInfo.Value)
	openid, _ := mysqlserver.GetOpenid(reqInfo.Numid)
	ability := defaultability.NewDefaultability(&client)
	//接口请求参数
	taobaoIgoGameEventReportGameEventDTO := domain.TaobaoIgoGameEventReportGameEventDTO{}

	taobaoIgoGameEventReportGameEventDTO.SetAppId(appid)
	taobaoIgoGameEventReportGameEventDTO.SetAppKey(appkey)
	taobaoIgoGameEventReportGameEventDTO.SetOpenId(openid)
	taobaoIgoGameEventReportGameEventDTO.SetType("mj_challenge")
	taobaoIgoGameEventReportGameEventDTO.SetDelta(reqInfo.Value) //差值
	// taobaoIgoGameEventReportGameEventDTO.SetMissionId(123)
	taobaoIgoGameEventReportGameEventDTO.SetTime(time.Now().UnixMilli())
	id := fmt.Sprintf("%v%v", openid, reqInfo.TallyTime)
	taobaoIgoGameEventReportGameEventDTO.SetId(id)
	/*
	   {"key":"value"}
	*/
	// taobaoIgoGameEventReportGameEventDTO.SetExtendInfo(make(map[string]interface{}))

	req := request.TaobaoIgoGameEventReportRequest{}
	req.SetEvent(taobaoIgoGameEventReportGameEventDTO)

	resp, err := ability.TaobaoIgoGameEventReport(&req)
	if err != nil {
		tbserversdk.ReturnError(w, 200, err.Error())
		logrus.Error(err)
	} else {
		logrus.Info(resp.Body)
		tbserversdk.ReturnSuccess(w, "success")
	}
}
func getRankHandle(w http.ResponseWriter, r *http.Request) {
	bodyContent, err := io.ReadAll(r.Body)
	var rankmap = make(map[int]define.PlayerRankInfo)
	var backvalue BackValue

	defer r.Body.Close()
	if err != nil {
		backvalue.Code = BODYREADFAIL
		backvalue.ErrMsg = err.Error()
		logrus.Error("read body fai")
		msgNew, _ := json.Marshal(backvalue)
		w.Write(msgNew)
		return
	}

	// fmt.Printf("body:%v\n", bodyContent)

	//解析json
	var conf GetRankParam
	errj := json.Unmarshal(bodyContent, &conf)
	if errj != nil {
		backvalue.Code = JSONREADFAIL
		backvalue.ErrMsg = errj.Error()
		logrus.Error("Unmarshal body fail")
		msgNew, _ := json.Marshal(backvalue)
		w.Write(msgNew)
		return
	}
	// fmt.Printf("data:%v,%v,%v\n", conf.Type, conf.Day, conf.TopN)
	// fmt.Printf("topmore:%v,%v,%v\n", conf.TopMore_1, conf.TopMore_2, conf.TopMore_3)

	switch conf.Type {
	case 1:
		rankinfo, err := getRankInfo(conf)
		if err != nil {
			backvalue.Code = GETRANKINFOFAIL
			backvalue.ErrMsg = err.Error()
			msgNew, _ := json.Marshal(backvalue)
			w.Write(msgNew)
			return
		}

		w.Header().Set("content-type", "text/json")
		rankmap = rankinfo

		backvalue.Code = SUCCESS
		backvalue.Msg = rankmap
		msgNew, _ := json.Marshal(backvalue)
		w.Write(msgNew)
	}
}

func getRankInfo(conf GetRankParam) (info map[int]define.PlayerRankInfo, err error) {
	rankinfo, err1 := rediserver.GetMatchTopN(conf.TopN)
	if err1 != nil {
		return nil, err1
	}
	var toprank define.PlayerRankInfo
	var rankmap = make(map[int]define.PlayerRankInfo)
	toprank.IsMyInfo = false
	for i, z := range rankinfo {
		i++
		ranknum, _ := rediserver.GetRankNum(z)
		numid, _ := strconv.Atoi(z.Member.(string))
		toprank.Numid = numid
		toprank.Rank = ranknum
		toprank.Score = int(math.Floor(z.Score))
		toprank.NickName, _ = mysqlserver.GetNickName(numid)
		rankmap[i] = toprank
	}
	lenfirst := len(rankmap)
	if lenfirst < conf.TopN {
		for i := 1; i <= (conf.TopN - lenfirst); i++ {
			toprank.Numid = 0
			toprank.Rank = 0
			toprank.Score = 0
			toprank.NickName = ""
			rankmap[lenfirst+i] = toprank
		}
	}

	lensecond := len(rankmap)
	for i := 1; i < 4; i++ {
		var t int
		if i == 1 {
			t = conf.TopMore_1
		} else if i == 2 {
			t = conf.TopMore_2
		} else if i == 3 {
			t = conf.TopMore_3
		}
		// logrus.Info("tttt=", t)
		numid, numrank, score, _ := rediserver.GetMatchN(t)
		if numid != -1 {
			toprank.Numid = numid
			toprank.Rank = numrank
			toprank.Score = int(math.Floor(score))
			toprank.NickName, _ = mysqlserver.GetNickName(numid)
		} else {
			toprank.Numid = 0
			toprank.Rank = 0
			toprank.Score = 0
			toprank.NickName = ""
		}
		rankmap[lensecond+i] = toprank
	}

	if conf.Numid != 0 {
		toprank.IsMyInfo = true
		myscore, myrank, err2 := rediserver.GetMemberMatchInfo(conf.Numid)
		if err2 != nil {
			logrus.Error("err2=", err2.Error())
			toprank.Numid = conf.Numid
			toprank.Rank = 0
			toprank.Score = 0
			toprank.NickName, _ = mysqlserver.GetNickName(conf.Numid)
		}
		toprank.NickName, _ = mysqlserver.GetNickName(conf.Numid)
		toprank.Numid = conf.Numid
		toprank.Rank = int(myrank)
		toprank.Score = int(math.Floor(myscore))
		rankmap[len(rankmap)+1] = toprank
	}

	return rankmap, nil
}
func addRankHandle(w http.ResponseWriter, r *http.Request) {
	// command :=r.FormValue("command")

	fmt.Println(r.Host)
	//读取请求体信息
	bodyContent, err := io.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		return
	}
	//解析json
	var conf MatchRankInfo
	errj := json.Unmarshal(bodyContent, &conf)
	if errj != nil {
		return
	}
	// fmt.Printf("data:%d,%d,%d\n", conf.Numid, conf.NowDay, conf.MatchScore)
	//格式化比赛分数
	flscore, err := formatScore(conf.MatchScore)
	if err != nil {
		log.Printf("format fail err:%v", err)
	}
	// fmt.Printf("flscore=%f\n", flscore)
	//加入/更新排行榜
	numid_str := fmt.Sprintf("%d", conf.Numid)
	if conf.Type == 1 {
		go rediserver.AddRank(conf.RoomGameid, redis.Z{Score: flscore, Member: numid_str})
	} else if conf.Type == 2 {
		go rediserver.RmvRank(conf.RoomGameid, redis.Z{Score: flscore, Member: numid_str})
	}

	//返回响应内容
	fmt.Fprintf(w, "add to rank success ~")
}

// 格式化score 将真正分数存为整数位,将第二天0点时间减去获取分数时的时间作为小数位
func formatScore(matchscore int) (fmscore float64, err error) {
	//现在的时间
	currentTime := time.Now()
	//第二天23.59.59的时间
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 23, 59, 59, 0, currentTime.Location())

	score := strconv.Itoa(matchscore)
	tim := strconv.Itoa(int(endTime.Unix() - currentTime.Unix()))

	flscore := fmt.Sprintf("%s.%s", score, tim)

	v1, err := strconv.ParseFloat(flscore, 64)
	if err == nil {
		logrus.Info("fscore=", v1)
		return v1, nil
	}
	return 0, err
}
