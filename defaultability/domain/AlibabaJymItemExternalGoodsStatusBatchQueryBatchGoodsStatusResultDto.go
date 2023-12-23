package domain


type AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusResultDto struct {
    /*
        商品状态列表     */
    GoodsStatusList  *[]AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto `json:"goods_status_list,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusResultDto) SetGoodsStatusList(v []AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto) *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusResultDto {
    s.GoodsStatusList = &v
    return s
}
