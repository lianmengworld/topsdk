package domain


type AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto struct {
    /*
        是否支持找回包赔     */
    SupportRetrieveCompensation  *bool `json:"support_retrieve_compensation,omitempty" `

    /*
        卖家账号信息商品属性对象数组     */
    SellerAccountPropertyList  *[]AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto `json:"seller_account_property_list,omitempty" `

    /*
        商品属性对象数组     */
    GoodsPropertyList  *[]AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto `json:"goods_property_list,omitempty" `

    /*
        游戏属性对象     */
    GameProperty  *AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto `json:"game_property,omitempty" `

    /*
        二级类目ID     */
    SecondCategoryId  *int64 `json:"second_category_id,omitempty" `

    /*
        一级类目ID     */
    FirstCategoryId  *int64 `json:"first_category_id,omitempty" `

    /*
        商品图片url列表     */
    ImageUrlList  *[]AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDto `json:"image_url_list,omitempty" `

    /*
        商品基本信息     */
    GoodsBaseInfo  *AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto `json:"goods_base_info,omitempty" `

    /*
        外部商品ID，用于标识外部系统每次提交过来的商品     */
    ExternalGoodsId  *string `json:"external_goods_id,omitempty" `

    /*
        是否支持议价     */
    CanBargain  *bool `json:"can_bargain,omitempty" `

    /*
        入淘商品首图，可不传     */
    TaobaoFirstImage  *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDTO `json:"taobao_first_image,omitempty" `

    /*
        多子属性对象数组，仅在支持多子属性的类目使用     */
    MultiPropertyValueList  *[]AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO `json:"multi_property_value_list,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetSupportRetrieveCompensation(v bool) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.SupportRetrieveCompensation = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetSellerAccountPropertyList(v []AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.SellerAccountPropertyList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetGoodsPropertyList(v []AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.GoodsPropertyList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetGameProperty(v AlibabaJymItemExternalGoodsBatchPublishGamePropertyDto) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.GameProperty = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetSecondCategoryId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.SecondCategoryId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetFirstCategoryId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.FirstCategoryId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetImageUrlList(v []AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDto) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.ImageUrlList = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetGoodsBaseInfo(v AlibabaJymItemExternalGoodsBatchPublishGoodsBaseInfoDto) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.GoodsBaseInfo = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetExternalGoodsId(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.ExternalGoodsId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetCanBargain(v bool) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.CanBargain = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetTaobaoFirstImage(v AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDTO) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.TaobaoFirstImage = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto) SetMultiPropertyValueList(v []AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishDto {
    s.MultiPropertyValueList = &v
    return s
}
