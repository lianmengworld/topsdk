package domain


type AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto struct {
    /*
        外部批次ID，用于幂等     */
    ExternalBatchId  *string `json:"external_batch_id,omitempty" `

    /*
        商品发布数据体     */
    GoodsList  *[]AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto `json:"goods_list,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto) SetExternalBatchId(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto {
    s.ExternalBatchId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto) SetGoodsList(v []AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishCommandDto {
    s.GoodsList = &v
    return s
}
