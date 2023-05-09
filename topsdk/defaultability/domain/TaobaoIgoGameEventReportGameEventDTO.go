package domain


type TaobaoIgoGameEventReportGameEventDTO struct {
    /*
        应用 appId     */
    AppId  *int64 `json:"app_id,omitempty" `

    /*
        应用appKey     */
    AppKey  *string `json:"app_key,omitempty" `

    /*
        用户openId     */
    OpenId  *string `json:"open_id,omitempty" `

    /*
        事件类型     */
    Type  *string `json:"type,omitempty" `

    /*
        变化量     */
    Delta  *int64 `json:"delta,omitempty" `

    /*
        任务id,平台分配     */
    MissionId  *int64 `json:"mission_id,omitempty" `

    /*
        发生时间，毫秒时间戳     */
    Time  *int64 `json:"time,omitempty" `

    /*
        事件唯一ID,用户幂等     */
    Id  *string `json:"id,omitempty" `

    /*
        扩展参数,可空     */
    ExtendInfo  *string `json:"extend_info,omitempty" `

}

func (s *TaobaoIgoGameEventReportGameEventDTO) SetAppId(v int64) *TaobaoIgoGameEventReportGameEventDTO {
    s.AppId = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetAppKey(v string) *TaobaoIgoGameEventReportGameEventDTO {
    s.AppKey = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetOpenId(v string) *TaobaoIgoGameEventReportGameEventDTO {
    s.OpenId = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetType(v string) *TaobaoIgoGameEventReportGameEventDTO {
    s.Type = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetDelta(v int64) *TaobaoIgoGameEventReportGameEventDTO {
    s.Delta = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetMissionId(v int64) *TaobaoIgoGameEventReportGameEventDTO {
    s.MissionId = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetTime(v int64) *TaobaoIgoGameEventReportGameEventDTO {
    s.Time = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetId(v string) *TaobaoIgoGameEventReportGameEventDTO {
    s.Id = &v
    return s
}
func (s *TaobaoIgoGameEventReportGameEventDTO) SetExtendInfo(v string) *TaobaoIgoGameEventReportGameEventDTO {
    s.ExtendInfo = &v
    return s
}
