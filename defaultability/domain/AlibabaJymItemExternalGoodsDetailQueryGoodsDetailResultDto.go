package domain


type AlibabaJymItemExternalGoodsDetailQueryGoodsDetailResultDto struct {
    /*
        商品详情信息     */
    GoodsDetail  *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto `json:"goods_detail,omitempty" `

    /*
        商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsDetailResultDto) SetGoodsDetail(v AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) *AlibabaJymItemExternalGoodsDetailQueryGoodsDetailResultDto {
    s.GoodsDetail = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsDetailResultDto) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsDetailResultDto {
    s.GoodsId = &v
    return s
}
