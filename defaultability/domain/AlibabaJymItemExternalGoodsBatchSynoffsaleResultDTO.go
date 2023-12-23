package domain


type AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO struct {
    /*
        调用错误信息     */
    ExtraErrMsg  *string `json:"extra_err_msg,omitempty" `

    /*
        业务对象     */
    Result  *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO `json:"result,omitempty" `

    /*
        调用错误码     */
    StateCode  *string `json:"state_code,omitempty" `

    /*
        是否调用成功     */
    Success  *string `json:"success,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO) SetExtraErrMsg(v string) *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO {
    s.ExtraErrMsg = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO) SetResult(v AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO) *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO {
    s.Result = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO) SetStateCode(v string) *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO {
    s.StateCode = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO) SetSuccess(v string) *AlibabaJymItemExternalGoodsBatchSynoffsaleResultDTO {
    s.Success = &v
    return s
}
