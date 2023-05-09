package response

import "goredis/topsdk/defaultability/domain"

type TaobaoInteractionGameEventResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	Result domain.TaobaoInteractionGameEventBizResult `json:"result,omitempty" `
}
