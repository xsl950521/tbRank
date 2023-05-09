package domain


type TaobaoInteractionGameEventBizResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        错误号     */
    MsgCode  *string `json:"msg_code,omitempty" `

    /*
        具体错误信息     */
    MsgInfo  *string `json:"msg_info,omitempty" `

    /*
        固定字段     */
    Model  *bool `json:"model,omitempty" `

    /*
        是否可重试，默认都不可以重试     */
    CanRetry  *bool `json:"can_retry,omitempty" `

}

func (s *TaobaoInteractionGameEventBizResult) SetSuccess(v bool) *TaobaoInteractionGameEventBizResult {
    s.Success = &v
    return s
}
func (s *TaobaoInteractionGameEventBizResult) SetMsgCode(v string) *TaobaoInteractionGameEventBizResult {
    s.MsgCode = &v
    return s
}
func (s *TaobaoInteractionGameEventBizResult) SetMsgInfo(v string) *TaobaoInteractionGameEventBizResult {
    s.MsgInfo = &v
    return s
}
func (s *TaobaoInteractionGameEventBizResult) SetModel(v bool) *TaobaoInteractionGameEventBizResult {
    s.Model = &v
    return s
}
func (s *TaobaoInteractionGameEventBizResult) SetCanRetry(v bool) *TaobaoInteractionGameEventBizResult {
    s.CanRetry = &v
    return s
}
