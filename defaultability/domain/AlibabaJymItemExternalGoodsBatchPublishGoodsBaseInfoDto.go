package domain


type AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto struct {
    /*
        标题     */
    Title  *string `json:"title,omitempty" `

    /*
        价格     */
    Price  *string `json:"price,omitempty" `

    /*
        库存     */
    Storage  *int64 `json:"storage,omitempty" `

    /*
        商品描述     */
    Description  *string `json:"description,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto) SetTitle(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto {
    s.Title = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto) SetPrice(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto {
    s.Price = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto) SetStorage(v int64) *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto {
    s.Storage = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto) SetDescription(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto {
    s.Description = &v
    return s
}
