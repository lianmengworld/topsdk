package domain


type AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDto struct {
    /*
        商品图片url     */
    ImageUrl  *string `json:"image_url,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDto) SetImageUrl(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsPublishImageDto {
    s.ImageUrl = &v
    return s
}
