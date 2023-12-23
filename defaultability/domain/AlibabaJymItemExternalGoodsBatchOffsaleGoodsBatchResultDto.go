package domain


type AlibabaJymItemExternalGoodsBatchOffsaleGoodsBatchResultDto struct {
    /*
        商品下架批次ID     */
    BatchId  *int64 `json:"batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchOffsaleGoodsBatchResultDto) SetBatchId(v int64) *AlibabaJymItemExternalGoodsBatchOffsaleGoodsBatchResultDto {
    s.BatchId = &v
    return s
}
