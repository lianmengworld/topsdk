package domain


type AlibabaJymItemExternalGoodsBatchModifypriceGoodsBatchResultDto struct {
    /*
        商品改价批次ID     */
    BatchId  *int64 `json:"batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchModifypriceGoodsBatchResultDto) SetBatchId(v int64) *AlibabaJymItemExternalGoodsBatchModifypriceGoodsBatchResultDto {
    s.BatchId = &v
    return s
}
