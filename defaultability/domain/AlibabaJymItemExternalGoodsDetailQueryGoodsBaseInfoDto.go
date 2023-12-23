package domain


type AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto struct {
    /*
        商品描述     */
    Description  *string `json:"description,omitempty" `

    /*
        库存     */
    Storage  *int64 `json:"storage,omitempty" `

    /*
        价格     */
    Price  *string `json:"price,omitempty" `

    /*
        标题     */
    Title  *string `json:"title,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto) SetDescription(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto {
    s.Description = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto) SetStorage(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto {
    s.Storage = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto) SetPrice(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto {
    s.Price = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto) SetTitle(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto {
    s.Title = &v
    return s
}
