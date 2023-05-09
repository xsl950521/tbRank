package request

import (
	"goredis/topsdk/defaultability/domain"
	"goredis/topsdk/util"
)

type TaobaoIgoGameRankingAwardRequest struct {
	/*
	   游戏排名信息     */
	Dto *domain.TaobaoIgoGameRankingAwardGameRankingDTO `json:"dto" required:"true" `
}

func (s *TaobaoIgoGameRankingAwardRequest) SetDto(v domain.TaobaoIgoGameRankingAwardGameRankingDTO) *TaobaoIgoGameRankingAwardRequest {
	s.Dto = &v
	return s
}

func (req *TaobaoIgoGameRankingAwardRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Dto != nil {
		paramMap["dto"] = util.ConvertStruct(*req.Dto)
	}
	return paramMap
}

func (req *TaobaoIgoGameRankingAwardRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
