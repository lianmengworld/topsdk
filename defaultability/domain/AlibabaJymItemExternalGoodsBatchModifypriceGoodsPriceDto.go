package domain


type AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto struct {
    /*
        外部商品id对象     */
    ExternalGoodsId  *AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto `json:"external_goods_id,omitempty" `

    /*
        商品价格     */
    Price  *string `json:"price,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto) SetExternalGoodsId(v AlibabaJymItemExternalGoodsBatchModifypriceExternalGoodsIdDto) *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto {
    s.ExternalGoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto) SetPrice(v string) *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto {
    s.Price = &v
    return s
}
