package domain


type AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto struct {
    /*
        原图url     */
    OriginImage  *string `json:"origin_image,omitempty" `

    /*
        图片id     */
    ImageId  *int64 `json:"image_id,omitempty" `

    /*
        备注     */
    Note  *string `json:"note,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto) SetOriginImage(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto {
    s.OriginImage = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto) SetImageId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto {
    s.ImageId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto) SetNote(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsImageDto {
    s.Note = &v
    return s
}
