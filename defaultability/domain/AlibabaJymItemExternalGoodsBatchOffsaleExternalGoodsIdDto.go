package domain


type AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto struct {
    /*
        外部商品ID，用于标识外部系统每次提交过来的商品     */
    ExternalGoodsId  *string `json:"external_goods_id,omitempty" `

    /*
        交易猫商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto) SetExternalGoodsId(v string) *AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto {
    s.ExternalGoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto {
    s.GoodsId = &v
    return s
}
