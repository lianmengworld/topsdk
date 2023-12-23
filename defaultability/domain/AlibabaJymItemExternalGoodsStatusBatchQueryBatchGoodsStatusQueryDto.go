package domain


type AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto struct {
    /*
        每页数量，支持最大100     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        页数     */
    Page  *int64 `json:"page,omitempty" `

    /*
        商品ID数组，size最大支持100     */
    GoodsIdList  *[]int64 `json:"goods_id_list,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto) SetPageSize(v int64) *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto {
    s.PageSize = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto) SetPage(v int64) *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto {
    s.Page = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto) SetGoodsIdList(v []int64) *AlibabaJymItemExternalGoodsStatusBatchQueryBatchGoodsStatusQueryDto {
    s.GoodsIdList = &v
    return s
}
