package domain


type AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto struct {
    /*
        批量上架商品id集合     */
    ExternalGoodsIdList  *[]AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto `json:"external_goods_id_list,omitempty" `

    /*
        外部批次ID，用于幂等     */
    ExternalBatchId  *string `json:"external_batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto) SetExternalGoodsIdList(v []AlibabaJymItemExternalGoodsBatchOnsaleExternalGoodsIdDto) *AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto {
    s.ExternalGoodsIdList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto) SetExternalBatchId(v string) *AlibabaJymItemExternalGoodsBatchOnsaleGoodsOnSaleCommandDto {
    s.ExternalBatchId = &v
    return s
}
