package domain


type AlibabaJymItemExternalGoodsDetailQueryGoodsDetailQueryDto struct {
    /*
        交易猫商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsDetailQueryDto) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsDetailQueryDto {
    s.GoodsId = &v
    return s
}
