package domain


type TaobaoInteractionAvatarQueryBizResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        具体内容     */
    Model  *TaobaoInteractionAvatarQueryAvatarDTO `json:"model,omitempty" `

    /*
        信息码     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        信息详情     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        是否可重试     */
    CanRetry  *bool `json:"can_retry,omitempty" `

}

func (s *TaobaoInteractionAvatarQueryBizResult) SetSuccess(v bool) *TaobaoInteractionAvatarQueryBizResult {
    s.Success = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryBizResult) SetModel(v TaobaoInteractionAvatarQueryAvatarDTO) *TaobaoInteractionAvatarQueryBizResult {
    s.Model = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryBizResult) SetMsgCode(v string) *TaobaoInteractionAvatarQueryBizResult {
    s.MsgCode = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryBizResult) SetMsgInfo(v string) *TaobaoInteractionAvatarQueryBizResult {
    s.MsgInfo = &v
    return s
}
func (s *TaobaoInteractionAvatarQueryBizResult) SetCanRetry(v bool) *TaobaoInteractionAvatarQueryBizResult {
    s.CanRetry = &v
    return s
}
