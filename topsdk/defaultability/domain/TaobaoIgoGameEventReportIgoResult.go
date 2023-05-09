package domain


type TaobaoIgoGameEventReportIgoResult struct {
    /*
        调用是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        是否完成     */
    Completed  *bool `json:"completed,omitempty" `

    /*
        调用失败后能否重试     */
    Retryable  *bool `json:"retryable,omitempty" `

    /*
        错误码     */
    BizCode  *string `json:"biz_code,omitempty" `

    /*
        错误描述     */
    BizMsg  *string `json:"biz_msg,omitempty" `

    /*
        固定字段     */
    Model  *bool `json:"model,omitempty" `

}

func (s *TaobaoIgoGameEventReportIgoResult) SetSuccess(v bool) *TaobaoIgoGameEventReportIgoResult {
    s.Success = &v
    return s
}
func (s *TaobaoIgoGameEventReportIgoResult) SetCompleted(v bool) *TaobaoIgoGameEventReportIgoResult {
    s.Completed = &v
    return s
}
func (s *TaobaoIgoGameEventReportIgoResult) SetRetryable(v bool) *TaobaoIgoGameEventReportIgoResult {
    s.Retryable = &v
    return s
}
func (s *TaobaoIgoGameEventReportIgoResult) SetBizCode(v string) *TaobaoIgoGameEventReportIgoResult {
    s.BizCode = &v
    return s
}
func (s *TaobaoIgoGameEventReportIgoResult) SetBizMsg(v string) *TaobaoIgoGameEventReportIgoResult {
    s.BizMsg = &v
    return s
}
func (s *TaobaoIgoGameEventReportIgoResult) SetModel(v bool) *TaobaoIgoGameEventReportIgoResult {
    s.Model = &v
    return s
}
