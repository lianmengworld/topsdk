package domain


type AlibabaJymItemExternalGoodsBatchDeleteGoodsBatchResultDTO struct {
    /*
        商品下架批次ID     */
    BatchId  *int64 `json:"batch_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchDeleteGoodsBatchResultDTO) SetBatchId(v int64) *AlibabaJymItemExternalGoodsBatchDeleteGoodsBatchResultDTO {
    s.BatchId = &v
    return s
}
