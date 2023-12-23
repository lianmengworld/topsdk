package domain


type AlibabaJymItemExternalGoodsBatchOnsaleGoodsBatchResultDto struct {
    /*
        商品上架批次ID     */
    BatchId  *int64 `json:"batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchOnsaleGoodsBatchResultDto) SetBatchId(v int64) *AlibabaJymItemExternalGoodsBatchOnsaleGoodsBatchResultDto {
    s.BatchId = &v
    return s
}
