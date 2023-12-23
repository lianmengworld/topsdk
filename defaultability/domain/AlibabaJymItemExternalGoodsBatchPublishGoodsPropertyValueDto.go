package domain


type AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto struct {
    /*
        属性值     */
    Value  *string `json:"value,omitempty" `

    /*
        属性值id，可不传     */
    ValueId  *int64 `json:"value_id,omitempty" `

    /*
        属性id     */
    PropertyId  *int64 `json:"property_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto) SetValue(v string) *AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto {
    s.Value = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto) SetValueId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto {
    s.ValueId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto) SetPropertyId(v int64) *AlibabaJymItemExternalGoodsBatchPublishGoodsPropertyValueDto {
    s.PropertyId = &v
    return s
}
