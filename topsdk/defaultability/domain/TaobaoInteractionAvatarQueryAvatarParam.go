package domain


type TaobaoInteractionAvatarQueryAvatarParam struct {
    /*
        用户 id     */
    OpenId  *string `json:"open_id,omitempty" `

    /*
         服务 appkey     */
    AppKey  *string `json:"app_key,omitempty" `

    /*
        TB_LIFE_MJ 麻将场景使用人生形象     */
    Type  *string `json:"type,omitempty" `

}

func (s *TaobaoInteractionAvatarQueryAvatarParam) SetOpenId(v string) *TaobaoInteractionAvatarQueryAvatarParam {
    s.OpenId = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryAvatarParam) SetAppKey(v string) *TaobaoInteractionAvatarQueryAvatarParam {
    s.AppKey = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryAvatarParam) SetType(v string) *TaobaoInteractionAvatarQueryAvatarParam {
    s.Type = &v
    return s
}
