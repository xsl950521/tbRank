package response

import "goredis/topsdk/defaultability/domain"

type TaobaoInteractionAvatarQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回值
	*/
	Result domain.TaobaoInteractionAvatarQueryBizResult `json:"result,omitempty" `
}
