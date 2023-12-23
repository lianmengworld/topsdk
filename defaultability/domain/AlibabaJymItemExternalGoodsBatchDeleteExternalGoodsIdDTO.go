package domain


type AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO struct {
    /*
        交易猫商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

    /*
        外部商品ID，用于标识外部系统每次提交过来的商品     */
    ExternalGoodsId  *string `json:"external_goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO {
    s.GoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO) SetExternalGoodsId(v string) *AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO {
    s.ExternalGoodsId = &v
    return s
}
