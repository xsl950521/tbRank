package response

import "goredis/topsdk/defaultability/domain"

type TaobaoIgoGameEventReportResponse struct {

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
	Result domain.TaobaoIgoGameEventReportIgoResult `json:"result,omitempty" `
}
