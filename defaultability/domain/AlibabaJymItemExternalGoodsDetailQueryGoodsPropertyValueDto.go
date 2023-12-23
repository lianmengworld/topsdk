package domain


type AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto struct {
    /*
        属性值     */
    Value  *string `json:"value,omitempty" `

    /*
        属性值id     */
    ValueId  *int64 `json:"value_id,omitempty" `

    /*
        属性id     */
    PropertyId  *int64 `json:"property_id,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto) SetValue(v string) *AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto {
    s.Value = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto) SetValueId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto {
    s.ValueId = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto) SetPropertyId(v int64) *AlibabaJymItemExternalGoodsDetailQueryGoodsPropertyValueDto {
    s.PropertyId = &v
    return s
}
