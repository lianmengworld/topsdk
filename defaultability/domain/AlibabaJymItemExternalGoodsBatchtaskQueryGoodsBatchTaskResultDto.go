package domain


type AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskResultDto struct {
    /*
        商品批次子任务对象集合     */
    GoodsBatchSubTaskList  *[]AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask `json:"goods_batch_sub_task_list,omitempty" `

    /*
        任务状态     */
    Status  *int64 `json:"status,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskResultDto) SetGoodsBatchSubTaskList(v []AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchSubTask) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskResultDto {
    s.GoodsBatchSubTaskList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskResultDto) SetStatus(v int64) *AlibabaJymItemExternalGoodsBatchtaskQueryGoodsBatchTaskResultDto {
    s.Status = &v
    return s
}
