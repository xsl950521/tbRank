package request

import (
	"goredis/topsdk/defaultability/domain"
	"goredis/topsdk/util"
)

type TaobaoIgoGameEventReportRequest struct {
	/*
	   游戏事件     */
	Event *domain.TaobaoIgoGameEventReportGameEventDTO `json:"event" required:"true" `
}

func (s *TaobaoIgoGameEventReportRequest) SetEvent(v domain.TaobaoIgoGameEventReportGameEventDTO) *TaobaoIgoGameEventReportRequest {
	s.Event = &v
	return s
}

func (req *TaobaoIgoGameEventReportRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Event != nil {
		paramMap["event"] = util.ConvertStruct(*req.Event)
	}
	return paramMap
}

func (req *TaobaoIgoGameEventReportRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
