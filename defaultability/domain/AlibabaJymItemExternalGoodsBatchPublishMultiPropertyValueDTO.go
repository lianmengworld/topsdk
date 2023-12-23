package domain


type AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO struct {
    /*
        父属性名     */
    PropertyName  *string `json:"property_name,omitempty" `

    /*
        属性值对象数组     */
    PropertyValueList  *[]AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO `json:"property_value_list,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO) SetPropertyName(v string) *AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO {
    s.PropertyName = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO) SetPropertyValueList(v []AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO) *AlibabaJymItemExternalGoodsBatchPublishMultiPropertyValueDTO {
    s.PropertyValueList = &v
    return s
}
