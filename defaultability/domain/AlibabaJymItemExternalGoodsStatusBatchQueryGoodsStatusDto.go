package domain


type AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto struct {
    /*
        价格     */
    Price  *string `json:"price,omitempty" `

    /*
        商品状态     */
    Status  *int64 `json:"status,omitempty" `

    /*
        商品ID     */
    GoodsId  *int64 `json:"goods_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto) SetPrice(v string) *AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto {
    s.Price = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto) SetStatus(v int64) *AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto {
    s.Status = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto) SetGoodsId(v int64) *AlibabaJymItemExternalGoodsStatusBatchQueryGoodsStatusDto {
    s.GoodsId = &v
    return s
}
