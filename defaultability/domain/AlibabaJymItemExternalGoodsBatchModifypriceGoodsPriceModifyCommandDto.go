package domain


type AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto struct {
    /*
        商品价格传输对象     */
    GoodsPriceList  *[]AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto `json:"goods_price_list,omitempty" `

    /*
        外部批次ID，用于幂等     */
    ExternalBatchId  *string `json:"external_batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto) SetGoodsPriceList(v []AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceDto) *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto {
    s.GoodsPriceList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto) SetExternalBatchId(v string) *AlibabaJymItemExternalGoodsBatchModifypriceGoodsPriceModifyCommandDto {
    s.ExternalBatchId = &v
    return s
}
