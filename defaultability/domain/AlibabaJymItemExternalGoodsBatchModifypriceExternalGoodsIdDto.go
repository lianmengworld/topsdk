package domain


type AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto struct {
    /*
        交易猫商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

    /*
        外部商品ID，用于标识外部系统每次提交过来的商品     */
    ExternalGoodsId  *string `json:"external_goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto {
    s.GoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto) SetExternalGoodsId(v string) *AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto {
    s.ExternalGoodsId = &v
    return s
}
