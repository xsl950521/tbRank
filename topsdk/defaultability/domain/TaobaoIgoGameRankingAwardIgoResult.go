package domain


type TaobaoIgoGameRankingAwardIgoResult struct {
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

func (s *TaobaoIgoGameRankingAwardIgoResult) SetSuccess(v bool) *TaobaoIgoGameRankingAwardIgoResult {
    s.Success = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardIgoResult) SetCompleted(v bool) *TaobaoIgoGameRankingAwardIgoResult {
    s.Completed = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardIgoResult) SetRetryable(v bool) *TaobaoIgoGameRankingAwardIgoResult {
    s.Retryable = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardIgoResult) SetBizCode(v string) *TaobaoIgoGameRankingAwardIgoResult {
    s.BizCode = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardIgoResult) SetBizMsg(v string) *TaobaoIgoGameRankingAwardIgoResult {
    s.BizMsg = &v
    return s
}
func (s *TaobaoIgoGameRankingAwardIgoResult) SetModel(v bool) *TaobaoIgoGameRankingAwardIgoResult {
    s.Model = &v
    return s
}
