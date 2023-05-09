package domain


type TaobaoInteractionGameEventGameEventDTO struct {
    /*
        12414     */
    UserId  *string `json:"user_id,omitempty" `

    /*
        事件类型     */
    Type  *string `json:"type,omitempty" `

    /*
        事件影响数值     */
    Delta  *int64 `json:"delta,omitempty" `

    /*
        事件的唯一ID，用于去重统计。     */
    Id  *string `json:"id,omitempty" `

    /*
        业务打标主场景;用于数据归类     */
    Scope  *string `json:"scope,omitempty" `

    /*
        业务打标子场景;用于数据归类     */
    SubScope  *string `json:"sub_scope,omitempty" `

}

func (s *TaobaoInteractionGameEventGameEventDTO) SetUserId(v string) *TaobaoInteractionGameEventGameEventDTO {
    s.UserId = &v
    return s
}
func (s *TaobaoInteractionGameEventGameEventDTO) SetType(v string) *TaobaoInteractionGameEventGameEventDTO {
    s.Type = &v
    return s
}
func (s *TaobaoInteractionGameEventGameEventDTO) SetDelta(v int64) *TaobaoInteractionGameEventGameEventDTO {
    s.Delta = &v
    return s
}
func (s *TaobaoInteractionGameEventGameEventDTO) SetId(v string) *TaobaoInteractionGameEventGameEventDTO {
    s.Id = &v
    return s
}
func (s *TaobaoInteractionGameEventGameEventDTO) SetScope(v string) *TaobaoInteractionGameEventGameEventDTO {
    s.Scope = &v
    return s
}
func (s *TaobaoInteractionGameEventGameEventDTO) SetSubScope(v string) *TaobaoInteractionGameEventGameEventDTO {
    s.SubScope = &v
    return s
}
