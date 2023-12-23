package domain


type AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO struct {
    /*
        批量删除商品id集合     */
    ExternalGoodsIdList  *[]AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO `json:"external_goods_id_list,omitempty" `

    /*
        外部批次ID，用于幂等     */
    ExternalBatchId  *string `json:"external_batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO) SetExternalGoodsIdList(v []AlibabaJymItemExternalGoodsBatchDeleteExternalGoodsIdDTO) *AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO {
    s.ExternalGoodsIdList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO) SetExternalBatchId(v string) *AlibabaJymItemExternalGoodsBatchDeleteGoodsDeleteCommandDTO {
    s.ExternalBatchId = &v
    return s
}
