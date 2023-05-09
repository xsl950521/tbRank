package request

import (
	"goredis/topsdk/defaultability/domain"
	"goredis/topsdk/util"
)

type TaobaoInteractionAvatarQueryRequest struct {
	/*
	   参数     */
	Param *domain.TaobaoInteractionAvatarQueryAvatarParam `json:"param" required:"true" `
}

func (s *TaobaoInteractionAvatarQueryRequest) SetParam(v domain.TaobaoInteractionAvatarQueryAvatarParam) *TaobaoInteractionAvatarQueryRequest {
	s.Param = &v
	return s
}

func (req *TaobaoInteractionAvatarQueryRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Param != nil {
		paramMap["param"] = util.ConvertStruct(*req.Param)
	}
	return paramMap
}

func (req *TaobaoInteractionAvatarQueryRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
