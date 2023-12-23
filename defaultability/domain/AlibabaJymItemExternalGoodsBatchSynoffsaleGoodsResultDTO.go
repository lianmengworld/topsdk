package domain


type AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO struct {
    /*
        状态码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        描述     */
    Desc  *string `json:"desc,omitempty" `

    /*
        是否下架成功     */
    Success  *bool `json:"success,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO) SetCode(v int64) *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO {
    s.Code = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO) SetDesc(v string) *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO {
    s.Desc = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO) SetSuccess(v bool) *AlibabaJymItemExternalGoodsBatchSynoffsaleGoodsResultDTO {
    s.Success = &v
    return s
}
