package domain


type AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto struct {
    /*
        批量下架商品id集合     */
    ExternalGoodsIdList  *[]AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto `json:"external_goods_id_list,omitempty" `

    /*
        外部批次ID，用于幂等     */
    ExternalBatchId  *string `json:"external_batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto) SetExternalGoodsIdList(v []AlibabaJymItemExternalGoodsBatchOffsaleExternalGoodsIdDto) *AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto {
    s.ExternalGoodsIdList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto) SetExternalBatchId(v string) *AlibabaJymItemExternalGoodsBatchOffsaleGoodsOffSaleCommandDto {
    s.ExternalBatchId = &v
    return s
}
