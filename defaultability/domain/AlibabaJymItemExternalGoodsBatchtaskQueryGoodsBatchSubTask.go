package domain


type AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask struct {
    /*
        子任务状态产生原因     */
    Reason  *string `json:"reason,omitempty" `

    /*
        子任务状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        子任务id     */
    SubBatchId  *int64 `json:"sub_batch_id,omitempty" `

    /*
        交易猫商品ID，如果商品发布失败则为空     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

    /*
        外部商品ID，用于标识外部系统每次提交过来的商品     */
    ExternalGoodsId  *string `json:"external_goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask) SetReason(v string) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask {
    s.Reason = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask) SetStatus(v int64) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask {
    s.Status = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask) SetSubBatchId(v int64) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask {
    s.SubBatchId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask {
    s.GoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask) SetExternalGoodsId(v string) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask {
    s.ExternalGoodsId = &v
    return s
}
