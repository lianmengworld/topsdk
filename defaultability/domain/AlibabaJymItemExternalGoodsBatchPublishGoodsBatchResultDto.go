package domain


type AlibabaJymItemExternalGoodsBatchPublishGoodsBatchResultDto struct {
    /*
        商品发布批次ID     */
    BatchId  *int64 `json:"batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsBatchResultDto) SetBatchId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGoodsBatchResultDto {
    s.BatchId = &v
    return s
}
