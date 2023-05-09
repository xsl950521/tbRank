package request

import (
	"goredis/topsdk/defaultability/domain"
	"goredis/topsdk/util"
)

type TaobaoInteractionGameEventRequest struct {
	/*
	   事件信息     */
	Event *domain.TaobaoInteractionGameEventGameEventDTO `json:"event" required:"true" `
}

func (s *TaobaoInteractionGameEventRequest) SetEvent(v domain.TaobaoInteractionGameEventGameEventDTO) *TaobaoInteractionGameEventRequest {
	s.Event = &v
	return s
}

func (req *TaobaoInteractionGameEventRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Event != nil {
		paramMap["event"] = util.ConvertStruct(*req.Event)
	}
	return paramMap
}

func (req *TaobaoInteractionGameEventRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
