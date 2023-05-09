package domain


type TaobaoIgoGameRankingAwardGameRankingDTO struct {
    /*
        应用 appId     */
    AppId  *int64 `json:"app_id,omitempty" `

    /*
        应用appKey     */
    AppKey  *string `json:"app_key,omitempty" `

    /*
        用户排名     */
    Ranking  *int64 `json:"ranking,omitempty" `

    /*
        任务id,平台分配     */
    MissionId  *int64 `json:"mission_id,omitempty" `

    /*
        业务时间，yyyymmdd格式     */
    BizTime  *string `json:"biz_time,omitempty" `

    /*
        事件唯一ID,用户幂等     */
    Id  *string `json:"id,omitempty" `

    /*
        扩展参数,可空     */
    ExtendInfo  *string `json:"extend_info,omitempty" `

    /*
        用户openid     */
    OpenId  *string `json:"open_id,omitempty" `

}

func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetAppId(v int64) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.AppId = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetAppKey(v string) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.AppKey = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetRanking(v int64) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.Ranking = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetMissionId(v int64) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.MissionId = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetBizTime(v string) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.BizTime = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetId(v string) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.Id = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetExtendInfo(v string) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.ExtendInfo = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardGameRankingDTO) SetOpenId(v string) *TaobaoIgoGameRankingAwardGameRankingDTO {
    s.OpenId = &v
    return s
}
