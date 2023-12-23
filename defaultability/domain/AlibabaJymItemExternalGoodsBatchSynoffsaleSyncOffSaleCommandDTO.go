package domain


type AlibabaJymItemExternalGoodsBatchSynoffsaleSyncOffSaleCommandDTO struct {
    /*
        交易猫商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchSynoffsaleSyncOffSaleCommandDTO) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsBatchSynoffsaleSyncOffSaleCommandDTO {
    s.GoodsId = &v
    return s
}
