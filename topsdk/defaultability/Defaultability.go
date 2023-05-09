package defaultability

import (
	"errors"
	"goredis/topsdk"
	"goredis/topsdk/defaultability/request"
	"goredis/topsdk/defaultability/response"
	"goredis/topsdk/util"
	"log"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
淘宝形象查询
*/
func (ability *Defaultability) TaobaoInteractionAvatarQuery(req *request.TaobaoInteractionAvatarQueryRequest) (*response.TaobaoInteractionAvatarQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.interaction.avatar.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoInteractionAvatarQueryResponse{}
	if err != nil {
		log.Println("taobaoInteractionAvatarQuery error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
互动开放游戏事件上报API
*/
func (ability *Defaultability) TaobaoIgoGameEventReport(req *request.TaobaoIgoGameEventReportRequest) (*response.TaobaoIgoGameEventReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.igo.game.event.report", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoIgoGameEventReportResponse{}
	if err != nil {
		log.Println("taobaoIgoGameEventReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
上报游戏特定事件用于统计用户状态，激励用户行为
*/
func (ability *Defaultability) TaobaoInteractionGameEvent(req *request.TaobaoInteractionGameEventRequest) (*response.TaobaoInteractionGameEventResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.interaction.game.event", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoInteractionGameEventResponse{}
	if err != nil {
		log.Println("taobaoInteractionGameEvent error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
排名权益发奖
*/
func (ability *Defaultability) TaobaoIgoGameRankingAward(req *request.TaobaoIgoGameRankingAwardRequest) (*response.TaobaoIgoGameRankingAwardResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.igo.game.ranking.award", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoIgoGameRankingAwardResponse{}
	if err != nil {
		log.Println("taobaoIgoGameRankingAward error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
