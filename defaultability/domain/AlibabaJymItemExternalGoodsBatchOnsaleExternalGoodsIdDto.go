package domain


type AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto struct {
    /*
        外部商品ID，用于标识外部系统每次提交过来的商品     */
    ExternalGoodsId  *string `json:"external_goods_id,omitempty" `

    /*
        交易猫商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto) SetExternalGoodsId(v string) *AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto {
    s.ExternalGoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto {
    s.GoodsId = &v
    return s
}
