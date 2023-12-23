package domain


type AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO struct {
    /*
        属性值     */
    Value  *string `json:"value,omitempty" `

    /*
        子属性列表     */
    ChildPropertyKeyValueList  *[]AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO `json:"child_property_key_value_list,omitempty" `

}

func (s *AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO) SetValue(v string) *AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO {
    s.Value = &v
    return s
}
func (s *AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO) SetChildPropertyKeyValueList(v []AlibabaJymItemExternalGoodsBatchPublishChildPropertyKeyValueDTO) *AlibabaJymItemExternalGoodsBatchPublishPropertyValueDTO {
    s.ChildPropertyKeyValueList = &v
    return s
}
