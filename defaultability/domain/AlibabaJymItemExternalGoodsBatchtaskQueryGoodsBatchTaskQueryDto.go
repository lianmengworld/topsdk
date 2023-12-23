package domain


type AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskQueryDto struct {
    /*
        任务批次ID     */
    BatchId  *int64 `json:"batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskQueryDto) SetBatchId(v int64) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskQueryDto {
    s.BatchId = &v
    return s
}
