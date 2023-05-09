package domain


type TaobaoInteractionAvatarQueryAvatarDTO struct {
    /*
        头像 URL 地址     */
    AvatarUrl  *string `json:"avatar_url,omitempty" `

    /*
        逐帧描述信息URL地址     */
    DescUrl  *string `json:"desc_url,omitempty" `

    /*
        头像类型     */
    Type  *string `json:"type,omitempty" `

}

func (s *TaobaoInteractionAvatarQueryAvatarDTO) SetAvatarUrl(v string) *TaobaoInteractionAvatarQueryAvatarDTO {
    s.AvatarUrl = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryAvatarDTO) SetDescUrl(v string) *TaobaoInteractionAvatarQueryAvatarDTO {
    s.DescUrl = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryAvatarDTO) SetType(v string) *TaobaoInteractionAvatarQueryAvatarDTO {
    s.Type = &v
    return s
}
