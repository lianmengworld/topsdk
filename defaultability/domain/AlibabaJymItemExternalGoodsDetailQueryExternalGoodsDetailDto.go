package domain


type AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto struct {
    /*
        商品基本信息     */
    GoodsBaseInfo  *AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto `json:"goods_base_info,omitempty" `

    /*
        类目     */
    Category  *AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto `json:"category,omitempty" `

    /*
        商品图片url列表     */
    ImageList  *[]AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto `json:"image_list,omitempty" `

    /*
        游戏属性     */
    GameProperty  *AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto `json:"game_property,omitempty" `

    /*
        游戏属性对     */
    GoodsPropertyList  *[]AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto `json:"goods_property_list,omitempty" `

    /*
        卖家账号信息商品属性对     */
    SellerAccountPropertyList  *[]AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto `json:"seller_account_property_list,omitempty" `

    /*
        是否支持找回包赔     */
    SupportRetrieveCompensation  *bool `json:"support_retrieve_compensation,omitempty" `

    /*
        是否支持议价     */
    CanBargain  *bool `json:"can_bargain,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetGoodsBaseInfo(v AlibabaJymItemExternalGoodsDetailQueryGoodsBaseInfoDto) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.GoodsBaseInfo = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetCategory(v AlibabaJymItemExternalGoodsDetailQueryGoodsCategoryDto) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.Category = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetImageList(v []AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.ImageList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetGameProperty(v AlibabaJymItemExternalGoodsDetailQueryGamePropertyDto) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.GameProperty = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetGoodsPropertyList(v []AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.GoodsPropertyList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetSellerAccountPropertyList(v []AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.SellerAccountPropertyList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetSupportRetrieveCompensation(v bool) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.SupportRetrieveCompensation = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto) SetCanBargain(v bool) *AlibabaJymItemExternalGoodsDetailQueryExternalGoodsDetailDto {
    s.CanBargain = &v
    return s
}
